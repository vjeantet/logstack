package tcpinput

import (
	"bitfan/processors/doc"
	"bitfan/processors/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	p := New()
	_, ok := p.(*processor)
	assert.Equal(t, ok, true, "New() should return a processor")
}
func TestDoc(t *testing.T) {
	assert.IsType(t, &doc.Processor{}, New().(*processor).Doc())
}
func TestConfigure(t *testing.T) {
	conf := map[string]interface{}{}
	ctx := testutils.NewProcessorContext()
	p := New()
	err := p.Configure(ctx, conf)
	assert.Nil(t, err, "Configure() processor without error")
}
