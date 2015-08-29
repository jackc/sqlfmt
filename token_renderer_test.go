package sqlfmt

import (
	"reflect"
	"testing"
)

func TestTokenRenderer(t *testing.T) {
	r := TokenRenderer(nil)
	r.Text("select", KeywordToken)
	r.Control(NewLineToken)
	r.Control(IndentToken)
	r.Text("foo", IdentifierToken)
	r.Control(NewLineToken)
	r.Control(UnindentToken)
	r.Text("from", KeywordToken)
	r.Control(NewLineToken)
	r.Control(IndentToken)
	r.Text("bar", IdentifierToken)

	expected := []RenderToken{
		{Type: KeywordToken, Value: "select"},
		{Type: NewLineToken},
		{Type: IndentToken},
		{Type: IdentifierToken, Value: "foo"},
		{Type: NewLineToken},
		{Type: UnindentToken},
		{Type: KeywordToken, Value: "from"},
		{Type: NewLineToken},
		{Type: IndentToken},
		{Type: IdentifierToken, Value: "bar"},
	}

	if !reflect.DeepEqual([]RenderToken(r), expected) {
		t.Errorf("Expected `%v`, got `%v`", expected, []RenderToken(r))
	}
}
