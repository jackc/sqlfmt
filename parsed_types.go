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
		r.Text("setof", KeywordToken)
	}

	t.Name.RenderTo(r)

	if t.OptInterval != nil {
		t.OptInterval.RenderTo(r)
	}

	if t.ArrayWord {
		r.Text("array", KeywordToken)
	}

	for _, ab := range t.ArrayBounds {
		r.Text("[", SymbolToken)
		r.Text(string(ab), ConstantToken)
		r.Text("]", SymbolToken)
	}

	if len(t.TypeMods) > 0 {
		r.Text("(", SymbolToken)
		for i, e := range t.TypeMods {
			e.RenderTo(r)
			if i < len(t.TypeMods)-1 {
				r.Text(",", SymbolToken)
			}
		}
		r.Text(")", SymbolToken)
	}

	if t.WithTimeZone {
		r.Text("with time zone", KeywordToken)
	}

	if t.CharSet != "" {
		r.Text("character set", KeywordToken)
		r.Text(t.CharSet, IdentifierToken)
	}
}

type AnyName []string

func (an AnyName) RenderTo(r Renderer) {
	for i, n := range an {
		r.Text(n, IdentifierToken)
		if i < len(an)-1 {
			r.Text(".", SymbolToken)
		}
	}
}

type ColumnRef struct {
	Name        string
	Indirection Indirection
}

func (cr ColumnRef) RenderTo(r Renderer) {
	r.Text(cr.Name, IdentifierToken)
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
		r.Text("[", SymbolToken)
		ie.LowerSubscript.RenderTo(r)
		if ie.UpperSubscript != nil {
			r.Text(":", SymbolToken)
			ie.UpperSubscript.RenderTo(r)
		}
		r.Text("]", SymbolToken)
	} else {
		r.Text(".", SymbolToken)
		r.Text(ie.Name, IdentifierToken)
	}
}

type StringConst string

func (s StringConst) RenderTo(r Renderer) {
	r.Text(string(s), ConstantToken)
}

type IntegerConst string

func (s IntegerConst) RenderTo(r Renderer) {
	r.Text(string(s), ConstantToken)
}

type FloatConst string

func (s FloatConst) RenderTo(r Renderer) {
	r.Text(string(s), ConstantToken)
}

type BoolConst bool

func (b BoolConst) RenderTo(r Renderer) {
	if b {
		r.Text("true", KeywordToken)
	} else {
		r.Text("false", KeywordToken)
	}
}

type NullConst struct{}

func (n NullConst) RenderTo(r Renderer) {
	r.Text("null", KeywordToken)
}

type BitConst string

func (b BitConst) RenderTo(r Renderer) {
	r.Text(string(b), ConstantToken)
}

type BooleanExpr struct {
	Left     Expr
	Operator string
	Right    Expr
}

func (e BooleanExpr) RenderTo(r Renderer) {
	e.Left.RenderTo(r)
	r.Control(NewLineToken)
	r.Text(e.Operator, SymbolToken)
	e.Right.RenderTo(r)
}

type BinaryExpr struct {
	Left     Expr
	Operator AnyName
	Right    Expr
}

func (e BinaryExpr) RenderTo(r Renderer) {
	e.Left.RenderTo(r)
	r.Control(SpaceToken)
	e.Operator.RenderTo(r)
	r.Control(SpaceToken)
	e.Right.RenderTo(r)
}

type ArrayConstructorExpr ArrayExpr

func (ace ArrayConstructorExpr) RenderTo(r Renderer) {
	r.Text("array", KeywordToken)
	ArrayExpr(ace).RenderTo(r)
}

type ArrayExpr []Expr

func (a ArrayExpr) RenderTo(r Renderer) {
	r.Text("[", SymbolToken)
	for i, e := range a {
		e.RenderTo(r)
		if i < len(a)-1 {
			r.Text(",", SymbolToken)
		}
	}
	r.Text("]", SymbolToken)
}

type TextOpWithEscapeExpr struct {
	Left     Expr
	Operator string
	Right    Expr
	Escape   Expr
}

func (e TextOpWithEscapeExpr) RenderTo(r Renderer) {
	e.Left.RenderTo(r)
	r.Text(e.Operator, SymbolToken)
	e.Right.RenderTo(r)

	if e.Escape != nil {
		r.Text("escape", KeywordToken)
		e.Escape.RenderTo(r)
	}
}

type UnaryExpr struct {
	Operator AnyName
	Expr     Expr
}

func (e UnaryExpr) RenderTo(r Renderer) {
	e.Operator.RenderTo(r)
	r.Control(RefuseSpaceToken)
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
	r.Text(s.Type, KeywordToken)
	r.Control(SpaceToken)
	s.Query.RenderTo(r)
}

type SubqueryOp struct {
	Operator bool
	Name     AnyName
}

func (s SubqueryOp) RenderTo(r Renderer) {
	if s.Operator {
		r.Text("operator", KeywordToken)
		r.Text("(", SymbolToken)
	}
	s.Name.RenderTo(r)
	if s.Operator {
		r.Text(")", SymbolToken)
	}
}

type WhenClause struct {
	When Expr
	Then Expr
}

func (w WhenClause) RenderTo(r Renderer) {
	r.Text("when", KeywordToken)
	w.When.RenderTo(r)
	r.Text("then", KeywordToken)
	r.Control(NewLineToken)
	r.Control(IndentToken)
	w.Then.RenderTo(r)
	r.Control(NewLineToken)
	r.Control(UnindentToken)
}

type InExpr struct {
	Value Expr
	Not   bool
	In    Expr
}

func (i InExpr) RenderTo(r Renderer) {
	i.Value.RenderTo(r)

	if i.Not {
		r.Text("not", KeywordToken)
	}

	r.Text("in", KeywordToken)
	r.Control(SpaceToken)

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
		r.Text("not", KeywordToken)
	}

	r.Text("between", KeywordToken)

	if b.Symmetric {
		r.Text("symmetric", KeywordToken)
	}

	b.Left.RenderTo(r)
	r.Text("and", KeywordToken)
	b.Right.RenderTo(r)
}

type CaseExpr struct {
	CaseArg     Expr
	WhenClauses []WhenClause
	Default     Expr
}

func (c CaseExpr) RenderTo(r Renderer) {
	r.Text("case", KeywordToken)

	if c.CaseArg != nil {
		c.CaseArg.RenderTo(r)
	}

	r.Control(NewLineToken)

	for _, w := range c.WhenClauses {
		w.RenderTo(r)
	}

	if c.Default != nil {
		r.Text("else", KeywordToken)
		r.Control(NewLineToken)
		r.Control(IndentToken)
		c.Default.RenderTo(r)
		r.Control(NewLineToken)
		r.Control(UnindentToken)
	}

	r.Text("end", KeywordToken)
	r.Control(NewLineToken)
}

type ParenExpr struct {
	Expr        Expr
	Indirection Indirection
}

func (e ParenExpr) RenderTo(r Renderer) {
	r.Text("(", SymbolToken)
	e.Expr.RenderTo(r)
	r.Text(")", SymbolToken)
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
	r.Text("::", SymbolToken)
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
	r.Text("interval", KeywordToken)
	if i.Precision != "" {
		r.Text("(", SymbolToken)
		i.Precision.RenderTo(r)
		r.Text(")", SymbolToken)
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
		r.Text(oi.Left, KeywordToken)
	}

	if oi.Right != "" {
		r.Text("to", KeywordToken)
		r.Text(oi.Right, KeywordToken)
	}

	if oi.Second != nil {
		oi.Second.RenderTo(r)
	}
}

type IntervalSecond struct {
	Precision IntegerConst
}

func (is IntervalSecond) RenderTo(r Renderer) {
	r.Text("second", KeywordToken)
	if is.Precision != "" {
		r.Text("(", SymbolToken)
		is.Precision.RenderTo(r)
		r.Text(")", SymbolToken)
	}
}

type ExtractExpr ExtractList

func (ee ExtractExpr) RenderTo(r Renderer) {
	r.Text("extract", KeywordToken)
	r.Text("(", SymbolToken)
	ExtractList(ee).RenderTo(r)
	r.Text(")", SymbolToken)
}

type ExtractList struct {
	Extract Expr
	Time    Expr
}

func (el ExtractList) RenderTo(r Renderer) {
	el.Extract.RenderTo(r)
	r.Text("from", KeywordToken)
	el.Time.RenderTo(r)
}

type OverlayExpr OverlayList

func (oe OverlayExpr) RenderTo(r Renderer) {
	r.Text("overlay", KeywordToken)
	r.Text("(", SymbolToken)
	OverlayList(oe).RenderTo(r)
	r.Text(")", SymbolToken)
}

type OverlayList struct {
	Dest    Expr
	Placing Expr
	From    Expr
	For     Expr
}

func (ol OverlayList) RenderTo(r Renderer) {
	ol.Dest.RenderTo(r)
	r.Text("placing", KeywordToken)
	ol.Placing.RenderTo(r)
	r.Text("from", KeywordToken)
	ol.From.RenderTo(r)

	if ol.For != nil {
		r.Text("for", KeywordToken)
		ol.For.RenderTo(r)
	}
}

type PositionExpr PositionList

func (pe PositionExpr) RenderTo(r Renderer) {
	r.Text("position", KeywordToken)
	r.Text("(", SymbolToken)
	PositionList(pe).RenderTo(r)
	r.Text(")", SymbolToken)
}

type PositionList struct {
	Substring Expr
	String    Expr
}

func (pl PositionList) RenderTo(r Renderer) {
	pl.Substring.RenderTo(r)
	r.Text("in", KeywordToken)
	pl.String.RenderTo(r)
}

type SubstrExpr SubstrList

func (se SubstrExpr) RenderTo(r Renderer) {
	r.Text("substring", KeywordToken)
	r.Text("(", SymbolToken)
	SubstrList(se).RenderTo(r)
	r.Text(")", SymbolToken)
}

type SubstrList struct {
	Source Expr
	From   Expr
	For    Expr
}

func (sl SubstrList) RenderTo(r Renderer) {
	sl.Source.RenderTo(r)
	r.Text("from", KeywordToken)
	sl.From.RenderTo(r)

	if sl.For != nil {
		r.Text("for", KeywordToken)
		sl.For.RenderTo(r)
	}
}

type TrimExpr struct {
	Direction string
	TrimList
}

func (te TrimExpr) RenderTo(r Renderer) {
	r.Text("trim", KeywordToken)
	r.Text("(", SymbolToken)
	if te.Direction != "" {
		r.Text(te.Direction, KeywordToken)
	}
	te.TrimList.RenderTo(r)
	r.Text(")", SymbolToken)
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
		r.Text("from", KeywordToken)
	}

	for i, e := range tl.Right {
		e.RenderTo(r)
		if i+1 < len(tl.Right) {
			r.Text(",", SymbolToken)
		}
	}
}

type XmlElement struct {
	Name       string
	Attributes XmlAttributes
	Body       []Expr
}

func (el XmlElement) RenderTo(r Renderer) {
	r.Text("xmlelement", KeywordToken)
	r.Text("(", SymbolToken)
	r.Text("name", KeywordToken)
	r.Text(el.Name, IdentifierToken)

	if el.Attributes != nil {
		r.Text(",", SymbolToken)
		el.Attributes.RenderTo(r)
	}

	if el.Body != nil {
		for _, e := range el.Body {
			r.Text(",", SymbolToken)
			e.RenderTo(r)
		}
	}

	r.Text(")", SymbolToken)
}

type XmlAttributes []XmlAttributeEl

func (attrs XmlAttributes) RenderTo(r Renderer) {
	r.Text("xmlattributes", KeywordToken)
	r.Text("(", SymbolToken)
	xmlAttributes(attrs).RenderTo(r)
	r.Text(")", SymbolToken)
}

type XmlAttributeEl struct {
	Value Expr
	Name  string
}

func (el XmlAttributeEl) RenderTo(r Renderer) {
	el.Value.RenderTo(r)
	if el.Name != "" {
		r.Text("as", KeywordToken)
		r.Text(el.Name, IdentifierToken)
	}
}

type XmlExists struct {
	Path Expr
	Body XmlExistsArgument
}

func (e XmlExists) RenderTo(r Renderer) {
	r.Text("xmlexists", KeywordToken)
	r.Text("(", SymbolToken)
	e.Path.RenderTo(r)
	e.Body.RenderTo(r)
	r.Text(")", SymbolToken)
}

type XmlExistsArgument struct {
	LeftByRef  bool
	Arg        Expr
	RightByRef bool
}

func (a XmlExistsArgument) RenderTo(r Renderer) {
	r.Text("passing", KeywordToken)

	if a.LeftByRef {
		r.Text("by ref", KeywordToken)
	}

	a.Arg.RenderTo(r)

	if a.RightByRef {
		r.Text("by ref", KeywordToken)
	}
}

type XmlForest []XmlAttributeEl

func (f XmlForest) RenderTo(r Renderer) {
	r.Text("xmlforest", KeywordToken)
	r.Text("(", SymbolToken)
	xmlAttributes(f).RenderTo(r)
	r.Text(")", SymbolToken)
}

type xmlAttributes []XmlAttributeEl

func (attrs xmlAttributes) RenderTo(r Renderer) {
	for i, a := range attrs {
		a.RenderTo(r)
		if i+1 < len(attrs) {
			r.Text(",", SymbolToken)
		}
	}
}

type XmlParse struct {
	Type             string
	Content          Expr
	WhitespaceOption string
}

func (p XmlParse) RenderTo(r Renderer) {
	r.Text("xmlparse", KeywordToken)
	r.Text("(", SymbolToken)
	r.Text(p.Type, KeywordToken)
	p.Content.RenderTo(r)
	if p.WhitespaceOption != "" {
		r.Text(p.WhitespaceOption, KeywordToken)
	}

	r.Text(")", SymbolToken)
}

type XmlPi struct {
	Name    string
	Content Expr
}

func (p XmlPi) RenderTo(r Renderer) {
	r.Text("xmlpi", KeywordToken)
	r.Text("(", SymbolToken)
	r.Text("name", KeywordToken)
	r.Text(p.Name, IdentifierToken)

	if p.Content != nil {
		r.Text(",", SymbolToken)
		p.Content.RenderTo(r)
	}
	r.Text(")", SymbolToken)
}

type XmlRoot struct {
	Xml        Expr
	Version    XmlRootVersion
	Standalone string
}

func (x XmlRoot) RenderTo(r Renderer) {
	r.Text("xmlroot", KeywordToken)
	r.Text("(", SymbolToken)
	x.Xml.RenderTo(r)
	r.Text(",", SymbolToken)
	x.Version.RenderTo(r)
	if x.Standalone != "" {
		r.Text(",", SymbolToken)
		r.Text("standalone", KeywordToken)
		r.Text(x.Standalone, KeywordToken)
	}
	r.Text(")", SymbolToken)
}

type XmlRootVersion struct {
	Expr Expr
}

func (rv XmlRootVersion) RenderTo(r Renderer) {
	r.Text("version", KeywordToken)
	if rv.Expr != nil {
		rv.Expr.RenderTo(r)
	} else {
		r.Text("no value", KeywordToken)
	}
}

type XmlSerialize struct {
	XmlType string
	Content Expr
	Type    PgType
}

func (s XmlSerialize) RenderTo(r Renderer) {
	r.Text("xmlserialize", KeywordToken)
	r.Text("(", SymbolToken)
	r.Text(s.XmlType, KeywordToken)
	s.Content.RenderTo(r)
	r.Text("as", KeywordToken)
	s.Type.RenderTo(r)
	r.Text(")", SymbolToken)
}

type CollateExpr struct {
	Expr      Expr
	Collation AnyName
}

func (c CollateExpr) RenderTo(r Renderer) {
	c.Expr.RenderTo(r)
	r.Text("collate", KeywordToken)
	c.Collation.RenderTo(r)
}

type NotExpr struct {
	Expr Expr
}

func (e NotExpr) RenderTo(r Renderer) {
	r.Text("not", KeywordToken)
	e.Expr.RenderTo(r)
}

type IsExpr struct {
	Expr Expr
	Not  bool
	Op   string // null, document, true, false, etc.
}

func (e IsExpr) RenderTo(r Renderer) {
	e.Expr.RenderTo(r)
	r.Text("is", KeywordToken)
	if e.Not {
		r.Text("not", KeywordToken)
	}
	r.Text(e.Op, KeywordToken)
}

type AliasedExpr struct {
	Expr  Expr
	Alias string
}

func (e AliasedExpr) RenderTo(r Renderer) {
	e.Expr.RenderTo(r)
	r.Text("as", KeywordToken)
	r.Text(e.Alias, IdentifierToken)
}

type IntoClause struct {
	Options  string
	OptTable bool
	Target   AnyName
}

func (i IntoClause) RenderTo(r Renderer) {
	r.Text("into", KeywordToken)

	if i.Options != "" {
		r.Text(i.Options, KeywordToken)
	}

	if i.OptTable {
		r.Text("table", KeywordToken)
	}

	i.Target.RenderTo(r)
	r.Control(NewLineToken)
}

type FromClause struct {
	Expr Expr
}

func (e FromClause) RenderTo(r Renderer) {
	r.Text("from", KeywordToken)
	r.Control(NewLineToken)
	r.Control(IndentToken)
	e.Expr.RenderTo(r)
	r.Control(NewLineToken)
	r.Control(UnindentToken)
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
		r.Text(",", SymbolToken)
		r.Control(NewLineToken)
	} else {
		r.Control(NewLineToken)
		r.Text(s.Join, KeywordToken)
	}

	s.Right.RenderTo(r)

	if len(s.Using) > 0 {
		r.Text("using", KeywordToken)
		r.Text("(", SymbolToken)

		for i, u := range s.Using {
			r.Text(u, IdentifierToken)
			if i+1 < len(s.Using) {
				r.Text(",", SymbolToken)
			}
		}

		r.Text(")", SymbolToken)
	}

	if s.On != nil {
		r.Text("on", KeywordToken)
		s.On.RenderTo(r)
	}
}

type WhereClause struct {
	Expr Expr
}

func (e WhereClause) RenderTo(r Renderer) {
	r.Text("where", KeywordToken)
	r.Control(NewLineToken)
	r.Control(IndentToken)
	e.Expr.RenderTo(r)
	r.Control(NewLineToken)
	r.Control(UnindentToken)
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
		r.Text(e.Order, KeywordToken)
	}
	if len(e.Using) > 0 {
		r.Text("using", KeywordToken)
		e.Using.RenderTo(r)
	}
	if e.Nulls != "" {
		r.Text("nulls", KeywordToken)
		r.Text(e.Nulls, KeywordToken)
	}
}

type OrderClause struct {
	Exprs []OrderExpr
}

func (e OrderClause) RenderTo(r Renderer) {
	r.Text("order by", KeywordToken)
	r.Control(NewLineToken)
	r.Control(IndentToken)

	for i, f := range e.Exprs {
		f.RenderTo(r)
		if i < len(e.Exprs)-1 {
			r.Text(",", SymbolToken)
		}
		r.Control(NewLineToken)
	}
	r.Control(UnindentToken)
}

type GroupByClause struct {
	Exprs []Expr
}

func (e GroupByClause) RenderTo(r Renderer) {
	r.Text("group by", KeywordToken)
	r.Control(NewLineToken)
	r.Control(IndentToken)

	for i, f := range e.Exprs {
		f.RenderTo(r)
		if i < len(e.Exprs)-1 {
			r.Text(",", SymbolToken)
		}
		r.Control(NewLineToken)
	}
	r.Control(UnindentToken)
}

type LimitClause struct {
	Limit  Expr
	Offset Expr
}

func (e LimitClause) RenderTo(r Renderer) {
	if e.Limit != nil {
		r.Text("limit", KeywordToken)
		e.Limit.RenderTo(r)
		r.Control(NewLineToken)
	}
	if e.Offset != nil {
		r.Text("offset", KeywordToken)
		e.Offset.RenderTo(r)
		r.Control(NewLineToken)
	}
}

type AtTimeZoneExpr struct {
	Expr     Expr
	TimeZone Expr
}

func (e AtTimeZoneExpr) RenderTo(r Renderer) {
	e.Expr.RenderTo(r)
	r.Text("at time zone", KeywordToken)
	e.TimeZone.RenderTo(r)
}

type LockingItem struct {
	Strength   string
	LockedRels []AnyName
	WaitPolicy string
}

func (li LockingItem) RenderTo(r Renderer) {
	r.Text("for", KeywordToken)
	r.Text(li.Strength, KeywordToken)

	if li.LockedRels != nil {
		r.Text("of", KeywordToken)

		for i, lr := range li.LockedRels {
			lr.RenderTo(r)
			if i < len(li.LockedRels)-1 {
				r.Text(",", SymbolToken)
			}
		}
	}

	if li.WaitPolicy != "" {
		r.Text(li.WaitPolicy, KeywordToken)
	}

	r.Control(NewLineToken)
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
	r.Text(string(fe), KeywordToken)
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
	r.Text("(", SymbolToken)

	if fa.Distinct {
		r.Text("distinct", KeywordToken)
	}

	if fa.Star {
		r.Text("*", SymbolToken)
	} else if len(fa.Args) > 0 {
		for i, a := range fa.Args {
			a.RenderTo(r)
			if i < len(fa.Args)-1 {
				r.Text(",", SymbolToken)
			}
		}
	}

	if fa.VariadicArg != nil {
		if len(fa.Args) > 0 {
			r.Text(",", SymbolToken)
		}

		r.Text("variadic", KeywordToken)
		fa.VariadicArg.RenderTo(r)
	}

	if fa.OrderClause != nil {
		fa.OrderClause.RenderTo(r)
	}

	r.Text(")", SymbolToken)
}

type FuncArg struct {
	Name   string
	NameOp string
	Expr   Expr
}

func (fa FuncArg) RenderTo(r Renderer) {
	if fa.Name != "" {
		r.Text(fa.Name, IdentifierToken)
		r.Text(fa.NameOp, SymbolToken)
	}
	fa.Expr.RenderTo(r)
}

type CastFunc struct {
	Name string
	Expr Expr
	Type PgType
}

func (cf CastFunc) RenderTo(r Renderer) {
	r.Text(cf.Name, KeywordToken)
	r.Text("(", SymbolToken)
	cf.Expr.RenderTo(r)
	r.Text("as", KeywordToken)
	cf.Type.RenderTo(r)
	r.Text(")", SymbolToken)
}

type IsOfExpr struct {
	Expr  Expr
	Not   bool
	Types []PgType
}

func (io IsOfExpr) RenderTo(r Renderer) {
	io.Expr.RenderTo(r)
	r.Text("is", KeywordToken)

	if io.Not {
		r.Text("not", KeywordToken)
	}

	r.Text("of", KeywordToken)
	r.Control(SpaceToken)
	r.Text("(", SymbolToken)

	for i, t := range io.Types {
		t.RenderTo(r)

		if i < len(io.Types)-1 {
			r.Text(",", SymbolToken)
		}
	}

	r.Text(")", SymbolToken)
}

type WithinGroupClause OrderClause

func (w WithinGroupClause) RenderTo(r Renderer) {
	r.Text("within group", KeywordToken)
	r.Control(SpaceToken)
	r.Text("(", SymbolToken)
	OrderClause(w).RenderTo(r)
	r.Text(")", SymbolToken)
}

type FilterClause struct {
	Expr
}

func (f FilterClause) RenderTo(r Renderer) {
	r.Text("filter", KeywordToken)
	r.Control(SpaceToken)
	r.Text("(", SymbolToken)
	r.Text("where", KeywordToken)
	f.Expr.RenderTo(r)
	r.Text(")", SymbolToken)
}

type DefaultExpr bool

func (d DefaultExpr) RenderTo(r Renderer) {
	r.Text("default", KeywordToken)
}

type Row struct {
	RowWord bool
	Exprs   []Expr
}

func (row Row) RenderTo(r Renderer) {
	if row.RowWord {
		r.Text("row", KeywordToken)
	}

	r.Text("(", SymbolToken)

	for i, e := range row.Exprs {
		e.RenderTo(r)
		if i < len(row.Exprs)-1 {
			r.Text(",", SymbolToken)
		}
	}

	r.Text(")", SymbolToken)
}

type ValuesRow []Expr

func (vr ValuesRow) RenderTo(r Renderer) {
	r.Text("(", SymbolToken)

	for i, e := range vr {
		e.RenderTo(r)
		if i < len(vr)-1 {
			r.Text(",", SymbolToken)
		}
	}

	r.Text(")", SymbolToken)
}

type ValuesClause []ValuesRow

func (vc ValuesClause) RenderTo(r Renderer) {
	r.Text("values", KeywordToken)
	r.Control(NewLineToken)
	r.Control(IndentToken)

	for i, row := range vc {
		row.RenderTo(r)
		if i < len(vc)-1 {
			r.Text(",", SymbolToken)
		}
		r.Control(NewLineToken)
	}

	r.Control(UnindentToken)
}

type OverClause struct {
	Name          string
	Specification *WindowSpecification
}

func (oc *OverClause) RenderTo(r Renderer) {
	r.Text("over", KeywordToken)
	r.Control(SpaceToken)
	if oc.Name != "" {
		r.Text(oc.Name, IdentifierToken)
	} else {
		oc.Specification.RenderTo(r)
	}
}

type WindowClause []WindowDefinition

func (wc WindowClause) RenderTo(r Renderer) {
	r.Text("window", KeywordToken)
	r.Control(NewLineToken)
	r.Control(IndentToken)

	for i, wd := range wc {
		wd.RenderTo(r)
		if i < len(wc)-1 {
			r.Text(",", SymbolToken)
		}
		r.Control(NewLineToken)
	}

	r.Control(UnindentToken)
}

type WindowDefinition struct {
	Name          string
	Specification WindowSpecification
}

func (wd WindowDefinition) RenderTo(r Renderer) {
	r.Text(wd.Name, IdentifierToken)
	r.Text("as", KeywordToken)
	r.Control(SpaceToken)
	wd.Specification.RenderTo(r)
}

type WindowSpecification struct {
	ExistingName    string
	PartitionClause PartitionClause
	OrderClause     *OrderClause
	FrameClause     *FrameClause
}

func (ws WindowSpecification) RenderTo(r Renderer) {
	r.Text("(", SymbolToken)

	if ws.ExistingName != "" {
		r.Text(ws.ExistingName, IdentifierToken)
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

	r.Text(")", SymbolToken)
}

type PartitionClause []Expr

func (pc PartitionClause) RenderTo(r Renderer) {
	r.Text("partition by", KeywordToken)

	for i, e := range pc {
		e.RenderTo(r)
		if i < len(pc)-1 {
			r.Text(",", SymbolToken)
		}
	}
}

type FrameClause struct {
	Mode  string
	Start *FrameBound
	End   *FrameBound
}

func (fc *FrameClause) RenderTo(r Renderer) {
	r.Text(fc.Mode, KeywordToken)

	if fc.End != nil {
		r.Text("between", KeywordToken)
		fc.Start.RenderTo(r)
		r.Text("and", KeywordToken)
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
		r.Text("current row", KeywordToken)
		return
	}

	if fb.BoundExpr != nil {
		fb.BoundExpr.RenderTo(r)
	} else {
		r.Text("unbounded", KeywordToken)
	}

	r.Text(fb.Direction, KeywordToken)
}

type RelationExpr struct {
	Name AnyName
	Star bool
	Only bool
}

func (re RelationExpr) RenderTo(r Renderer) {
	if re.Only {
		r.Text("only", KeywordToken)
	}

	re.Name.RenderTo(r)

	if re.Star {
		r.Text("*", SymbolToken)
	}

	r.Control(NewLineToken)
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
		r.Text("table", KeywordToken)
		s.Table.RenderTo(r)
		return
	}

	if s.ValuesClause != nil {
		s.ValuesClause.RenderTo(r)
		return
	}

	if s.LeftSelect != nil {
		s.LeftSelect.RenderTo(r)
		r.Control(NewLineToken)
		r.Text(s.SetOp, KeywordToken)

		if s.SetAll {
			r.Text("all", KeywordToken)
		}

		r.Control(NewLineToken)

		s.RightSelect.RenderTo(r)

		return
	}

	r.Text("select", KeywordToken)

	if s.DistinctList != nil {
		r.Text("distinct", KeywordToken)

		if len(s.DistinctList) > 0 {
			r.Text("on", KeywordToken)
			r.Text("(", SymbolToken)

			for i, f := range s.DistinctList {
				f.RenderTo(r)
				if i < len(s.DistinctList)-1 {
					r.Text(",", SymbolToken)
				}
			}
			r.Text(")", SymbolToken)
		}

	}

	r.Control(NewLineToken)
	r.Control(IndentToken)
	for i, f := range s.TargetList {
		f.RenderTo(r)
		if i < len(s.TargetList)-1 {
			r.Text(",", SymbolToken)
		}
		r.Control(NewLineToken)
	}
	r.Control(UnindentToken)

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
		r.Text("having", KeywordToken)
		r.Control(NewLineToken)
		r.Control(IndentToken)
		s.HavingClause.RenderTo(r)
		r.Control(NewLineToken)
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
		r.Text("(", SymbolToken)
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
		r.Text(")", SymbolToken)
		r.Control(NewLineToken)
	}

	if s.Semicolon {
		r.Text(";", SymbolToken)
		r.Control(NewLineToken)
	}
}

type ExistsExpr SelectStmt

func (e ExistsExpr) RenderTo(r Renderer) {
	r.Text("exists", KeywordToken)

	SelectStmt(e).RenderTo(r)
}

type ArraySubselect SelectStmt

func (a ArraySubselect) RenderTo(r Renderer) {
	r.Text("array", KeywordToken)

	SelectStmt(a).RenderTo(r)
}
