package sqlfmt

import (
	"bytes"
)

type TokenRenderer []RenderToken

func (r *TokenRenderer) Text(val string, tokenType int) {
	*r = TokenRenderer(append([]RenderToken(*r), RenderToken{Type: tokenType, Value: val}))
}

func (r *TokenRenderer) Control(tokenType int) {
	*r = TokenRenderer(append([]RenderToken(*r), RenderToken{Type: tokenType}))
}

func RenderTokens(r Renderer, tokens []RenderToken) {
	for _, t := range tokens {
		if t.Value != "" {
			r.Text(t.Value, t.Type)
		} else {
			r.Control(t.Type)
		}
	}
}

func TryOneLine(tokens []RenderToken, maxLineLength int) []RenderToken {
	buf := &bytes.Buffer{}
	r := NewTextRenderer(buf)
	RenderTokens(r, tokens)
	if buf.Len() < maxLineLength {
		filteredTokens := []RenderToken{}
		for _, t := range tokens {
			switch t.Type {
			case NewLineToken, IndentToken, UnindentToken:
				// filter these out
			default:
				filteredTokens = append(filteredTokens, t)
			}
		}

		tokens = filteredTokens
	}

	return tokens
}
