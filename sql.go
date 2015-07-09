//line sql.y:2
package main

import __yyfmt__ "fmt"

//line sql.y:3
//line sql.y:7
type sqlSymType struct {
	yys         int
	sqlSelect   *SelectStmt
	fields      []Expr
	expr        Expr
	src         string
	identifiers []string
	fromClause  *FromClause
	whereClause *WhereClause
	orderExpr   OrderExpr
	orderClause *OrderClause
}

const COMMA = 57346
const PERIOD = 57347
const SELECT = 57348
const AS = 57349
const FROM = 57350
const CROSS = 57351
const NATURAL = 57352
const JOIN = 57353
const USING = 57354
const ON = 57355
const IDENTIFIER = 57356
const STRING_LITERAL = 57357
const NUMBER_LITERAL = 57358
const OPERATOR = 57359
const NOT = 57360
const LPAREN = 57361
const RPAREN = 57362
const WHERE = 57363
const ORDER = 57364
const BY = 57365
const ASC = 57366
const DESC = 57367

var sqlToknames = []string{
	"COMMA",
	"PERIOD",
	"SELECT",
	"AS",
	"FROM",
	"CROSS",
	"NATURAL",
	"JOIN",
	"USING",
	"ON",
	"IDENTIFIER",
	"STRING_LITERAL",
	"NUMBER_LITERAL",
	"OPERATOR",
	"NOT",
	"LPAREN",
	"RPAREN",
	"WHERE",
	"ORDER",
	"BY",
	"ASC",
	"DESC",
}
var sqlStatenames = []string{}

const sqlEofCode = 1
const sqlErrCode = 2
const sqlMaxDepth = 200

//line sql.y:242

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0

//line yacctab:1
var sqlExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const sqlNprod = 37
const sqlPrivate = 57344

var sqlTokenNames []string
var sqlStates []string

const sqlLast = 76

var sqlAct = []int{

	11, 45, 10, 34, 27, 8, 20, 20, 62, 23,
	21, 52, 53, 4, 44, 58, 29, 30, 7, 27,
	17, 12, 13, 14, 63, 15, 16, 39, 41, 27,
	25, 8, 43, 2, 46, 46, 47, 26, 48, 32,
	27, 51, 12, 13, 14, 64, 15, 16, 61, 42,
	31, 40, 54, 55, 56, 57, 35, 50, 59, 49,
	4, 36, 37, 38, 6, 28, 33, 24, 19, 22,
	18, 60, 5, 9, 3, 1,
}
var sqlPact = []int{

	54, -1000, -1000, 10, 28, -16, -1000, 28, 28, 63,
	-1000, 23, 60, -1000, -1000, 28, 7, -1000, -15, 62,
	-20, 52, -1000, 2, 28, 37, -1000, 28, 35, -1000,
	12, -6, -1000, 28, 28, 28, 48, 46, 28, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -13, -1000, -1000, 28,
	28, 42, -1000, -1000, -1000, -1000, -4, 28, 34, 2,
	4, -1000, 31, -1000, -1000,
}
var sqlPgo = []int{

	0, 75, 33, 74, 73, 2, 0, 72, 71, 69,
	64, 1, 20, 68,
}
var sqlR1 = []int{

	0, 1, 2, 2, 2, 2, 3, 4, 4, 5,
	5, 5, 6, 6, 6, 6, 6, 6, 6, 6,
	8, 8, 9, 9, 9, 9, 9, 7, 7, 10,
	11, 11, 11, 12, 12, 13, 13,
}
var sqlR2 = []int{

	0, 1, 1, 3, 4, 2, 2, 1, 3, 1,
	3, 2, 1, 3, 1, 1, 3, 2, 3, 3,
	1, 3, 3, 4, 4, 7, 5, 2, 2, 2,
	1, 2, 2, 0, 1, 3, 3,
}
var sqlChk = []int{

	-1000, -1, -2, -3, 6, -7, -10, 8, 21, -4,
	-5, -6, 14, 15, 16, 18, 19, -12, -10, -13,
	22, -5, -9, -6, 4, 7, 14, 17, 5, -6,
	-6, -2, -12, 4, 23, 4, 9, 10, 11, -5,
	14, -6, 14, 20, 20, -11, -6, -11, -5, 11,
	11, -5, 24, 25, -5, -5, 12, 13, 19, -6,
	-8, 14, 4, 20, 14,
}
var sqlDef = []int{

	0, -2, 1, 2, 0, 33, 5, 0, 0, 6,
	7, 9, 12, 14, 15, 0, 0, 3, 33, 34,
	0, 27, 28, 29, 0, 0, 11, 0, 0, 17,
	0, 0, 4, 0, 0, 0, 0, 0, 0, 8,
	10, 16, 13, 18, 19, 36, 30, 35, 22, 0,
	0, 0, 31, 32, 23, 24, 0, 0, 0, 26,
	0, 20, 0, 25, 21,
}
var sqlTok1 = []int{

	1,
}
var sqlTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25,
}
var sqlTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var sqlDebug = 0

type sqlLexer interface {
	Lex(lval *sqlSymType) int
	Error(s string)
}

const sqlFlag = -1000

func sqlTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(sqlToknames) {
		if sqlToknames[c-4] != "" {
			return sqlToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func sqlStatname(s int) string {
	if s >= 0 && s < len(sqlStatenames) {
		if sqlStatenames[s] != "" {
			return sqlStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func sqllex1(lex sqlLexer, lval *sqlSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = sqlTok1[0]
		goto out
	}
	if char < len(sqlTok1) {
		c = sqlTok1[char]
		goto out
	}
	if char >= sqlPrivate {
		if char < sqlPrivate+len(sqlTok2) {
			c = sqlTok2[char-sqlPrivate]
			goto out
		}
	}
	for i := 0; i < len(sqlTok3); i += 2 {
		c = sqlTok3[i+0]
		if c == char {
			c = sqlTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = sqlTok2[1] /* unknown char */
	}
	if sqlDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", sqlTokname(c), uint(char))
	}
	return c
}

func sqlParse(sqllex sqlLexer) int {
	var sqln int
	var sqllval sqlSymType
	var sqlVAL sqlSymType
	sqlS := make([]sqlSymType, sqlMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	sqlstate := 0
	sqlchar := -1
	sqlp := -1
	goto sqlstack

ret0:
	return 0

ret1:
	return 1

sqlstack:
	/* put a state and value onto the stack */
	if sqlDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", sqlTokname(sqlchar), sqlStatname(sqlstate))
	}

	sqlp++
	if sqlp >= len(sqlS) {
		nyys := make([]sqlSymType, len(sqlS)*2)
		copy(nyys, sqlS)
		sqlS = nyys
	}
	sqlS[sqlp] = sqlVAL
	sqlS[sqlp].yys = sqlstate

sqlnewstate:
	sqln = sqlPact[sqlstate]
	if sqln <= sqlFlag {
		goto sqldefault /* simple state */
	}
	if sqlchar < 0 {
		sqlchar = sqllex1(sqllex, &sqllval)
	}
	sqln += sqlchar
	if sqln < 0 || sqln >= sqlLast {
		goto sqldefault
	}
	sqln = sqlAct[sqln]
	if sqlChk[sqln] == sqlchar { /* valid shift */
		sqlchar = -1
		sqlVAL = sqllval
		sqlstate = sqln
		if Errflag > 0 {
			Errflag--
		}
		goto sqlstack
	}

sqldefault:
	/* default state action */
	sqln = sqlDef[sqlstate]
	if sqln == -2 {
		if sqlchar < 0 {
			sqlchar = sqllex1(sqllex, &sqllval)
		}

		/* look through exception table */
		xi := 0
		for {
			if sqlExca[xi+0] == -1 && sqlExca[xi+1] == sqlstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			sqln = sqlExca[xi+0]
			if sqln < 0 || sqln == sqlchar {
				break
			}
		}
		sqln = sqlExca[xi+1]
		if sqln < 0 {
			goto ret0
		}
	}
	if sqln == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			sqllex.Error("syntax error")
			Nerrs++
			if sqlDebug >= 1 {
				__yyfmt__.Printf("%s", sqlStatname(sqlstate))
				__yyfmt__.Printf(" saw %s\n", sqlTokname(sqlchar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for sqlp >= 0 {
				sqln = sqlPact[sqlS[sqlp].yys] + sqlErrCode
				if sqln >= 0 && sqln < sqlLast {
					sqlstate = sqlAct[sqln] /* simulate a shift of "error" */
					if sqlChk[sqlstate] == sqlErrCode {
						goto sqlstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if sqlDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", sqlS[sqlp].yys)
				}
				sqlp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if sqlDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", sqlTokname(sqlchar))
			}
			if sqlchar == sqlEofCode {
				goto ret1
			}
			sqlchar = -1
			goto sqlnewstate /* try again in the same state */
		}
	}

	/* reduction by production sqln */
	if sqlDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", sqln, sqlStatname(sqlstate))
	}

	sqlnt := sqln
	sqlpt := sqlp
	_ = sqlpt // guard against "declared and not used"

	sqlp -= sqlR2[sqln]
	sqlVAL = sqlS[sqlp+1]

	/* consult goto table to find next state */
	sqln = sqlR1[sqln]
	sqlg := sqlPgo[sqln]
	sqlj := sqlg + sqlS[sqlp].yys + 1

	if sqlj >= sqlLast {
		sqlstate = sqlAct[sqlg]
	} else {
		sqlstate = sqlAct[sqlj]
		if sqlChk[sqlstate] != -sqln {
			sqlstate = sqlAct[sqlg]
		}
	}
	// dummy call; replaced with literal code
	switch sqlnt {

	case 1:
		//line sql.y:63
		{
			sqlVAL.sqlSelect = sqlS[sqlpt-0].sqlSelect
		}
	case 2:
		//line sql.y:69
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-0].fields
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 3:
		//line sql.y:75
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-2].fields
			sqlVAL.sqlSelect.FromClause = sqlS[sqlpt-1].fromClause
			sqlVAL.sqlSelect.OrderClause = sqlS[sqlpt-0].orderClause
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 4:
		//line sql.y:83
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-3].fields
			sqlVAL.sqlSelect.FromClause = sqlS[sqlpt-2].fromClause
			sqlVAL.sqlSelect.WhereClause = sqlS[sqlpt-1].whereClause
			sqlVAL.sqlSelect.OrderClause = sqlS[sqlpt-0].orderClause
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 5:
		//line sql.y:92
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-1].fields
			sqlVAL.sqlSelect.WhereClause = sqlS[sqlpt-0].whereClause
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 6:
		//line sql.y:101
		{
			sqlVAL.fields = sqlS[sqlpt-0].fields
		}
	case 7:
		//line sql.y:107
		{
			sqlVAL.fields = []Expr{sqlS[sqlpt-0].expr}
		}
	case 8:
		//line sql.y:111
		{
			sqlVAL.fields = append(sqlS[sqlpt-2].fields, sqlS[sqlpt-0].expr)
		}
	case 9:
		//line sql.y:117
		{
			sqlVAL.expr = sqlS[sqlpt-0].expr
		}
	case 10:
		//line sql.y:121
		{
			sqlVAL.expr = AliasedExpr{Expr: sqlS[sqlpt-2].expr, Alias: sqlS[sqlpt-0].src}
		}
	case 11:
		//line sql.y:125
		{
			sqlVAL.expr = AliasedExpr{Expr: sqlS[sqlpt-1].expr, Alias: sqlS[sqlpt-0].src}
		}
	case 12:
		//line sql.y:131
		{
			sqlVAL.expr = ColumnRef{Column: sqlS[sqlpt-0].src}
		}
	case 13:
		//line sql.y:135
		{
			sqlVAL.expr = ColumnRef{Table: sqlS[sqlpt-2].src, Column: sqlS[sqlpt-0].src}
		}
	case 14:
		//line sql.y:139
		{
			sqlVAL.expr = StringLiteral(sqlS[sqlpt-0].src)
		}
	case 15:
		//line sql.y:143
		{
			sqlVAL.expr = IntegerLiteral(sqlS[sqlpt-0].src)
		}
	case 16:
		//line sql.y:147
		{
			sqlVAL.expr = BinaryExpr{Left: sqlS[sqlpt-2].expr, Operator: sqlS[sqlpt-1].src, Right: sqlS[sqlpt-0].expr}
		}
	case 17:
		//line sql.y:151
		{
			sqlVAL.expr = NotExpr{Expr: sqlS[sqlpt-0].expr}
		}
	case 18:
		//line sql.y:155
		{
			sqlVAL.expr = ParenExpr{Expr: sqlS[sqlpt-1].expr}
		}
	case 19:
		//line sql.y:159
		{
			sqlVAL.expr = ParenExpr{Expr: sqlS[sqlpt-1].sqlSelect}
		}
	case 20:
		//line sql.y:165
		{
			sqlVAL.identifiers = []string{sqlS[sqlpt-0].src}
		}
	case 21:
		//line sql.y:169
		{
			sqlVAL.identifiers = append(sqlS[sqlpt-2].identifiers, sqlS[sqlpt-0].src)
		}
	case 22:
		//line sql.y:175
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-2].expr, Join: sqlS[sqlpt-1].src, Right: sqlS[sqlpt-0].expr}
		}
	case 23:
		//line sql.y:179
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-3].expr, Join: "cross join", Right: sqlS[sqlpt-0].expr}
		}
	case 24:
		//line sql.y:183
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-3].expr, Join: "natural join", Right: sqlS[sqlpt-0].expr}
		}
	case 25:
		//line sql.y:187
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-6].expr, Join: "join", Right: sqlS[sqlpt-4].expr, Using: sqlS[sqlpt-1].identifiers}
		}
	case 26:
		//line sql.y:191
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-4].expr, Join: "join", Right: sqlS[sqlpt-2].expr, On: sqlS[sqlpt-0].expr}
		}
	case 27:
		//line sql.y:197
		{
			sqlVAL.fromClause = &FromClause{Expr: sqlS[sqlpt-0].expr}
		}
	case 28:
		//line sql.y:201
		{
			sqlVAL.fromClause = &FromClause{Expr: sqlS[sqlpt-0].expr}
		}
	case 29:
		//line sql.y:207
		{
			sqlVAL.whereClause = &WhereClause{Expr: sqlS[sqlpt-0].expr}
		}
	case 30:
		//line sql.y:213
		{
			sqlVAL.orderExpr = OrderExpr{Expr: sqlS[sqlpt-0].expr}
		}
	case 31:
		//line sql.y:217
		{
			sqlVAL.orderExpr = OrderExpr{Expr: sqlS[sqlpt-1].expr, Order: sqlS[sqlpt-0].src}
		}
	case 32:
		//line sql.y:221
		{
			sqlVAL.orderExpr = OrderExpr{Expr: sqlS[sqlpt-1].expr, Order: sqlS[sqlpt-0].src}
		}
	case 33:
		//line sql.y:227
		{
			sqlVAL.orderClause = nil
		}
	case 34:
		sqlVAL.orderClause = sqlS[sqlpt-0].orderClause
	case 35:
		//line sql.y:234
		{
			sqlVAL.orderClause = &OrderClause{Exprs: []OrderExpr{sqlS[sqlpt-0].orderExpr}}
		}
	case 36:
		//line sql.y:238
		{
			sqlS[sqlpt-2].orderClause.Exprs = append(sqlS[sqlpt-2].orderClause.Exprs, sqlS[sqlpt-0].orderExpr)
			sqlVAL.orderClause = sqlS[sqlpt-2].orderClause
		}
	}
	goto sqlstack /* stack new state and value */
}
