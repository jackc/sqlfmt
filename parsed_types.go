package sqlfmt

import (
	"errors"
)

func Parse(lexer *sqlLex) (stmt *SelectStmt, err error) {
	if rc := yyParse(lexer); rc != 0 {
		return nil, errors.New("Parse failed")
	}

	return lexer.stmt, nil
}

type Expr interface {
	RenderTo(Renderer)
}

type PgType struct {
	Name         AnyName
	OptInterval  *OptInterval
	Setof        bool
	ArrayWord    bool
	ArrayBounds  []IntegerConst
	TypeMods     []Expr
	CharSet      string
	WithTimeZone bool
}

func (t PgType) RenderTo(r Renderer) {
	if t.Setof {
		r.Keyword("setof")
	}

	t.Name.RenderTo(r)

	if t.OptInterval != nil {
		t.OptInterval.RenderTo(r)
	}

	if t.ArrayWord {
		r.Keyword("array")
	}

	for _, ab := range t.ArrayBounds {
		r.Symbol("[")
		r.Constant(string(ab))
		r.Symbol("]")
	}

	if len(t.TypeMods) > 0 {
		r.Symbol("(")
		for i, e := range t.TypeMods {
			e.RenderTo(r)
			if i < len(t.TypeMods)-1 {
				r.Symbol(",")
			}
		}
		r.Symbol(")")
	}

	if t.WithTimeZone {
		r.Keyword("with time zone")
	}

	if t.CharSet != "" {
		r.Keyword("character set")
		r.Identifier(t.CharSet)
	}
}

type AnyName []string

func (an AnyName) RenderTo(r Renderer) {
	for i, n := range an {
		r.Identifier(n)
		if i < len(an)-1 {
			r.Symbol(".")
		}
	}
}

type ColumnRef struct {
	Name        string
	Indirection Indirection
}

func (cr ColumnRef) RenderTo(r Renderer) {
	r.Identifier(cr.Name)
	if cr.Indirection != nil {
		cr.Indirection.RenderTo(r)
	}
}

type Indirection []IndirectionEl

func (i Indirection) RenderTo(r Renderer) {
	for _, e := range i {
		e.RenderTo(r)
	}
}

type IndirectionEl struct {
	Name           string
	LowerSubscript Expr
	UpperSubscript Expr
}

func (ie IndirectionEl) RenderTo(r Renderer) {
	if ie.LowerSubscript != nil {
		r.Symbol("[")
		ie.LowerSubscript.RenderTo(r)
		if ie.UpperSubscript != nil {
			r.Symbol(":")
			ie.UpperSubscript.RenderTo(r)
		}
		r.Symbol("]")
	} else {
		r.Symbol(".")
		r.Identifier(ie.Name)
	}
}

type StringConst string

func (s StringConst) RenderTo(r Renderer) {
	r.Constant(string(s))
}

type IntegerConst string

func (s IntegerConst) RenderTo(r Renderer) {
	r.Constant(string(s))
}

type FloatConst string

func (s FloatConst) RenderTo(r Renderer) {
	r.Constant(string(s))
}

type BoolConst bool

func (b BoolConst) RenderTo(r Renderer) {
	if b {
		r.Keyword("true")
	} else {
		r.Keyword("false")
	}
}

type NullConst struct{}

func (n NullConst) RenderTo(r Renderer) {
	r.Keyword("null")
}

type BitConst string

func (b BitConst) RenderTo(r Renderer) {
	r.Constant(string(b))
}

type BooleanExpr struct {
	Left     Expr
	Operator string
	Right    Expr
}

func (e BooleanExpr) RenderTo(r Renderer) {
	e.Left.RenderTo(r)
	r.NewLine()
	r.Symbol(e.Operator)
	e.Right.RenderTo(r)
}

type BinaryExpr struct {
	Left     Expr
	Operator AnyName
	Right    Expr
}

func (e BinaryExpr) RenderTo(r Renderer) {
	e.Left.RenderTo(r)
	r.Space()
	e.Operator.RenderTo(r)
	r.Space()
	e.Right.RenderTo(r)
}

type ArrayConstructorExpr ArrayExpr

func (ace ArrayConstructorExpr) RenderTo(r Renderer) {
	r.Keyword("array")
	ArrayExpr(ace).RenderTo(r)
}

type ArrayExpr []Expr

func (a ArrayExpr) RenderTo(r Renderer) {
	r.Symbol("[")
	for i, e := range a {
		e.RenderTo(r)
		if i < len(a)-1 {
			r.Symbol(",")
		}
	}
	r.Symbol("]")
}

type TextOpWithEscapeExpr struct {
	Left     Expr
	Operator string
	Right    Expr
	Escape   Expr
}

func (e TextOpWithEscapeExpr) RenderTo(r Renderer) {
	e.Left.RenderTo(r)
	r.Symbol(e.Operator)
	e.Right.RenderTo(r)

	if e.Escape != nil {
		r.Keyword("escape")
		e.Escape.RenderTo(r)
	}
}

type UnaryExpr struct {
	Operator AnyName
	Expr     Expr
}

func (e UnaryExpr) RenderTo(r Renderer) {
	e.Operator.RenderTo(r)
	r.RefuseSpace()
	e.Expr.RenderTo(r)
}

type PostfixExpr struct {
	Expr     Expr
	Operator AnyName
}

func (e PostfixExpr) RenderTo(r Renderer) {
	e.Expr.RenderTo(r)
	e.Operator.RenderTo(r)
}

type SubqueryOpExpr struct {
	Value Expr
	Op    SubqueryOp
	Type  string
	Query Expr
}

func (s SubqueryOpExpr) RenderTo(r Renderer) {
	s.Value.RenderTo(r)
	s.Op.RenderTo(r)
	r.Keyword(s.Type)
	r.Space()
	s.Query.RenderTo(r)
}

type SubqueryOp struct {
	Operator bool
	Name     AnyName
}

func (s SubqueryOp) RenderTo(r Renderer) {
	if s.Operator {
		r.Keyword("operator")
		r.Symbol("(")
	}
	s.Name.RenderTo(r)
	if s.Operator {
		r.Symbol(")")
	}
}

type WhenClause struct {
	When Expr
	Then Expr
}

func (w WhenClause) RenderTo(r Renderer) {
	r.Keyword("when")
	w.When.RenderTo(r)
	r.Keyword("then")
	r.NewLine()
	r.Indent()
	w.Then.RenderTo(r)
	r.NewLine()
	r.Unindent()
}

type InExpr struct {
	Value Expr
	Not   bool
	In    Expr
}

func (i InExpr) RenderTo(r Renderer) {
	i.Value.RenderTo(r)

	if i.Not {
		r.Keyword("not")
	}

	r.Keyword("in")
	r.Space()

	i.In.RenderTo(r)
}

type BetweenExpr struct {
	Expr      Expr
	Not       bool
	Symmetric bool
	Left      Expr
	Right     Expr
}

func (b BetweenExpr) RenderTo(r Renderer) {
	b.Expr.RenderTo(r)

	if b.Not {
		r.Keyword("not")
	}

	r.Keyword("between")

	if b.Symmetric {
		r.Keyword("symmetric")
	}

	b.Left.RenderTo(r)
	r.Keyword("and")
	b.Right.RenderTo(r)
}

type CaseExpr struct {
	CaseArg     Expr
	WhenClauses []WhenClause
	Default     Expr
}

func (c CaseExpr) RenderTo(r Renderer) {
	r.Keyword("case")

	if c.CaseArg != nil {
		c.CaseArg.RenderTo(r)
	}

	r.NewLine()

	for _, w := range c.WhenClauses {
		w.RenderTo(r)
	}

	if c.Default != nil {
		r.Keyword("else")
		r.NewLine()
		r.Indent()
		c.Default.RenderTo(r)
		r.NewLine()
		r.Unindent()
	}

	r.Keyword("end")
	r.NewLine()
}

type ParenExpr struct {
	Expr        Expr
	Indirection Indirection
}

func (e ParenExpr) RenderTo(r Renderer) {
	r.Symbol("(")
	e.Expr.RenderTo(r)
	r.Symbol(")")
	if e.Indirection != nil {
		e.Indirection.RenderTo(r)
	}
}

type TypecastExpr struct {
	Expr     Expr
	Typename PgType
}

func (t TypecastExpr) RenderTo(r Renderer) {
	t.Expr.RenderTo(r)
	r.Symbol("::")
	t.Typename.RenderTo(r)
}

type ConstTypeExpr struct {
	Typename PgType
	Expr     Expr
}

func (t ConstTypeExpr) RenderTo(r Renderer) {
	t.Typename.RenderTo(r)
	t.Expr.RenderTo(r)
}

type ConstIntervalExpr struct {
	Precision   IntegerConst
	Value       Expr
	OptInterval *OptInterval
}

func (i ConstIntervalExpr) RenderTo(r Renderer) {
	r.Keyword("interval")
	if i.Precision != "" {
		r.Symbol("(")
		i.Precision.RenderTo(r)
		r.Symbol(")")
	}

	i.Value.RenderTo(r)

	if i.OptInterval != nil {
		i.OptInterval.RenderTo(r)
	}
}

type OptInterval struct {
	Left   string
	Right  string
	Second *IntervalSecond
}

func (oi OptInterval) RenderTo(r Renderer) {
	if oi.Left != "" {
		r.Keyword(oi.Left)
	}

	if oi.Right != "" {
		r.Keyword("to")
		r.Keyword(oi.Right)
	}

	if oi.Second != nil {
		oi.Second.RenderTo(r)
	}
}

type IntervalSecond struct {
	Precision IntegerConst
}

func (is IntervalSecond) RenderTo(r Renderer) {
	r.Keyword("second")
	if is.Precision != "" {
		r.Symbol("(")
		is.Precision.RenderTo(r)
		r.Symbol(")")
	}
}

type ExtractExpr ExtractList

func (ee ExtractExpr) RenderTo(r Renderer) {
	r.Keyword("extract")
	r.Symbol("(")
	ExtractList(ee).RenderTo(r)
	r.Symbol(")")
}

type ExtractList struct {
	Extract Expr
	Time    Expr
}

func (el ExtractList) RenderTo(r Renderer) {
	el.Extract.RenderTo(r)
	r.Keyword("from")
	el.Time.RenderTo(r)
}

type OverlayExpr OverlayList

func (oe OverlayExpr) RenderTo(r Renderer) {
	r.Keyword("overlay")
	r.Symbol("(")
	OverlayList(oe).RenderTo(r)
	r.Symbol(")")
}

type OverlayList struct {
	Dest    Expr
	Placing Expr
	From    Expr
	For     Expr
}

func (ol OverlayList) RenderTo(r Renderer) {
	ol.Dest.RenderTo(r)
	r.Keyword("placing")
	ol.Placing.RenderTo(r)
	r.Keyword("from")
	ol.From.RenderTo(r)

	if ol.For != nil {
		r.Keyword("for")
		ol.For.RenderTo(r)
	}
}

type PositionExpr PositionList

func (pe PositionExpr) RenderTo(r Renderer) {
	r.Keyword("position")
	r.Symbol("(")
	PositionList(pe).RenderTo(r)
	r.Symbol(")")
}

type PositionList struct {
	Substring Expr
	String    Expr
}

func (pl PositionList) RenderTo(r Renderer) {
	pl.Substring.RenderTo(r)
	r.Keyword("in")
	pl.String.RenderTo(r)
}

type SubstrExpr SubstrList

func (se SubstrExpr) RenderTo(r Renderer) {
	r.Keyword("substring")
	r.Symbol("(")
	SubstrList(se).RenderTo(r)
	r.Symbol(")")
}

type SubstrList struct {
	Source Expr
	From   Expr
	For    Expr
}

func (sl SubstrList) RenderTo(r Renderer) {
	sl.Source.RenderTo(r)
	r.Keyword("from")
	sl.From.RenderTo(r)

	if sl.For != nil {
		r.Keyword("for")
		sl.For.RenderTo(r)
	}
}

type TrimExpr struct {
	Direction string
	TrimList
}

func (te TrimExpr) RenderTo(r Renderer) {
	r.Keyword("trim")
	r.Symbol("(")
	if te.Direction != "" {
		r.Keyword(te.Direction)
	}
	te.TrimList.RenderTo(r)
	r.Symbol(")")
}

type TrimList struct {
	Left  Expr
	From  bool
	Right []Expr
}

func (tl TrimList) RenderTo(r Renderer) {
	if tl.Left != nil {
		tl.Left.RenderTo(r)
	}

	if tl.From {
		r.Keyword("from")
	}

	for i, e := range tl.Right {
		e.RenderTo(r)
		if i+1 < len(tl.Right) {
			r.Symbol(",")
		}
	}
}

type XmlElement struct {
	Name       string
	Attributes XmlAttributes
	Body       []Expr
}

func (el XmlElement) RenderTo(r Renderer) {
	r.Keyword("xmlelement")
	r.Symbol("(")
	r.Keyword("name")
	r.Identifier(el.Name)

	if el.Attributes != nil {
		r.Symbol(",")
		el.Attributes.RenderTo(r)
	}

	if el.Body != nil {
		for _, e := range el.Body {
			r.Symbol(",")
			e.RenderTo(r)
		}
	}

	r.Symbol(")")
}

type XmlAttributes []XmlAttributeEl

func (attrs XmlAttributes) RenderTo(r Renderer) {
	r.Keyword("xmlattributes")
	r.Symbol("(")
	xmlAttributes(attrs).RenderTo(r)
	r.Symbol(")")
}

type XmlAttributeEl struct {
	Value Expr
	Name  string
}

func (el XmlAttributeEl) RenderTo(r Renderer) {
	el.Value.RenderTo(r)
	if el.Name != "" {
		r.Keyword("as")
		r.Identifier(el.Name)
	}
}

type XmlExists struct {
	Path Expr
	Body XmlExistsArgument
}

func (e XmlExists) RenderTo(r Renderer) {
	r.Keyword("xmlexists")
	r.Symbol("(")
	e.Path.RenderTo(r)
	e.Body.RenderTo(r)
	r.Symbol(")")
}

type XmlExistsArgument struct {
	LeftByRef  bool
	Arg        Expr
	RightByRef bool
}

func (a XmlExistsArgument) RenderTo(r Renderer) {
	r.Keyword("passing")

	if a.LeftByRef {
		r.Keyword("by ref")
	}

	a.Arg.RenderTo(r)

	if a.RightByRef {
		r.Keyword("by ref")
	}
}

type XmlForest []XmlAttributeEl

func (f XmlForest) RenderTo(r Renderer) {
	r.Keyword("xmlforest")
	r.Symbol("(")
	xmlAttributes(f).RenderTo(r)
	r.Symbol(")")
}

type xmlAttributes []XmlAttributeEl

func (attrs xmlAttributes) RenderTo(r Renderer) {
	for i, a := range attrs {
		a.RenderTo(r)
		if i+1 < len(attrs) {
			r.Symbol(",")
		}
	}
}

type XmlParse struct {
	Type             string
	Content          Expr
	WhitespaceOption string
}

func (p XmlParse) RenderTo(r Renderer) {
	r.Keyword("xmlparse")
	r.Symbol("(")
	r.Keyword(p.Type)
	p.Content.RenderTo(r)
	if p.WhitespaceOption != "" {
		r.Keyword(p.WhitespaceOption)
	}

	r.Symbol(")")
}

type XmlPi struct {
	Name    string
	Content Expr
}

func (p XmlPi) RenderTo(r Renderer) {
	r.Keyword("xmlpi")
	r.Symbol("(")
	r.Keyword("name")
	r.Identifier(p.Name)

	if p.Content != nil {
		r.Symbol(",")
		p.Content.RenderTo(r)
	}
	r.Symbol(")")
}

type XmlRoot struct {
	Xml        Expr
	Version    XmlRootVersion
	Standalone string
}

func (x XmlRoot) RenderTo(r Renderer) {
	r.Keyword("xmlroot")
	r.Symbol("(")
	x.Xml.RenderTo(r)
	r.Symbol(",")
	x.Version.RenderTo(r)
	if x.Standalone != "" {
		r.Symbol(",")
		r.Keyword("standalone")
		r.Keyword(x.Standalone)
	}
	r.Symbol(")")
}

type XmlRootVersion struct {
	Expr Expr
}

func (rv XmlRootVersion) RenderTo(r Renderer) {
	r.Keyword("version")
	if rv.Expr != nil {
		rv.Expr.RenderTo(r)
	} else {
		r.Keyword("no value")
	}
}

type XmlSerialize struct {
	XmlType string
	Content Expr
	Type    PgType
}

func (s XmlSerialize) RenderTo(r Renderer) {
	r.Keyword("xmlserialize")
	r.Symbol("(")
	r.Keyword(s.XmlType)
	s.Content.RenderTo(r)
	r.Keyword("as")
	s.Type.RenderTo(r)
	r.Symbol(")")
}

type CollateExpr struct {
	Expr      Expr
	Collation AnyName
}

func (c CollateExpr) RenderTo(r Renderer) {
	c.Expr.RenderTo(r)
	r.Keyword("collate")
	c.Collation.RenderTo(r)
}

type NotExpr struct {
	Expr Expr
}

func (e NotExpr) RenderTo(r Renderer) {
	r.Keyword("not")
	e.Expr.RenderTo(r)
}

type IsExpr struct {
	Expr Expr
	Not  bool
	Op   string // null, document, true, false, etc.
}

func (e IsExpr) RenderTo(r Renderer) {
	e.Expr.RenderTo(r)
	r.Keyword("is")
	if e.Not {
		r.Keyword("not")
	}
	r.Keyword(e.Op)
}

type AliasedExpr struct {
	Expr  Expr
	Alias string
}

func (e AliasedExpr) RenderTo(r Renderer) {
	e.Expr.RenderTo(r)
	r.Keyword("as")
	r.Identifier(e.Alias)
}

type IntoClause struct {
	Options  string
	OptTable bool
	Target   AnyName
}

func (i IntoClause) RenderTo(r Renderer) {
	r.Keyword("into")

	if i.Options != "" {
		r.Keyword(i.Options)
	}

	if i.OptTable {
		r.Keyword("table")
	}

	i.Target.RenderTo(r)
	r.NewLine()
}

type FromClause struct {
	Expr Expr
}

func (e FromClause) RenderTo(r Renderer) {
	r.Keyword("from")
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
		r.Symbol(",")
		r.NewLine()
	} else {
		r.NewLine()
		r.Keyword(s.Join)
	}

	s.Right.RenderTo(r)

	if len(s.Using) > 0 {
		r.Keyword("using")
		r.Symbol("(")

		for i, u := range s.Using {
			r.Identifier(u)
			if i+1 < len(s.Using) {
				r.Symbol(",")
			}
		}

		r.Symbol(")")
	}

	if s.On != nil {
		r.Keyword("on")
		s.On.RenderTo(r)
	}
}

type WhereClause struct {
	Expr Expr
}

func (e WhereClause) RenderTo(r Renderer) {
	r.Keyword("where")
	r.NewLine()
	r.Indent()
	e.Expr.RenderTo(r)
	r.NewLine()
	r.Unindent()
}

type OrderExpr struct {
	Expr  Expr
	Order string
	Using AnyName
	Nulls string
}

func (e OrderExpr) RenderTo(r Renderer) {
	e.Expr.RenderTo(r)
	if e.Order != "" {
		r.Keyword(e.Order)
	}
	if len(e.Using) > 0 {
		r.Keyword("using")
		e.Using.RenderTo(r)
	}
	if e.Nulls != "" {
		r.Keyword("nulls")
		r.Keyword(e.Nulls)
	}
}

type OrderClause struct {
	Exprs []OrderExpr
}

func (e OrderClause) RenderTo(r Renderer) {
	r.Keyword("order by")
	r.NewLine()
	r.Indent()

	for i, f := range e.Exprs {
		f.RenderTo(r)
		if i < len(e.Exprs)-1 {
			r.Symbol(",")
		}
		r.NewLine()
	}
	r.Unindent()
}

type GroupByClause struct {
	Exprs []Expr
}

func (e GroupByClause) RenderTo(r Renderer) {
	r.Keyword("group by")
	r.NewLine()
	r.Indent()

	for i, f := range e.Exprs {
		f.RenderTo(r)
		if i < len(e.Exprs)-1 {
			r.Symbol(",")
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
		r.Keyword("limit")
		e.Limit.RenderTo(r)
		r.NewLine()
	}
	if e.Offset != nil {
		r.Keyword("offset")
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
	r.Keyword("at time zone")
	e.TimeZone.RenderTo(r)
}

type LockingItem struct {
	Strength   string
	LockedRels []AnyName
	WaitPolicy string
}

func (li LockingItem) RenderTo(r Renderer) {
	r.Keyword("for")
	r.Keyword(li.Strength)

	if li.LockedRels != nil {
		r.Keyword("of")

		for i, lr := range li.LockedRels {
			lr.RenderTo(r)
			if i < len(li.LockedRels)-1 {
				r.Symbol(",")
			}
		}
	}

	if li.WaitPolicy != "" {
		r.Keyword(li.WaitPolicy)
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

type FuncExprNoParens string

func (fe FuncExprNoParens) RenderTo(r Renderer) {
	r.Keyword(string(fe))
}

type FuncExpr struct {
	FuncApplication
	WithinGroupClause *WithinGroupClause
	FilterClause      *FilterClause
	OverClause        *OverClause
}

func (fe FuncExpr) RenderTo(r Renderer) {
	fe.FuncApplication.RenderTo(r)

	if fe.WithinGroupClause != nil {
		fe.WithinGroupClause.RenderTo(r)
	}

	if fe.FilterClause != nil {
		fe.FilterClause.RenderTo(r)
	}

	if fe.OverClause != nil {
		fe.OverClause.RenderTo(r)
	}
}

type FuncApplication struct {
	Name AnyName

	Distinct bool

	Star        bool
	Args        []FuncArg
	VariadicArg *FuncArg

	OrderClause *OrderClause
}

func (fa FuncApplication) RenderTo(r Renderer) {
	fa.Name.RenderTo(r)
	r.Symbol("(")

	if fa.Distinct {
		r.Keyword("distinct")
	}

	if fa.Star {
		r.Symbol("*")
	} else if len(fa.Args) > 0 {
		for i, a := range fa.Args {
			a.RenderTo(r)
			if i < len(fa.Args)-1 {
				r.Symbol(",")
			}
		}
	}

	if fa.VariadicArg != nil {
		if len(fa.Args) > 0 {
			r.Symbol(",")
		}

		r.Keyword("variadic")
		fa.VariadicArg.RenderTo(r)
	}

	if fa.OrderClause != nil {
		fa.OrderClause.RenderTo(r)
	}

	r.Symbol(")")
}

type FuncArg struct {
	Name   string
	NameOp string
	Expr   Expr
}

func (fa FuncArg) RenderTo(r Renderer) {
	if fa.Name != "" {
		r.Identifier(fa.Name)
		r.Symbol(fa.NameOp)
	}
	fa.Expr.RenderTo(r)
}

type CastFunc struct {
	Name string
	Expr Expr
	Type PgType
}

func (cf CastFunc) RenderTo(r Renderer) {
	r.Keyword(cf.Name)
	r.Symbol("(")
	cf.Expr.RenderTo(r)
	r.Keyword("as")
	cf.Type.RenderTo(r)
	r.Symbol(")")
}

type IsOfExpr struct {
	Expr  Expr
	Not   bool
	Types []PgType
}

func (io IsOfExpr) RenderTo(r Renderer) {
	io.Expr.RenderTo(r)
	r.Keyword("is")

	if io.Not {
		r.Keyword("not")
	}

	r.Keyword("of")
	r.Space()
	r.Symbol("(")

	for i, t := range io.Types {
		t.RenderTo(r)

		if i < len(io.Types)-1 {
			r.Symbol(",")
		}
	}

	r.Symbol(")")
}

type WithinGroupClause OrderClause

func (w WithinGroupClause) RenderTo(r Renderer) {
	r.Keyword("within group")
	r.Space()
	r.Symbol("(")
	OrderClause(w).RenderTo(r)
	r.Symbol(")")
}

type FilterClause struct {
	Expr
}

func (f FilterClause) RenderTo(r Renderer) {
	r.Keyword("filter")
	r.Space()
	r.Symbol("(")
	r.Keyword("where")
	f.Expr.RenderTo(r)
	r.Symbol(")")
}

type DefaultExpr bool

func (d DefaultExpr) RenderTo(r Renderer) {
	r.Keyword("default")
}

type Row struct {
	RowWord bool
	Exprs   []Expr
}

func (row Row) RenderTo(r Renderer) {
	if row.RowWord {
		r.Keyword("row")
	}

	r.Symbol("(")

	for i, e := range row.Exprs {
		e.RenderTo(r)
		if i < len(row.Exprs)-1 {
			r.Symbol(",")
		}
	}

	r.Symbol(")")
}

type ValuesRow []Expr

func (vr ValuesRow) RenderTo(r Renderer) {
	r.Symbol("(")

	for i, e := range vr {
		e.RenderTo(r)
		if i < len(vr)-1 {
			r.Symbol(",")
		}
	}

	r.Symbol(")")
}

type ValuesClause []ValuesRow

func (vc ValuesClause) RenderTo(r Renderer) {
	r.Keyword("values")
	r.NewLine()
	r.Indent()

	for i, row := range vc {
		row.RenderTo(r)
		if i < len(vc)-1 {
			r.Symbol(",")
		}
		r.NewLine()
	}

	r.Unindent()
}

type OverClause struct {
	Name          string
	Specification *WindowSpecification
}

func (oc *OverClause) RenderTo(r Renderer) {
	r.Keyword("over")
	r.Space()
	if oc.Name != "" {
		r.Identifier(oc.Name)
	} else {
		oc.Specification.RenderTo(r)
	}
}

type WindowClause []WindowDefinition

func (wc WindowClause) RenderTo(r Renderer) {
	r.Keyword("window")
	r.NewLine()
	r.Indent()

	for i, wd := range wc {
		wd.RenderTo(r)
		if i < len(wc)-1 {
			r.Symbol(",")
		}
		r.NewLine()
	}

	r.Unindent()
}

type WindowDefinition struct {
	Name          string
	Specification WindowSpecification
}

func (wd WindowDefinition) RenderTo(r Renderer) {
	r.Identifier(wd.Name)
	r.Keyword("as")
	r.Space()
	wd.Specification.RenderTo(r)
}

type WindowSpecification struct {
	ExistingName    string
	PartitionClause PartitionClause
	OrderClause     *OrderClause
	FrameClause     *FrameClause
}

func (ws WindowSpecification) RenderTo(r Renderer) {
	r.Symbol("(")

	if ws.ExistingName != "" {
		r.Identifier(ws.ExistingName)
	}

	if ws.PartitionClause != nil {
		ws.PartitionClause.RenderTo(r)
	}

	if ws.OrderClause != nil {
		ws.OrderClause.RenderTo(r)
	}

	if ws.FrameClause != nil {
		ws.FrameClause.RenderTo(r)
	}

	r.Symbol(")")
}

type PartitionClause []Expr

func (pc PartitionClause) RenderTo(r Renderer) {
	r.Keyword("partition by")

	for i, e := range pc {
		e.RenderTo(r)
		if i < len(pc)-1 {
			r.Symbol(",")
		}
	}
}

type FrameClause struct {
	Mode  string
	Start *FrameBound
	End   *FrameBound
}

func (fc *FrameClause) RenderTo(r Renderer) {
	r.Keyword(fc.Mode)

	if fc.End != nil {
		r.Keyword("between")
		fc.Start.RenderTo(r)
		r.Keyword("and")
		fc.End.RenderTo(r)
	} else {
		fc.Start.RenderTo(r)
	}
}

type FrameBound struct {
	CurrentRow bool

	BoundExpr Expr
	Direction string
}

func (fb FrameBound) RenderTo(r Renderer) {
	if fb.CurrentRow {
		r.Keyword("current row")
		return
	}

	if fb.BoundExpr != nil {
		fb.BoundExpr.RenderTo(r)
	} else {
		r.Keyword("unbounded")
	}

	r.Keyword(fb.Direction)
}

type RelationExpr struct {
	Name AnyName
	Star bool
	Only bool
}

func (re RelationExpr) RenderTo(r Renderer) {
	if re.Only {
		r.Keyword("only")
	}

	re.Name.RenderTo(r)

	if re.Star {
		r.Symbol("*")
	}

	r.NewLine()
}

type SimpleSelect struct {
	DistinctList  []Expr
	TargetList    []Expr
	IntoClause    *IntoClause
	FromClause    *FromClause
	WhereClause   *WhereClause
	GroupByClause *GroupByClause
	HavingClause  Expr
	WindowClause  WindowClause

	ValuesClause ValuesClause

	LeftSelect  *SelectStmt
	SetOp       string
	SetAll      bool
	RightSelect *SelectStmt

	Table *RelationExpr
}

func (s SimpleSelect) RenderTo(r Renderer) {
	if s.Table != nil {
		r.Keyword("table")
		s.Table.RenderTo(r)
		return
	}

	if s.ValuesClause != nil {
		s.ValuesClause.RenderTo(r)
		return
	}

	if s.LeftSelect != nil {
		s.LeftSelect.RenderTo(r)
		r.NewLine()
		r.Keyword(s.SetOp)

		if s.SetAll {
			r.Keyword("all")
		}

		r.NewLine()

		s.RightSelect.RenderTo(r)

		return
	}

	r.Keyword("select")

	if s.DistinctList != nil {
		r.Keyword("distinct")

		if len(s.DistinctList) > 0 {
			r.Keyword("on")
			r.Symbol("(")

			for i, f := range s.DistinctList {
				f.RenderTo(r)
				if i < len(s.DistinctList)-1 {
					r.Symbol(",")
				}
			}
			r.Symbol(")")
		}

	}

	r.NewLine()
	r.Indent()
	for i, f := range s.TargetList {
		f.RenderTo(r)
		if i < len(s.TargetList)-1 {
			r.Symbol(",")
		}
		r.NewLine()
	}
	r.Unindent()

	if s.IntoClause != nil {
		s.IntoClause.RenderTo(r)
	}

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
		r.Keyword("having")
		r.NewLine()
		r.Indent()
		s.HavingClause.RenderTo(r)
		r.NewLine()
	}

	if s.WindowClause != nil {
		s.WindowClause.RenderTo(r)
	}
}

type SelectStmt struct {
	SimpleSelect
	OrderClause   *OrderClause
	LimitClause   *LimitClause
	LockingClause *LockingClause

	ParenWrapped bool
	Semicolon    bool
}

func (s SelectStmt) RenderTo(r Renderer) {
	if s.ParenWrapped {
		r.Symbol("(")
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
		r.Symbol(")")
		r.NewLine()
	}

	if s.Semicolon {
		r.Symbol(";")
		r.NewLine()
	}
}

type ExistsExpr SelectStmt

func (e ExistsExpr) RenderTo(r Renderer) {
	r.Keyword("exists")

	SelectStmt(e).RenderTo(r)
}

type ArraySubselect SelectStmt

func (a ArraySubselect) RenderTo(r Renderer) {
	r.Keyword("array")

	SelectStmt(a).RenderTo(r)
}
