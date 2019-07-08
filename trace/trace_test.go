package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var out bytes.Buffer
	tracer := New(&out)
	if tracer == nil {
		t.Error("call to New should not return nil. want Tracer. got nil")
	} else {
		payload := "Trace this"
		tracer.Trace(payload)
		if out.String() != "Trace this\n" {
			t.Errorf("want trace to be '%s', got '%s'.", payload, out.String())
		}
	}
}
