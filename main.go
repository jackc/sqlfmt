//go:generate -command yacc go tool yacc
//go:generate yacc -o sql.go -p "sql" sql.y
package main

import (
	"io/ioutil"
	"os"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}
	lexer := NewSqlLexer(string(input))
	sqlParse(lexer)

	r := NewTextRenderer(os.Stdout)
	lexer.stmt.RenderTo(r)
}

type ColumnRef struct {
	Table  string
	Column string
}

type SelectExpr struct {
	Expr  ColumnRef
	Alias string
}

type SelectStmt struct {
	Fields    []SelectExpr
	FromTable string
}

func (s SelectStmt) RenderTo(r Renderer) {
	r.Text("select", "keyword")
	r.NewLine()
	r.Indent()
	for i, f := range s.Fields {
		if f.Expr.Table != "" {
			r.Text(f.Expr.Table, "identifier")
			r.Text(".", "period")
		}
		r.Text(f.Expr.Column, "identifier")

		if f.Alias != "" {
			r.Text(" ", "space")
			r.Text("as", "keyword")
			r.Text(" ", "space")
			r.Text(f.Alias, "identifier")
		}
		if i < len(s.Fields)-1 {
			r.Text(",", "comma")
		}
		r.NewLine()
	}
	r.Unindent()

	if s.FromTable != "" {
		r.Text("from", "keyword")
		r.NewLine()
		r.Indent()
		r.Text(s.FromTable, "identifier")
		r.NewLine()
	}
}
