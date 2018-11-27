// Code generated by "bitfanDoc "; DO NOT EDIT
package use

import "bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "use",
  ImportPath: "bitfan/processors/use",
  Doc:        "When used in input (input->filter->o) the processor will receive events from the last filter from the pipeline used in configuration file.\n\nWhen used in filter (i->filter->o) the processor will\n\n* pass the event to the first filter plugin found in the used configuration file\n* receive events from the last filter plugin found in the used configuration file\n\nWhen used in output (i->filter->output->o) the processor will\n\n* pass the event to the first filter plugin found in the used configuration file",
  DocShort:   "Include a config file",
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
        Name:           "Path",
        Alias:          "path",
        Doc:            "Path to configuration to import in this pipeline, it could be a local file or an url\ncan be relative path to the current configuration.\n\nSPLIT and JOIN : in filter Section, set multiples path to make a split and join into your pipeline",
        Required:       true,
        Type:           "array",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "path=> [\"meteo-input.conf\"]",
      },
      &doc.ProcessorOption{
        Name:           "Var",
        Alias:          "var",
        Doc:            "You can set variable references in the used configuration by using ${var}.\neach reference will be replaced by the value of the variable found in this option\n\nThe replacement is case-sensitive.",
        Required:       false,
        Type:           "hash",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "var => {\"hostname\"=>\"myhost\",\"varname\"=>\"varvalue\"}",
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