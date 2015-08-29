package sqlfmt

import (
	"bytes"
	"testing"
)

func TestTextRenderer(t *testing.T) {
	var buf bytes.Buffer

	tr := NewTextRenderer(&buf)

	tr.Text("select", KeywordToken)
	tr.Control(NewLineToken)
	tr.Control(IndentToken)
	tr.Text("foo", IdentifierToken)
	tr.Text(",", SymbolToken)
	tr.Text("bar", IdentifierToken)
	tr.Control(NewLineToken)
	tr.Control(UnindentToken)
	tr.Text("from", KeywordToken)
	tr.Control(NewLineToken)
	tr.Control(IndentToken)
	tr.Text("baz", IdentifierToken)

	expected := `select
  foo, bar
from
  baz`

	if buf.String() != expected {
		t.Errorf("Expected `%s`, got `%s`", expected, buf.String())
	}
}
