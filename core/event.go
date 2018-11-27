package core

import (
	"time"

	"github.com/clbanning/mxj"
	"bitfan/processors"
)

// event represents data sent to agents (or received by agents)
type event struct {
	fields mxj.Map
}

func (e *event) Fields() *mxj.Map {
	return &e.fields
}

func (e *event) SetFields(f map[string]interface{}) {
	e.fields = f
}

func (e *event) Message() string {
	return e.Fields().ValueOrEmptyForPathString("message")
}

func (e *event) SetMessage(s string) {
	e.Fields().SetValueForPath(s, "message")
}

func (e *event) Clone() processors.IPacket {
	nf, _ := e.Fields().Copy()
	nf["@timestamp"], _ = e.Fields().ValueForPath("@timestamp")
	return newPacket(nf)
}

func newPacket(fields map[string]interface{}) processors.IPacket {
	if fields == nil {
		fields = mxj.Map{}
	}

	if _, k := fields["@timestamp"]; !k {
		fields["@timestamp"] = time.Now()
	} else {
		switch v := fields["@timestamp"].(type) {
		case string:
			var err error
			if fields["@timestamp"], err = time.Parse(time.RFC3339, v); err != nil {
				fields["@timestamp"] = v
			}
		}
	}

	return &event{
		fields: fields,
	}
}
