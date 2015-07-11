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

type PgType struct {
	Name string
}

func (t PgType) RenderTo(r Renderer) {
	r.Text(t.Name, "typename")
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

type UnaryExpr struct {
	Operator string
	Expr     Expr
}

func (e UnaryExpr) RenderTo(r Renderer) {
	r.Text(e.Operator, "operator")
	e.Expr.RenderTo(r)
}

type WhenClause struct {
	When Expr
	Then Expr
}

func (w WhenClause) RenderTo(r Renderer) {
	r.Text("when", "keyword")
	r.Text(" ", "space")
	w.When.RenderTo(r)
	r.Text(" ", "space")
	r.Text("then", "keyword")
	r.NewLine()
	r.Indent()
	w.Then.RenderTo(r)
	r.NewLine()
	r.Unindent()
}

type CaseExpr struct {
	CaseArg     Expr
	WhenClauses []WhenClause
	Default     Expr
}

func (c CaseExpr) RenderTo(r Renderer) {
	r.Text("case", "keyword")

	if c.CaseArg != nil {
		r.Text(" ", "space")
		c.CaseArg.RenderTo(r)
	}

	r.NewLine()

	for _, w := range c.WhenClauses {
		w.RenderTo(r)
	}

	if c.Default != nil {
		r.Text("else", "keyword")
		r.NewLine()
		r.Indent()
		c.Default.RenderTo(r)
		r.NewLine()
		r.Unindent()
	}

	r.Text("end", "keyword")
	r.NewLine()
}

type ParenExpr struct {
	Expr Expr
}

func (e ParenExpr) RenderTo(r Renderer) {
	r.Text("(", "lparen")
	e.Expr.RenderTo(r)
	r.Text(")", "rparen")
}

type TypecastExpr struct {
	Expr     Expr
	Typename PgType
}

func (t TypecastExpr) RenderTo(r Renderer) {
	t.Expr.RenderTo(r)
	r.Text("::", "typecast")
	t.Typename.RenderTo(r)
}

type NotExpr struct {
	Expr Expr
}

func (e NotExpr) RenderTo(r Renderer) {
	r.Text("not", "keyword")
	r.Text(" ", "space")
	e.Expr.RenderTo(r)
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
	On    Expr
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

	if s.On != nil {
		r.Text(" ", "space")
		r.Text("on", "keyword")
		r.Text(" ", "space")
		s.On.RenderTo(r)
	}
}

type WhereClause struct {
	Expr Expr
}

func (e WhereClause) RenderTo(r Renderer) {
	r.Text("where", "keyword")
	r.NewLine()
	r.Indent()
	e.Expr.RenderTo(r)
	r.NewLine()
	r.Unindent()
}

type OrderExpr struct {
	Expr  Expr
	Order string
	Using string
	Nulls string
}

func (e OrderExpr) RenderTo(r Renderer) {
	e.Expr.RenderTo(r)
	if e.Order != "" {
		r.Text(" ", "space")
		r.Text(e.Order, "keyword")
	}
	if e.Using != "" {
		r.Text(" ", "space")
		r.Text("using", "keyword")
		r.Text(" ", "space")
		r.Text(e.Using, "operator")
	}
	if e.Nulls != "" {
		r.Text(" ", "space")
		r.Text("nulls", "keyword")
		r.Text(" ", "space")
		r.Text(e.Nulls, "keyword")
	}
}

type OrderClause struct {
	Exprs []OrderExpr
}

func (e OrderClause) RenderTo(r Renderer) {
	r.Text("order by", "keyword")
	r.NewLine()
	r.Indent()

	for i, f := range e.Exprs {
		f.RenderTo(r)
		if i < len(e.Exprs)-1 {
			r.Text(",", "comma")
		}
		r.NewLine()
	}
	r.Unindent()
}

type GroupByClause struct {
	Exprs []Expr
}

func (e GroupByClause) RenderTo(r Renderer) {
	r.Text("group by", "keyword")
	r.NewLine()
	r.Indent()

	for i, f := range e.Exprs {
		f.RenderTo(r)
		if i < len(e.Exprs)-1 {
			r.Text(",", "comma")
		}
		r.NewLine()
	}
	r.Unindent()
}

type LimitClause struct {
	Limit  Expr
	Offset Expr
}

func (e LimitClause) RenderTo(r Renderer) {
	if e.Limit != nil {
		r.Text("limit", "keyword")
		r.Text(" ", "space")
		e.Limit.RenderTo(r)
		r.NewLine()
	}
	if e.Offset != nil {
		r.Text("offset", "keyword")
		r.Text(" ", "space")
		e.Offset.RenderTo(r)
		r.NewLine()
	}
}

type LockingItem struct {
	Strength   string
	LockedRels []string
	WaitPolicy string
}

func (li LockingItem) RenderTo(r Renderer) {
	r.Text("for", "keyword")
	r.Text(" ", "space")
	r.Text(li.Strength, "keyword")

	if li.LockedRels != nil {
		r.Text(" ", "space")
		r.Text("of", "keyword")
		r.Text(" ", "space")

		for i, lr := range li.LockedRels {
			r.Text(lr, "identifier")
			if i < len(li.LockedRels)-1 {
				r.Text(",", "comma")
				r.Text(" ", "space")
			}
		}
	}

	if li.WaitPolicy != "" {
		r.Text(" ", "space")
		r.Text(li.WaitPolicy, "keyword")
	}

	r.NewLine()
}

type LockingClause struct {
	Locks []LockingItem
}

func (lc LockingClause) RenderTo(r Renderer) {
	for _, li := range lc.Locks {
		li.RenderTo(r)
	}
}

type SelectStmt struct {
	DistinctList  []Expr
	TargetList    []Expr
	FromClause    *FromClause
	WhereClause   *WhereClause
	OrderClause   *OrderClause
	GroupByClause *GroupByClause
	HavingClause  Expr
	LimitClause   *LimitClause
	LockingClause *LockingClause
	ParenWrapped  bool
}

func (s SelectStmt) RenderTo(r Renderer) {
	if s.ParenWrapped {
		r.Text("(", "lparen")
	}

	r.Text("select", "keyword")

	if s.DistinctList != nil {
		r.Text(" ", "space")
		r.Text("distinct", "keyword")

		if len(s.DistinctList) > 0 {
			r.Text(" ", "space")
			r.Text("on", "keyword")
			r.Text("(", "lparen")

			for i, f := range s.DistinctList {
				f.RenderTo(r)
				if i < len(s.DistinctList)-1 {
					r.Text(",", "comma")
					r.Text(" ", "space")
				}
			}
			r.Text(")", "rparen")
		}

	}

	r.NewLine()
	r.Indent()
	for i, f := range s.TargetList {
		f.RenderTo(r)
		if i < len(s.TargetList)-1 {
			r.Text(",", "comma")
		}
		r.NewLine()
	}
	r.Unindent()

	if s.FromClause != nil {
		s.FromClause.RenderTo(r)
	}

	if s.WhereClause != nil {
		s.WhereClause.RenderTo(r)
	}

	if s.GroupByClause != nil {
		s.GroupByClause.RenderTo(r)
	}

	if s.HavingClause != nil {
		r.Text("having", "keyword")
		r.NewLine()
		r.Indent()
		s.HavingClause.RenderTo(r)
		r.NewLine()
	}

	if s.OrderClause != nil {
		s.OrderClause.RenderTo(r)
	}

	if s.LimitClause != nil {
		s.LimitClause.RenderTo(r)
	}

	if s.LockingClause != nil {
		s.LockingClause.RenderTo(r)
	}

	if s.ParenWrapped {
		r.Text(")", "rparen")
		r.NewLine()
	}
}
