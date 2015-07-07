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
	if rc := sqlParse(lexer); rc != 0 {
		os.Exit(1)
	}

	r := NewTextRenderer(os.Stdout)
	lexer.stmt.RenderTo(r)
}

type Expr interface {
	RenderTo(Renderer)
}

type ColumnRef struct {
	Table  string
	Column string
}

func (cr ColumnRef) RenderTo(r Renderer) {
	if cr.Table != "" {
		r.Text(cr.Table, "identifier")
		r.Text(".", "period")
	}
	r.Text(cr.Column, "identifier")
}

type StringLiteral string

func (s StringLiteral) RenderTo(r Renderer) {
	r.Text(string(s), "stringLiteral")
}

type IntegerLiteral string

func (s IntegerLiteral) RenderTo(r Renderer) {
	r.Text(string(s), "integerLiteral")
}

type BinaryExpr struct {
	Left     Expr
	Operator string
	Right    Expr
}

func (e BinaryExpr) RenderTo(r Renderer) {
	e.Left.RenderTo(r)
	r.Text(" ", "space")
	r.Text(e.Operator, "operator")
	r.Text(" ", "space")
	e.Right.RenderTo(r)
}

type ParenExpr struct {
	Expr Expr
}

func (e ParenExpr) RenderTo(r Renderer) {
	r.Text("(", "lparen")
	e.Expr.RenderTo(r)
	r.Text(")", "rparen")
}

type AliasedExpr struct {
	Expr  Expr
	Alias string
}

func (e AliasedExpr) RenderTo(r Renderer) {
	e.Expr.RenderTo(r)
	r.Text(" ", "space")
	r.Text("as", "keyword")
	r.Text(" ", "space")
	r.Text(e.Alias, "identifier")
}

type FromClause struct {
	Expr Expr
}

func (e FromClause) RenderTo(r Renderer) {
	r.Text("from", "keyword")
	r.NewLine()
	r.Indent()
	e.Expr.RenderTo(r)
	r.NewLine()
	r.Unindent()
}

type JoinExpr struct {
	Left  Expr
	Join  string
	Right Expr
	Using []string
}

func (s JoinExpr) RenderTo(r Renderer) {
	s.Left.RenderTo(r)

	if s.Join == "," {
		r.Text(",", "comma")
		r.NewLine()
	} else {
		r.NewLine()
		r.Text(s.Join, "keyword")
		r.Text(" ", "space")
	}

	s.Right.RenderTo(r)

	if len(s.Using) > 0 {
		r.Text(" ", "space")
		r.Text("using", "keyword")
		r.Text("(", "lparen")

		for i, u := range s.Using {
			r.Text(u, "identifier")
			if i+1 < len(s.Using) {
				r.Text(",", "comma")
				r.Text(" ", "space")
			}
		}

		r.Text(")", "rparen")
	}
}

type SelectStmt struct {
	Fields     []Expr
	FromClause *FromClause
}

func (s SelectStmt) RenderTo(r Renderer) {
	r.Text("select", "keyword")
	r.NewLine()
	r.Indent()
	for i, f := range s.Fields {
		f.RenderTo(r)
		if i < len(s.Fields)-1 {
			r.Text(",", "comma")
		}
		r.NewLine()
	}
	r.Unindent()

	if s.FromClause != nil {
		s.FromClause.RenderTo(r)
	}
}
