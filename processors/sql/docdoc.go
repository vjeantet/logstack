// Code generated by "bitfanDoc "; DO NOT EDIT
package sqlprocessor

import "bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "sqlprocessor",
  ImportPath: "bitfan/processors/sql",
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
        Name:           "Driver",
        Alias:          "driver",
        Doc:            "GOLANG driver class to load, for example, \"mysql\".",
        Required:       true,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "driver => \"mysql\"",
      },
      &doc.ProcessorOption{
        Name:           "EventBy",
        Alias:          "event_by",
        Doc:            "Send an event row by row or one event with all results\npossible values \"row\", \"result\"",
        Required:       false,
        Type:           "string",
        DefaultValue:   "\"row\"",
        PossibleValues: []string{
          "\"row\"",
          "\"result\"",
        },
        ExampleLS: "",
      },
      &doc.ProcessorOption{
        Name:           "Statement",
        Alias:          "statement",
        Doc:            "SQL Statement\n\nWhen there are more than 1 statement, only data from the last one will generate events.",
        Required:       true,
        Type:           "location",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "statement => \"SELECT * FROM mytable\"",
      },
      &doc.ProcessorOption{
        Name:           "Interval",
        Alias:          "interval",
        Doc:            "Set an interval when this processor is used as a input",
        Required:       false,
        Type:           "interval",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "interval => \"10\"",
      },
      &doc.ProcessorOption{
        Name:           "ConnectionString",
        Alias:          "connection_string",
        Doc:            "",
        Required:       true,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "connection_string => \"username:password@tcp(192.168.1.2:3306)/mydatabase?charset=utf8\"",
      },
      &doc.ProcessorOption{
        Name:           "Var",
        Alias:          "var",
        Doc:            "You can set variable to be used in Statements by using ${var}.\neach reference will be replaced by the value of the variable found in Statement's content\nThe replacement is case-sensitive.",
        Required:       false,
        Type:           "hash",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "var => {\"hostname\"=>\"myhost\",\"varname\"=>\"varvalue\"}",
      },
      &doc.ProcessorOption{
        Name:           "Target",
        Alias:          "target",
        Doc:            "Define the target field for placing the retrieved data. If this setting is omitted,\nthe data will be stored in the \"data\" field\nSet the value to \".\" to store value to the root (top level) of the event",
        Required:       false,
        Type:           "string",
        DefaultValue:   "\"data\"",
        PossibleValues: []string{},
        ExampleLS:      "target => \"data\"",
      },
    },
  },
  Ports: []*doc.ProcessorPort{},
}
}