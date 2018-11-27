// Code generated by "bitfanDoc "; DO NOT EDIT
package drop

import "bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "drop",
  ImportPath: "bitfan/processors/filter-drop",
  Doc:        "Drops everything received\nDrops everything that gets to this filter.\n\nThis is best used in combination with conditionals, for example:\n```\nfilter {\n  if [loglevel] == \"debug\" {\n    drop { }\n  }\n}\n```\nThe above will only pass events to the drop filter if the loglevel field is debug. This will cause all events matching to be dropped.",
  DocShort:   "Drops all events",
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
        Name:           "Percentage",
        Alias:          "",
        Doc:            "Drop all the events within a pre-configured percentage.\nThis is useful if you just need a percentage but not the whole.",
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