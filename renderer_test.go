package sqlfmt

import (
	"bytes"
	"testing"
)

func TestTextRenderer(t *testing.T) {
	var buf bytes.Buffer

	tr := NewTextRenderer(&buf)

	tr.Text("select", "")
	tr.NewLine()
	tr.Indent()
	tr.Text("foo", "")
	tr.NewLine()
	tr.Unindent()
	tr.Text("from", "")
	tr.NewLine()
	tr.Indent()
	tr.Text("bar", "")

	expected := `select
  foo
from
  bar`

	if buf.String() != expected {
		t.Errorf("Expected `%s`, got `%s`", expected, buf.String())
	}
}
