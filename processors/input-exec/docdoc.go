// Code generated by "bitfanDoc "; DO NOT EDIT
package execinput

import "bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
		Name:       "execinput",
		ImportPath: "bitfan/processors/input-exec",
		Doc:        "",
		DocShort:   "",
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
					Name:           "Command",
					Alias:          "",
					Doc:            "",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Args",
					Alias:          "",
					Doc:            "",
					Required:       false,
					Type:           "array",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Interval",
					Alias:          "",
					Doc:            "",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Codec",
					Alias:          "codec",
					Doc:            "The codec used for input data. Input codecs are a convenient method for decoding\nyour data before it enters the input, without needing a separate filter in your bitfan pipeline",
					Required:       false,
					Type:           "codec",
					DefaultValue:   "\"plain\"",
					PossibleValues: []string{},
					ExampleLS:      "",
				},
			},
		},
		Ports: []*doc.ProcessorPort{},
	}
}
