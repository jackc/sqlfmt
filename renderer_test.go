package sqlfmt

import (
	"bytes"
	"testing"
)

func TestTextRenderer(t *testing.T) {
	var buf bytes.Buffer

	tr := NewTextRenderer(&buf)

	tr.Keyword("select")
	tr.NewLine()
	tr.Indent()
	tr.Identifier("foo")
	tr.Symbol(",")
	tr.Identifier("bar")
	tr.NewLine()
	tr.Unindent()
	tr.Keyword("from")
	tr.NewLine()
	tr.Indent()
	tr.Identifier("baz")

	expected := `select
  foo, bar
from
  baz`

	if buf.String() != expected {
		t.Errorf("Expected `%s`, got `%s`", expected, buf.String())
	}
}
