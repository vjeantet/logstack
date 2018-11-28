// Code generated by "bitfanDoc "; DO NOT EDIT
package change

import "bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
		Name:       "change",
		ImportPath: "bitfan/processors/filter-change",
		Doc:        "This rule will monitor a certain field and match if that field changes. The field must change with respect to the last event",
		DocShort:   "drop event when field value is the same in the last event",
		Options: &doc.ProcessorOptions{
			Doc: "",
			Options: []*doc.ProcessorOption{
				&doc.ProcessorOption{
					Name:           "processors.CommonOptions",
					Alias:          ",squash",
					Doc:            "",
					Required:       false,
					Type:           "processors.CommonOptions",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "CompareField",
					Alias:          "compare_field",
					Doc:            "The name of the field to use to compare to the blacklist.\n\nIf the field is null, those events will be ignored.",
					Required:       true,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "compare_field => \"message\"",
				},
				&doc.ProcessorOption{
					Name:           "IgnoreMissing",
					Alias:          "ignore_missing",
					Doc:            "If true, events without a compare_key field will not count as changed.",
					Required:       false,
					Type:           "bool",
					DefaultValue:   "true",
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Timeframe",
					Alias:          "timeframe",
					Doc:            "The maximum time in seconds between changes. After this time period, Bitfan will forget the old value of the compare_field field.",
					Required:       false,
					Type:           "int",
					DefaultValue:   "0 (no timeframe)",
					PossibleValues: []string{},
					ExampleLS:      "timeframe => 10",
				},
			},
		},
		Ports: []*doc.ProcessorPort{
			&doc.ProcessorPort{
				Default: true,
				Name:    "PORT_SUCCESS",
				Number:  0,
				Doc:     "",
			},
		},
	}
}
