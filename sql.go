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
}
var sqlStatenames = []string{}

const sqlEofCode = 1
const sqlErrCode = 2
const sqlMaxDepth = 200

//line sql.y:199

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0

//line yacctab:1
var sqlExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const sqlNprod = 30
const sqlPrivate = 57344

var sqlTokenNames []string
var sqlStates []string

const sqlLast = 62

var sqlAct = []int{

	11, 7, 10, 12, 13, 14, 51, 15, 16, 20,
	18, 8, 24, 47, 8, 37, 26, 27, 4, 38,
	24, 2, 52, 53, 33, 35, 12, 13, 14, 22,
	15, 16, 39, 50, 36, 42, 23, 34, 28, 24,
	45, 46, 41, 43, 44, 29, 40, 48, 4, 6,
	30, 31, 32, 25, 21, 17, 19, 49, 5, 9,
	3, 1,
}
var sqlPact = []int{

	42, -1000, -1000, -7, -11, -10, -1000, -11, -11, 50,
	-1000, 22, 48, -1000, -1000, -11, 12, -1000, 41, -1000,
	3, -11, 23, -1000, -11, 20, -1000, -5, -1, -11,
	35, 31, -11, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-11, -11, 28, -1000, -1000, -6, -11, 19, 3, 2,
	-1000, 9, -1000, -1000,
}
var sqlPgo = []int{

	0, 61, 21, 60, 59, 2, 0, 58, 57, 56,
	49,
}
var sqlR1 = []int{

	0, 1, 2, 2, 2, 2, 3, 4, 4, 5,
	5, 5, 6, 6, 6, 6, 6, 6, 6, 6,
	8, 8, 9, 9, 9, 9, 9, 7, 7, 10,
}
var sqlR2 = []int{

	0, 1, 1, 2, 3, 2, 2, 1, 3, 1,
	3, 2, 1, 3, 1, 1, 3, 2, 3, 3,
	1, 3, 3, 4, 4, 7, 5, 2, 2, 2,
}
var sqlChk = []int{

	-1000, -1, -2, -3, 6, -7, -10, 8, 21, -4,
	-5, -6, 14, 15, 16, 18, 19, -10, -5, -9,
	-6, 4, 7, 14, 17, 5, -6, -6, -2, 4,
	9, 10, 11, -5, 14, -6, 14, 20, 20, -5,
	11, 11, -5, -5, -5, 12, 13, 19, -6, -8,
	14, 4, 20, 14,
}
var sqlDef = []int{

	0, -2, 1, 2, 0, 3, 5, 0, 0, 6,
	7, 9, 12, 14, 15, 0, 0, 4, 27, 28,
	29, 0, 0, 11, 0, 0, 17, 0, 0, 0,
	0, 0, 0, 8, 10, 16, 13, 18, 19, 22,
	0, 0, 0, 23, 24, 0, 0, 0, 26, 0,
	20, 0, 25, 21,
}
var sqlTok1 = []int{

	1,
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
		//line sql.y:54
		{
			sqlVAL.sqlSelect = sqlS[sqlpt-0].sqlSelect
		}
	case 2:
		//line sql.y:60
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-0].fields
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 3:
		//line sql.y:66
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-1].fields
			sqlVAL.sqlSelect.FromClause = sqlS[sqlpt-0].fromClause
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 4:
		//line sql.y:73
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-2].fields
			sqlVAL.sqlSelect.FromClause = sqlS[sqlpt-1].fromClause
			sqlVAL.sqlSelect.WhereClause = sqlS[sqlpt-0].whereClause
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 5:
		//line sql.y:81
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-1].fields
			sqlVAL.sqlSelect.WhereClause = sqlS[sqlpt-0].whereClause
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 6:
		//line sql.y:90
		{
			sqlVAL.fields = sqlS[sqlpt-0].fields
		}
	case 7:
		//line sql.y:96
		{
			sqlVAL.fields = []Expr{sqlS[sqlpt-0].expr}
		}
	case 8:
		//line sql.y:100
		{
			sqlVAL.fields = append(sqlS[sqlpt-2].fields, sqlS[sqlpt-0].expr)
		}
	case 9:
		//line sql.y:106
		{
			sqlVAL.expr = sqlS[sqlpt-0].expr
		}
	case 10:
		//line sql.y:110
		{
			sqlVAL.expr = AliasedExpr{Expr: sqlS[sqlpt-2].expr, Alias: sqlS[sqlpt-0].src}
		}
	case 11:
		//line sql.y:114
		{
			sqlVAL.expr = AliasedExpr{Expr: sqlS[sqlpt-1].expr, Alias: sqlS[sqlpt-0].src}
		}
	case 12:
		//line sql.y:120
		{
			sqlVAL.expr = ColumnRef{Column: sqlS[sqlpt-0].src}
		}
	case 13:
		//line sql.y:124
		{
			sqlVAL.expr = ColumnRef{Table: sqlS[sqlpt-2].src, Column: sqlS[sqlpt-0].src}
		}
	case 14:
		//line sql.y:128
		{
			sqlVAL.expr = StringLiteral(sqlS[sqlpt-0].src)
		}
	case 15:
		//line sql.y:132
		{
			sqlVAL.expr = IntegerLiteral(sqlS[sqlpt-0].src)
		}
	case 16:
		//line sql.y:136
		{
			sqlVAL.expr = BinaryExpr{Left: sqlS[sqlpt-2].expr, Operator: sqlS[sqlpt-1].src, Right: sqlS[sqlpt-0].expr}
		}
	case 17:
		//line sql.y:140
		{
			sqlVAL.expr = NotExpr{Expr: sqlS[sqlpt-0].expr}
		}
	case 18:
		//line sql.y:144
		{
			sqlVAL.expr = ParenExpr{Expr: sqlS[sqlpt-1].expr}
		}
	case 19:
		//line sql.y:148
		{
			sqlVAL.expr = ParenExpr{Expr: sqlS[sqlpt-1].sqlSelect}
		}
	case 20:
		//line sql.y:154
		{
			sqlVAL.identifiers = []string{sqlS[sqlpt-0].src}
		}
	case 21:
		//line sql.y:158
		{
			sqlVAL.identifiers = append(sqlS[sqlpt-2].identifiers, sqlS[sqlpt-0].src)
		}
	case 22:
		//line sql.y:164
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-2].expr, Join: sqlS[sqlpt-1].src, Right: sqlS[sqlpt-0].expr}
		}
	case 23:
		//line sql.y:168
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-3].expr, Join: "cross join", Right: sqlS[sqlpt-0].expr}
		}
	case 24:
		//line sql.y:172
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-3].expr, Join: "natural join", Right: sqlS[sqlpt-0].expr}
		}
	case 25:
		//line sql.y:176
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-6].expr, Join: "join", Right: sqlS[sqlpt-4].expr, Using: sqlS[sqlpt-1].identifiers}
		}
	case 26:
		//line sql.y:180
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-4].expr, Join: "join", Right: sqlS[sqlpt-2].expr, On: sqlS[sqlpt-0].expr}
		}
	case 27:
		//line sql.y:186
		{
			sqlVAL.fromClause = &FromClause{Expr: sqlS[sqlpt-0].expr}
		}
	case 28:
		//line sql.y:190
		{
			sqlVAL.fromClause = &FromClause{Expr: sqlS[sqlpt-0].expr}
		}
	case 29:
		//line sql.y:196
		{
			sqlVAL.whereClause = &WhereClause{Expr: sqlS[sqlpt-0].expr}
		}
	}
	goto sqlstack /* stack new state and value */
}
