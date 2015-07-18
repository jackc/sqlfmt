//go:generate -command yacc go tool yacc
//go:generate yacc -o sql.go -p "sql" sql.y
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const Version = "0.0.1"

var options struct {
	version bool
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage:  %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&options.version, "version", false, "print version and exit")
	flag.Parse()

	if options.version {
		fmt.Printf("sqlfmt v%v\n", Version)
		os.Exit(0)
	}

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

type QualifiedName []string

func (qn QualifiedName) RenderTo(r Renderer) {
	for i, s := range qn {
		r.Text(s, "identifier")
		if i+1 < len(qn) {
			r.Text(".", "period")
		}
	}
}

type StringLiteral string

func (s StringLiteral) RenderTo(r Renderer) {
	r.Text(string(s), "stringLiteral")
}

type IntegerLiteral string

func (s IntegerLiteral) RenderTo(r Renderer) {
	r.Text(string(s), "integerLiteral")
}

type BoolLiteral bool

func (b BoolLiteral) RenderTo(r Renderer) {
	if b {
		r.Text("true", "boolLiteral")
	} else {
		r.Text("false", "boolLiteral")
	}
}

type BooleanExpr struct {
	Left     Expr
	Operator string
	Right    Expr
}

func (e BooleanExpr) RenderTo(r Renderer) {
	e.Left.RenderTo(r)
	r.NewLine()
	r.Text(e.Operator, "operator")
	r.Space()
	e.Right.RenderTo(r)
}

type BinaryExpr struct {
	Left     Expr
	Operator string
	Right    Expr
}

func (e BinaryExpr) RenderTo(r Renderer) {
	e.Left.RenderTo(r)
	r.Space()
	r.Text(e.Operator, "operator")
	r.Space()
	e.Right.RenderTo(r)
}

type TextOpWithEscapeExpr struct {
	Left     Expr
	Operator string
	Right    Expr
	Escape   Expr
}

func (e TextOpWithEscapeExpr) RenderTo(r Renderer) {
	e.Left.RenderTo(r)
	r.Space()
	r.Text(e.Operator, "operator")
	r.Space()
	e.Right.RenderTo(r)

	if e.Escape != nil {
		r.Space()
		r.Text("escape", "keyword")
		r.Space()
		e.Escape.RenderTo(r)
	}
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
	r.Space()
	w.When.RenderTo(r)
	r.Space()
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
		r.Space()
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

type CollateExpr struct {
	Expr      Expr
	Collation QualifiedName
}

func (c CollateExpr) RenderTo(r Renderer) {
	c.Expr.RenderTo(r)
	r.Space()
	r.Text("collate", "keyword")
	r.Space()
	c.Collation.RenderTo(r)
}

type NotExpr struct {
	Expr Expr
}

func (e NotExpr) RenderTo(r Renderer) {
	r.Text("not", "keyword")
	r.Space()
	e.Expr.RenderTo(r)
}

type IsNullExpr struct {
	Expr Expr
	Not  bool
}

func (e IsNullExpr) RenderTo(r Renderer) {
	e.Expr.RenderTo(r)
	r.Space()
	r.Text("is", "keyword")
	r.Space()
	if e.Not {
		r.Text("not", "keyword")
		r.Space()
	}
	r.Text("null", "keyword")
}

type IsBoolOpExpr struct {
	Expr Expr
	Not  bool
	Op   string
}

func (e IsBoolOpExpr) RenderTo(r Renderer) {
	e.Expr.RenderTo(r)
	r.Space()
	r.Text("is", "keyword")
	r.Space()
	if e.Not {
		r.Text("not", "keyword")
		r.Space()
	}
	r.Text(e.Op, "keyword")
}

type AliasedExpr struct {
	Expr  Expr
	Alias string
}

func (e AliasedExpr) RenderTo(r Renderer) {
	e.Expr.RenderTo(r)
	r.Space()
	r.Text("as", "keyword")
	r.Space()
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
		r.Space()
	}

	s.Right.RenderTo(r)

	if len(s.Using) > 0 {
		r.Space()
		r.Text("using", "keyword")
		r.Text("(", "lparen")

		for i, u := range s.Using {
			r.Text(u, "identifier")
			if i+1 < len(s.Using) {
				r.Text(",", "comma")
				r.Space()
			}
		}

		r.Text(")", "rparen")
	}

	if s.On != nil {
		r.Space()
		r.Text("on", "keyword")
		r.Space()
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
		r.Space()
		r.Text(e.Order, "keyword")
	}
	if e.Using != "" {
		r.Space()
		r.Text("using", "keyword")
		r.Space()
		r.Text(e.Using, "operator")
	}
	if e.Nulls != "" {
		r.Space()
		r.Text("nulls", "keyword")
		r.Space()
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
		r.Space()
		e.Limit.RenderTo(r)
		r.NewLine()
	}
	if e.Offset != nil {
		r.Text("offset", "keyword")
		r.Space()
		e.Offset.RenderTo(r)
		r.NewLine()
	}
}

type AtTimeZoneExpr struct {
	Expr     Expr
	TimeZone Expr
}

func (e AtTimeZoneExpr) RenderTo(r Renderer) {
	e.Expr.RenderTo(r)
	r.Space()
	r.Text("at time zone", "keyword")
	r.Space()
	e.TimeZone.RenderTo(r)
}

type LockingItem struct {
	Strength   string
	LockedRels []string
	WaitPolicy string
}

func (li LockingItem) RenderTo(r Renderer) {
	r.Text("for", "keyword")
	r.Space()
	r.Text(li.Strength, "keyword")

	if li.LockedRels != nil {
		r.Space()
		r.Text("of", "keyword")
		r.Space()

		for i, lr := range li.LockedRels {
			r.Text(lr, "identifier")
			if i < len(li.LockedRels)-1 {
				r.Text(",", "comma")
				r.Space()
			}
		}
	}

	if li.WaitPolicy != "" {
		r.Space()
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

type FuncApplication struct {
	Name string
}

func (fa FuncApplication) RenderTo(r Renderer) {
	r.Text(fa.Name, "identifier")
	r.Text("(", "lparen")
	r.Text(")", "lparen")
}

type DefaultExpr bool

func (d DefaultExpr) RenderTo(r Renderer) {
	r.Text("default", "keyword")
}

type ValuesRow []Expr

func (vr ValuesRow) RenderTo(r Renderer) {
	r.Text("(", "lparen")

	for i, e := range vr {
		e.RenderTo(r)
		if i < len(vr)-1 {
			r.Text(",", "comma")
			r.Space()
		}
	}

	r.Text(")", "rparen")
}

type ValuesClause []ValuesRow

func (vc ValuesClause) RenderTo(r Renderer) {
	r.Text("values", "keyword")
	r.NewLine()
	r.Indent()

	for i, row := range vc {
		row.RenderTo(r)
		if i < len(vc)-1 {
			r.Text(",", "comma")
		}
		r.NewLine()
	}

	r.Unindent()
}

type SimpleSelect struct {
	DistinctList  []Expr
	TargetList    []Expr
	FromClause    *FromClause
	WhereClause   *WhereClause
	GroupByClause *GroupByClause
	HavingClause  Expr

	ValuesClause ValuesClause

	LeftSelect  *SelectStmt
	SetOp       string
	SetAll      bool
	RightSelect *SelectStmt
}

func (s SimpleSelect) RenderTo(r Renderer) {
	if s.ValuesClause != nil {
		s.ValuesClause.RenderTo(r)
		return
	}

	if s.LeftSelect != nil {
		s.LeftSelect.RenderTo(r)
		r.NewLine()
		r.Text(s.SetOp, "keyword")

		if s.SetAll {
			r.Space()
			r.Text("all", "keyword")
		}

		r.NewLine()

		s.RightSelect.RenderTo(r)

		return
	}

	r.Text("select", "keyword")

	if s.DistinctList != nil {
		r.Space()
		r.Text("distinct", "keyword")

		if len(s.DistinctList) > 0 {
			r.Space()
			r.Text("on", "keyword")
			r.Text("(", "lparen")

			for i, f := range s.DistinctList {
				f.RenderTo(r)
				if i < len(s.DistinctList)-1 {
					r.Text(",", "comma")
					r.Space()
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
}

type SelectStmt struct {
	SimpleSelect
	OrderClause   *OrderClause
	LimitClause   *LimitClause
	LockingClause *LockingClause

	ParenWrapped bool
}

func (s SelectStmt) RenderTo(r Renderer) {
	if s.ParenWrapped {
		r.Text("(", "lparen")
	}

	s.SimpleSelect.RenderTo(r)

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
