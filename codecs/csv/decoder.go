//go:generate bitfanDoc -codec csv
// Parses comma-separated value data into individual fields
package csvcodec

import (
	"encoding/csv"
	"fmt"
	"io"

	"bitfan/commons"
	"github.com/mitchellh/mapstructure"
)

type decoder struct {
	more        bool
	r           *csv.Reader
	columnnames []string
	options     decoderOptions
	comma       rune
	log         commons.Logger
	title       bool
}

// Parses comma-separated value data into individual fields
type decoderOptions struct {

	// Define the column separator value. If this is not specified, the default is a comma ,. Optional
	// @Default ","
	Separator string `mapstructure:"separator"`

	// Define whether column names should autogenerated or not. Defaults to true.
	// If set to false, columns not having a header specified will not be parsed.
	// @Default true
	AutogenerateColumnNames bool `mapstructure:"autogenerate_column_names"`

	// Define the character used to quote CSV fields. If this is not specified the default is a double quote ". Optional.
	// @Default "\""
	QuoteChar string `mapstructure:"quote_char"`

	// Define a list of column names (in the order they appear in the CSV, as
	// if it were a header line).
	//
	// If columns is not configured, or there are not enough columns specified,
	// the default column names are "column1", "column2", etc.
	//
	// In the case that there are more columns in the data than specified in this column
	// list, extra columns will be auto-numbered:
	// (e.g. "user_defined_1", "user_defined_2", "column3", "column4", etc.)
	Columns []string `mapstructure:"columns"`

	// Define the comment character.
	// Lines beginning with the Comment character without preceding whitespace are ignored.
	Comment string `mapstructure:"comment"`
}

func NewDecoder(r io.Reader) *decoder {
	d := &decoder{
		r:    csv.NewReader(r),
		more: true,
		options: decoderOptions{
			Separator:               ",",
			AutogenerateColumnNames: true,
			QuoteChar:               "\"",
		},
		title: true,
		comma: ',',
	}

	return d
}
func (d *decoder) SetOptions(conf map[string]interface{}, logger commons.Logger, cwl string) error {
	d.log = logger

	if err := mapstructure.Decode(conf, &d.options); err != nil {
		return err
	}

	d.r.Comma = []rune(d.options.Separator)[0]
	d.comma = d.r.Comma

	if len(d.options.Columns) > 0 {
		d.columnnames = d.options.Columns
	}

	if d.options.Comment != "" {
		d.r.Comment = []rune(d.options.Comment)[0]
	}

	return nil
}

func (d *decoder) Decode(data *interface{}) error {
	*data = nil
	record, err := d.r.Read()

	if err == io.EOF {
		d.more = false
		return err
	}

	if d.title && d.options.AutogenerateColumnNames {
		d.columnnames = record
		d.title = false
		return d.Decode(data)
	}

	*data = map[string]interface{}{}

	for i := range record {
		var column string
		if i >= len(d.columnnames) {
			column = fmt.Sprintf("column%d", i+1)
		} else {
			column = d.columnnames[i]
		}
		(*data).(map[string]interface{})[column] = record[i]
	}

	return nil
}

func (d *decoder) More() bool {
	return d.more
}

func (d *decoder) Buffer() []byte {
	return []byte{}
}
