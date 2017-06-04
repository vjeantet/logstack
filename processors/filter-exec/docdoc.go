// Code generated by "bitfanDoc "; DO NOT EDIT
package exec

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "exec",
  ImportPath: "/Users/sodadi/go/src/github.com/vjeantet/bitfan/processors/filter-exec",
  Doc:        "",
  DocShort:   "drop event when field value is the same in the last event",
  Options:    &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:           "AddField",
        Alias:          "add_field",
        Doc:            "If this filter is successful, add any arbitrary fields to this event.",
        Required:       false,
        Type:           "hash",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "AddTag",
        Alias:          "add_tag",
        Doc:            "If this filter is successful, add arbitrary tags to the event. Tags can be dynamic\nand include parts of the event using the %{field} syntax.",
        Required:       false,
        Type:           "array",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "RemoveField",
        Alias:          "remove_field",
        Doc:            "If this filter is successful, remove arbitrary fields from this event. Example:\n` kv {\n`   remove_field => [ \"foo_%{somefield}\" ]\n` }",
        Required:       false,
        Type:           "array",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "RemoveTag",
        Alias:          "remove_tag",
        Doc:            "If this filter is successful, remove arbitrary tags from the event. Tags can be dynamic and include parts of the event using the %{field} syntax.\nExample:\n` kv {\n`   remove_tag => [ \"foo_%{somefield}\" ]\n` }\nIf the event has field \"somefield\" == \"hello\" this filter, on success, would remove the tag foo_hello if it is present. The second example would remove a sad, unwanted tag as well.",
        Required:       false,
        Type:           "array",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Command",
        Alias:          "command",
        Doc:            "",
        Required:       true,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Args",
        Alias:          "args",
        Doc:            "",
        Required:       false,
        Type:           "array",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Stdin",
        Alias:          "stdin",
        Doc:            "Pass the complete event to stdin as a json string",
        Required:       false,
        Type:           "bool",
        DefaultValue:   "false",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Target",
        Alias:          "target",
        Doc:            "Where do the output should be stored\nSet \".\" when output is json formated and want to replace current event fields with output\nresponse. (usefull )",
        Required:       false,
        Type:           "string",
        DefaultValue:   "\"stdout\"",
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