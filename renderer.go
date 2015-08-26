package sqlfmt

import (
	"io"
)

type Renderer interface {
	Keyword(val string)
	Identifier(val string)
	Symbol(val string)
	Constant(val string)
	Space()
	NewLine()
	Indent()
	Unindent()
}

type TextRenderer struct {
	w            io.Writer
	err          error
	indentLvl    int
	indent       string
	lineIndented bool
	newLine      bool
}

func NewTextRenderer(w io.Writer) *TextRenderer {
	return &TextRenderer{w: w, indent: "  "}
}

func (tr *TextRenderer) text(val string) {
	tr.newLine = false

	if !tr.lineIndented {
		for i := 0; i < tr.indentLvl; i++ {
			_, tr.err = io.WriteString(tr.w, tr.indent)
			if tr.err != nil {
				return
			}
		}

		tr.lineIndented = true
	}

	_, tr.err = io.WriteString(tr.w, val)
}

func (tr *TextRenderer) Keyword(val string) {
	if tr.err != nil {
		return
	}

	tr.text(val)
}

func (tr *TextRenderer) Identifier(val string) {
	if tr.err != nil {
		return
	}

	tr.text(val)
}

func (tr *TextRenderer) Symbol(val string) {
	if tr.err != nil {
		return
	}

	tr.text(val)
}

func (tr *TextRenderer) Constant(val string) {
	if tr.err != nil {
		return
	}

	tr.text(val)
}

func (tr *TextRenderer) Space() {
	if tr.err != nil {
		return
	}

	_, tr.err = io.WriteString(tr.w, " ")
}

func (tr *TextRenderer) NewLine() {
	if tr.err != nil {
		return
	}

	if tr.newLine {
		return
	}

	tr.newLine = true

	_, tr.err = io.WriteString(tr.w, "\n")
	if tr.err != nil {
		return
	}

	tr.lineIndented = false
}

func (tr *TextRenderer) Indent() {
	tr.indentLvl = tr.indentLvl + 1
}

func (tr *TextRenderer) Unindent() {
	tr.indentLvl = tr.indentLvl - 1
}
