// Code generated by "bitfanDoc "; DO NOT EDIT
package null

import "bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "null",
  ImportPath: "bitfan/processors/output-null",
  Doc:        "Drops everything received",
  DocShort:   "Drops everything received",
  Options:    (*doc.ProcessorOptions)(nil),
  Ports:      []*doc.ProcessorPort{},
}
}