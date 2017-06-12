package location

// Code comes from https://github.com/spf13/hugo/tree/master/tpl

import (
	"html"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cast"
	"github.com/vjeantet/jodaTime"
)

type templateFunctions struct{}

// HTMLUnescape returns a copy of with HTML escape requences converted to plain
// text.
func (t *templateFunctions) htmlUnescape(s interface{}) (string, error) {
	ss, err := cast.ToStringE(s)
	if err != nil {
		return "", err
	}

	return html.UnescapeString(ss), nil
}

// HTMLEscape returns a copy of s with reserved HTML characters escaped.
func (t *templateFunctions) htmlEscape(s interface{}) (string, error) {
	ss, err := cast.ToStringE(s)
	if err != nil {
		return "", err
	}

	return html.EscapeString(ss), nil
}

// ToUpper returns a copy of the input s with all Unicode letters mapped to their
// upper case.
func (t *templateFunctions) upper(s interface{}) (string, error) {
	ss, err := cast.ToStringE(s)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(ss), nil
}

// ToLower returns a copy of the input s with all Unicode letters mapped to their
// lower case.
func (t *templateFunctions) lower(s interface{}) (string, error) {
	ss, err := cast.ToStringE(s)
	if err != nil {
		return "", err
	}

	return strings.ToLower(ss), nil
}

// Trim returns a string with all leading and trailing characters defined
// contained in cutset removed.
func (t *templateFunctions) trim(s, cutset interface{}) (string, error) {
	ss, err := cast.ToStringE(s)
	if err != nil {
		return "", err
	}

	sc, err := cast.ToStringE(cutset)
	if err != nil {
		return "", err
	}

	return strings.Trim(ss, sc), nil
}

// Format converts the textual representation of the datetime string into
// the other form or returns it of the time.Time value. These are formatted
// with the layout string
func (t *templateFunctions) dateFormat(layout string, v interface{}) string {
	ti, err := cast.ToTimeE(v)
	if err != nil {
		return ""
	}
	return jodaTime.Format(layout, ti)
}

func (t *templateFunctions) timeStampFormat(layout string, v map[string]interface{}) string {
	ts := v["@timestamp"]
	ti, err := cast.ToTimeE(ts)
	if err != nil {
		return ""
	}
	return jodaTime.Format(layout, ti)
}

// AsTime converts the textual representation of the datetime string into
// a time.Time interface.
func (t *templateFunctions) asTime(v interface{}) interface{} {
	ti, err := cast.ToTimeE(v)
	if err != nil {
		return nil
	}

	return ti
}

// Now returns the current local time.
func (t *templateFunctions) now() time.Time {
	return time.Now()
}

// NumFmt formats a number with the given precision using the
// negative, decimal, and grouping options.  The `options`
// parameter is a string consisting of `<negative> <decimal> <grouping>`.  The
// default `options` value is `- . ,`.
//
// Note that numbers are rounded up at 5 or greater.
// So, with precision set to 0, 1.5 becomes `2`, and 1.4 becomes `1`.
func (t *templateFunctions) numFmt(precision, number interface{}, options ...interface{}) string {
	prec, err := cast.ToIntE(precision)
	if err != nil {
		return ""
	}

	n, err := cast.ToFloat64E(number)
	if err != nil {
		return ""
	}

	var neg, dec, grp string

	if len(options) == 0 {
		// TODO(moorereason): move to site config
		neg, dec, grp = "-", ".", ","
	} else {
		s, err := cast.ToStringE(options[0])
		if err != nil {
			return ""
		}

		rs := strings.Fields(s)
		switch len(rs) {
		case 0:
		case 1:
			neg = rs[0]
		case 2:
			neg, dec = rs[0], rs[1]
		case 3:
			neg, dec, grp = rs[0], rs[1], rs[2]
		default:
			return ""
		}
	}

	// Logic from MIT Licensed github.com/go-playground/locales/
	// Original Copyright (c) 2016 Go Playground

	s := strconv.FormatFloat(math.Abs(n), 'f', prec, 64)
	L := len(s) + 2 + len(s[:len(s)-1-prec])/3

	var count int
	inWhole := prec == 0
	b := make([]byte, 0, L)

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			for j := len(dec) - 1; j >= 0; j-- {
				b = append(b, dec[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(grp) - 1; j >= 0; j-- {
					b = append(b, grp[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if n < 0 {
		for j := len(neg) - 1; j >= 0; j-- {
			b = append(b, neg[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}
