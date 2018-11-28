// Code generated by "bitfanDoc "; DO NOT EDIT
package websocketinput

import "bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
		Name:       "websocketinput",
		ImportPath: "bitfan/processors/input-websocket",
		Doc:        "Receive event on a ws connection",
		DocShort:   "Reads events from standard input",
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
					Name:           "Codec",
					Alias:          "",
					Doc:            "The codec used for outputed data.",
					Required:       false,
					Type:           "codec",
					DefaultValue:   "\"json\"",
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Uri",
					Alias:          "",
					Doc:            "URI path",
					Required:       false,
					Type:           "string",
					DefaultValue:   "\"wsin\"",
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "MaxMessageSize",
					Alias:          "max_message_size",
					Doc:            "Maximum message size allowed from peer.",
					Required:       false,
					Type:           "int",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
			},
		},
		Ports: []*doc.ProcessorPort{},
	}
}
