// Code generated by "bitfanDoc "; DO NOT EDIT
package whitelist

import "bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "whitelist",
  ImportPath: "bitfan/processors/filter-whitelist",
  Doc:        "Similar to blacklist, this processor will compare a certain field to a whitelist, and match\nif the list does not contain the term",
  DocShort:   "drop event when term is in a given list",
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
        Name:           "CompareField",
        Alias:          "compare_field",
        Doc:            "The name of the field to use to compare to the whitelist.\nIf the field is null, those events will be ignored.",
        Required:       true,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "compare_field => \"message\"",
      },
      &doc.ProcessorOption{
        Name:           "IgnoreMissing",
        Alias:          "ignore_missing",
        Doc:            "If true, events without a compare_key field will not match.",
        Required:       false,
        Type:           "bool",
        DefaultValue:   "true",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Terms",
        Alias:          "terms",
        Doc:            "A list of whitelisted terms.\nThe compare_field term must be in this list or else it will match.",
        Required:       true,
        Type:           "array",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "terms => [\"val1\",\"val2\",\"val3\"]",
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