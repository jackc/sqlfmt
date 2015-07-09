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

const SELECT = 57346
const AS = 57347
const FROM = 57348
const CROSS = 57349
const NATURAL = 57350
const JOIN = 57351
const USING = 57352
const ON = 57353
const IDENTIFIER = 57354
const STRING_LITERAL = 57355
const NUMBER_LITERAL = 57356
const OPERATOR = 57357
const NOT = 57358
const WHERE = 57359
const ORDER = 57360
const BY = 57361
const ASC = 57362
const DESC = 57363

var sqlToknames = []string{
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

//line sql.y:238

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

	11, 45, 10, 4, 62, 44, 28, 63, 58, 23,
	21, 12, 13, 14, 33, 15, 29, 30, 12, 13,
	14, 17, 15, 16, 24, 27, 34, 39, 41, 27,
	16, 36, 37, 38, 46, 46, 47, 27, 48, 43,
	32, 51, 52, 53, 8, 20, 35, 20, 7, 2,
	25, 64, 54, 55, 4, 61, 42, 26, 59, 8,
	27, 56, 57, 50, 40, 49, 31, 6, 19, 22,
	60, 5, 9, 18, 3, 1,
}
var sqlPact = []int{

	50, -1000, -1000, 42, 6, 27, -1000, 6, 6, 2,
	-1000, 45, -17, -1000, -1000, 6, -1, -1000, 29, -8,
	7, 24, -1000, 10, 6, 52, -1000, 6, 44, -1000,
	14, -20, -1000, 6, 6, 6, 56, 54, 6, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 22, -1000, -1000, 6,
	6, 51, -1000, -1000, -1000, -1000, -16, 6, 43, 10,
	-18, -1000, 39, -1000, -1000,
}
var sqlPgo = []int{

	0, 75, 49, 74, 72, 2, 0, 71, 70, 69,
	67, 1, 21, 68,
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

	-1000, -1, -2, -3, 4, -7, -10, 6, 17, -4,
	-5, -6, 12, 13, 14, 16, 24, -12, -10, -13,
	18, -5, -9, -6, 22, 5, 12, 15, 23, -6,
	-6, -2, -12, 22, 19, 22, 7, 8, 9, -5,
	12, -6, 12, 25, 25, -11, -6, -11, -5, 9,
	9, -5, 20, 21, -5, -5, 10, 11, 24, -6,
	-8, 12, 22, 25, 12,
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

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	24, 25, 3, 3, 22, 3, 23,
}
var sqlTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
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
		//line sql.y:59
		{
			sqlVAL.sqlSelect = sqlS[sqlpt-0].sqlSelect
		}
	case 2:
		//line sql.y:65
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-0].fields
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 3:
		//line sql.y:71
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-2].fields
			sqlVAL.sqlSelect.FromClause = sqlS[sqlpt-1].fromClause
			sqlVAL.sqlSelect.OrderClause = sqlS[sqlpt-0].orderClause
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 4:
		//line sql.y:79
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-3].fields
			sqlVAL.sqlSelect.FromClause = sqlS[sqlpt-2].fromClause
			sqlVAL.sqlSelect.WhereClause = sqlS[sqlpt-1].whereClause
			sqlVAL.sqlSelect.OrderClause = sqlS[sqlpt-0].orderClause
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 5:
		//line sql.y:88
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-1].fields
			sqlVAL.sqlSelect.WhereClause = sqlS[sqlpt-0].whereClause
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 6:
		//line sql.y:97
		{
			sqlVAL.fields = sqlS[sqlpt-0].fields
		}
	case 7:
		//line sql.y:103
		{
			sqlVAL.fields = []Expr{sqlS[sqlpt-0].expr}
		}
	case 8:
		//line sql.y:107
		{
			sqlVAL.fields = append(sqlS[sqlpt-2].fields, sqlS[sqlpt-0].expr)
		}
	case 9:
		//line sql.y:113
		{
			sqlVAL.expr = sqlS[sqlpt-0].expr
		}
	case 10:
		//line sql.y:117
		{
			sqlVAL.expr = AliasedExpr{Expr: sqlS[sqlpt-2].expr, Alias: sqlS[sqlpt-0].src}
		}
	case 11:
		//line sql.y:121
		{
			sqlVAL.expr = AliasedExpr{Expr: sqlS[sqlpt-1].expr, Alias: sqlS[sqlpt-0].src}
		}
	case 12:
		//line sql.y:127
		{
			sqlVAL.expr = ColumnRef{Column: sqlS[sqlpt-0].src}
		}
	case 13:
		//line sql.y:131
		{
			sqlVAL.expr = ColumnRef{Table: sqlS[sqlpt-2].src, Column: sqlS[sqlpt-0].src}
		}
	case 14:
		//line sql.y:135
		{
			sqlVAL.expr = StringLiteral(sqlS[sqlpt-0].src)
		}
	case 15:
		//line sql.y:139
		{
			sqlVAL.expr = IntegerLiteral(sqlS[sqlpt-0].src)
		}
	case 16:
		//line sql.y:143
		{
			sqlVAL.expr = BinaryExpr{Left: sqlS[sqlpt-2].expr, Operator: sqlS[sqlpt-1].src, Right: sqlS[sqlpt-0].expr}
		}
	case 17:
		//line sql.y:147
		{
			sqlVAL.expr = NotExpr{Expr: sqlS[sqlpt-0].expr}
		}
	case 18:
		//line sql.y:151
		{
			sqlVAL.expr = ParenExpr{Expr: sqlS[sqlpt-1].expr}
		}
	case 19:
		//line sql.y:155
		{
			sqlVAL.expr = ParenExpr{Expr: sqlS[sqlpt-1].sqlSelect}
		}
	case 20:
		//line sql.y:161
		{
			sqlVAL.identifiers = []string{sqlS[sqlpt-0].src}
		}
	case 21:
		//line sql.y:165
		{
			sqlVAL.identifiers = append(sqlS[sqlpt-2].identifiers, sqlS[sqlpt-0].src)
		}
	case 22:
		//line sql.y:171
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-2].expr, Join: ",", Right: sqlS[sqlpt-0].expr}
		}
	case 23:
		//line sql.y:175
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-3].expr, Join: "cross join", Right: sqlS[sqlpt-0].expr}
		}
	case 24:
		//line sql.y:179
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-3].expr, Join: "natural join", Right: sqlS[sqlpt-0].expr}
		}
	case 25:
		//line sql.y:183
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-6].expr, Join: "join", Right: sqlS[sqlpt-4].expr, Using: sqlS[sqlpt-1].identifiers}
		}
	case 26:
		//line sql.y:187
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-4].expr, Join: "join", Right: sqlS[sqlpt-2].expr, On: sqlS[sqlpt-0].expr}
		}
	case 27:
		//line sql.y:193
		{
			sqlVAL.fromClause = &FromClause{Expr: sqlS[sqlpt-0].expr}
		}
	case 28:
		//line sql.y:197
		{
			sqlVAL.fromClause = &FromClause{Expr: sqlS[sqlpt-0].expr}
		}
	case 29:
		//line sql.y:203
		{
			sqlVAL.whereClause = &WhereClause{Expr: sqlS[sqlpt-0].expr}
		}
	case 30:
		//line sql.y:209
		{
			sqlVAL.orderExpr = OrderExpr{Expr: sqlS[sqlpt-0].expr}
		}
	case 31:
		//line sql.y:213
		{
			sqlVAL.orderExpr = OrderExpr{Expr: sqlS[sqlpt-1].expr, Order: sqlS[sqlpt-0].src}
		}
	case 32:
		//line sql.y:217
		{
			sqlVAL.orderExpr = OrderExpr{Expr: sqlS[sqlpt-1].expr, Order: sqlS[sqlpt-0].src}
		}
	case 33:
		//line sql.y:223
		{
			sqlVAL.orderClause = nil
		}
	case 34:
		sqlVAL.orderClause = sqlS[sqlpt-0].orderClause
	case 35:
		//line sql.y:230
		{
			sqlVAL.orderClause = &OrderClause{Exprs: []OrderExpr{sqlS[sqlpt-0].orderExpr}}
		}
	case 36:
		//line sql.y:234
		{
			sqlS[sqlpt-2].orderClause.Exprs = append(sqlS[sqlpt-2].orderClause.Exprs, sqlS[sqlpt-0].orderExpr)
			sqlVAL.orderClause = sqlS[sqlpt-2].orderClause
		}
	}
	goto sqlstack /* stack new state and value */
}
