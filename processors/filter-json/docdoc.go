// Code generated by "bitfanDoc "; DO NOT EDIT
package json

import "bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "json",
  ImportPath: "bitfan/processors/filter-json",
  Doc:        "",
  DocShort:   "Parses JSON events",
  Options:    &doc.ProcessorOptions{
    Doc:     "",
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
        Name:           "SkipOnInvalidJson",
        Alias:          "skip_on_invalid_json",
        Doc:            "Allow to skip filter on invalid json",
        Required:       false,
        Type:           "bool",
        DefaultValue:   "false",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Source",
        Alias:          "source",
        Doc:            "The configuration for the JSON filter",
        Required:       true,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Target",
        Alias:          "target",
        Doc:            "Define the target field for placing the parsed data. If this setting is omitted,\nthe JSON data will be stored at the root (top level) of the event",
        Required:       false,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "TagOnFailure",
        Alias:          "tag_on_failure",
        Doc:            "Append values to the tags field when there has been no successful match",
        Required:       false,
        Type:           "array",
        DefaultValue:   "[\"_jsonparsefailure\"]",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
    },
  },
  Ports: []*doc.ProcessorPort{},
}
}