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
	tr.NewLine()
	tr.Unindent()
	tr.Keyword("from")
	tr.NewLine()
	tr.Indent()
	tr.Identifier("bar")

	expected := `select
  foo
from
  bar`

	if buf.String() != expected {
		t.Errorf("Expected `%s`, got `%s`", expected, buf.String())
	}
}
