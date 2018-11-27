// Code generated by "bitfanDoc "; DO NOT EDIT
package elasticinput

import "bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "elasticinput",
  ImportPath: "bitfan/processors/input-elasticsearch",
  Doc:        "",
  DocShort:   "",
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
        Name:           "Hosts",
        Alias:          "",
        Doc:            "",
        Required:       false,
        Type:           "array",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Query",
        Alias:          "",
        Doc:            "",
        Required:       false,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Size",
        Alias:          "",
        Doc:            "",
        Required:       false,
        Type:           "int",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "User",
        Alias:          "",
        Doc:            "",
        Required:       false,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Password",
        Alias:          "",
        Doc:            "",
        Required:       false,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
    },
  },
  Ports: []*doc.ProcessorPort{},
}
}