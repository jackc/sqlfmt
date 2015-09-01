package sqlfmt

import (
	"io"
)

const (
	NullToken        = iota
	KeywordToken     = iota
	IdentifierToken  = iota
	SymbolToken      = iota
	ConstantToken    = iota
	SpaceToken       = iota
	RefuseSpaceToken = iota
	NewLineToken     = iota
	IndentToken      = iota
	UnindentToken    = iota
)

type RenderToken struct {
	Type  int
	Value string
}

type Renderer interface {
	Text(val string, typ int)
	Control(typ int)
}

type TextRenderer struct {
	w               io.Writer
	err             error
	indentLvl       int
	indent          string
	lineIndented    bool
	newLine         bool
	lastRenderToken RenderToken
}

func (left RenderToken) SpaceBetween(right RenderToken) bool {
	if left.Type == RefuseSpaceToken {
		return false
	}

	switch left.Type {
	case KeywordToken, IdentifierToken, ConstantToken:
		switch right.Type {
		case KeywordToken, IdentifierToken, ConstantToken:
			return true
		case SymbolToken:
			switch right.Value {
			case "[", "(", "]", ")", ".", ",", "::", ":":
				return false
			}
			return true
		}
	case SymbolToken:
		switch left.Value {
		case ".", "(", "[", "::", ":":
			return false
		}

		if right.Type == NewLineToken {
			return false
		}

		if left.Value == "," {
			return true
		}

		if right.Type == SymbolToken {
			switch right.Value {
			case ".", "(", "[", "::", ")", "]", ",", ":":
				return false
			}
		}

		return true
	}

	return false
}

func NewTextRenderer(w io.Writer) *TextRenderer {
	return &TextRenderer{w: w, indent: "  "}
}

func (tr *TextRenderer) Text(val string, tokenType int) {
	if !tr.lineIndented {
		for i := 0; i < tr.indentLvl; i++ {
			_, tr.err = io.WriteString(tr.w, tr.indent)
			if tr.err != nil {
				return
			}
		}

		tr.lineIndented = true
	}

	token := RenderToken{Type: tokenType, Value: val}

	if tr.newLine {
		tr.newLine = false
	} else if tr.lastRenderToken.SpaceBetween(token) {
		_, tr.err = io.WriteString(tr.w, " ")
		if tr.err != nil {
			return
		}
	}

	tr.lastRenderToken = token

	_, tr.err = io.WriteString(tr.w, val)
}

func (tr *TextRenderer) Control(typ int) {
	if tr.err != nil {
		return
	}

	switch typ {
	case SpaceToken:
		_, tr.err = io.WriteString(tr.w, " ")
	case NewLineToken:
		tr.renderNewLine()
	case IndentToken:
		tr.indentLvl = tr.indentLvl + 1
	case UnindentToken:
		tr.indentLvl = tr.indentLvl - 1
	}

	tr.lastRenderToken = RenderToken{Type: typ}
}

func (tr *TextRenderer) renderNewLine() {
	if tr.newLine {
		return
	}

	tr.newLine = true
	tr.lastRenderToken = RenderToken{Type: NewLineToken}

	_, tr.err = io.WriteString(tr.w, "\n")
	if tr.err != nil {
		return
	}

	tr.lineIndented = false
}

func (tr *TextRenderer) Error() error {
	return tr.err
}
