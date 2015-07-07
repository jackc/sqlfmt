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
}

const COMMA = 57346
const PERIOD = 57347
const SELECT = 57348
const AS = 57349
const FROM = 57350
const CROSS = 57351
const JOIN = 57352
const USING = 57353
const ON = 57354
const IDENTIFIER = 57355
const STRING_LITERAL = 57356
const NUMBER_LITERAL = 57357
const OPERATOR = 57358
const LPAREN = 57359
const RPAREN = 57360

var sqlToknames = []string{
	"COMMA",
	"PERIOD",
	"SELECT",
	"AS",
	"FROM",
	"CROSS",
	"JOIN",
	"USING",
	"ON",
	"IDENTIFIER",
	"STRING_LITERAL",
	"NUMBER_LITERAL",
	"OPERATOR",
	"LPAREN",
	"RPAREN",
}
var sqlStatenames = []string{}

const sqlEofCode = 1
const sqlErrCode = 2
const sqlMaxDepth = 200

//line sql.y:164

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0

//line yacctab:1
var sqlExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const sqlNprod = 25
const sqlPrivate = 57344

var sqlTokenNames []string
var sqlStates []string

const sqlLast = 51

var sqlAct = []int{

	9, 8, 4, 10, 11, 12, 31, 13, 14, 10,
	11, 12, 42, 13, 21, 19, 19, 30, 26, 38,
	28, 6, 17, 2, 44, 32, 43, 34, 18, 41,
	29, 19, 36, 37, 23, 35, 27, 22, 39, 24,
	25, 33, 4, 20, 16, 15, 40, 5, 7, 3,
	1,
}
var sqlPact = []int{

	36, -1000, -1000, 13, -10, -1000, -10, 40, -1000, 15,
	38, -1000, -1000, -4, 30, -1000, -10, 23, -1000, -10,
	17, -1, -12, -10, 31, -10, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -10, 21, -1000, 2, -10, 16, 0,
	8, -1000, 11, -1000, -1000,
}
var sqlPgo = []int{

	0, 50, 23, 49, 48, 1, 0, 47, 46, 45,
}
var sqlR1 = []int{

	0, 1, 2, 2, 3, 4, 4, 5, 5, 5,
	6, 6, 6, 6, 6, 6, 6, 8, 8, 9,
	9, 9, 9, 7, 7,
}
var sqlR2 = []int{

	0, 1, 1, 2, 2, 1, 3, 1, 3, 2,
	1, 3, 1, 1, 3, 3, 3, 1, 3, 3,
	4, 7, 5, 2, 2,
}
var sqlChk = []int{

	-1000, -1, -2, -3, 6, -7, 8, -4, -5, -6,
	13, 14, 15, 17, -5, -9, 4, 7, 13, 16,
	5, -6, -2, 4, 9, 10, -5, 13, -6, 13,
	18, 18, -5, 10, -5, -5, 11, 12, 17, -6,
	-8, 13, 4, 18, 13,
}
var sqlDef = []int{

	0, -2, 1, 2, 0, 3, 0, 4, 5, 7,
	10, 12, 13, 0, 23, 24, 0, 0, 9, 0,
	0, 0, 0, 0, 0, 0, 6, 8, 14, 11,
	15, 16, 19, 0, 0, 20, 0, 0, 0, 22,
	0, 17, 0, 21, 18,
}
var sqlTok1 = []int{

	1,
}
var sqlTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18,
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
		//line sql.y:48
		{
			sqlVAL.sqlSelect = sqlS[sqlpt-0].sqlSelect
		}
	case 2:
		//line sql.y:54
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-0].fields
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 3:
		//line sql.y:60
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.Fields = sqlS[sqlpt-1].fields
			sqlVAL.sqlSelect.FromClause = sqlS[sqlpt-0].fromClause
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 4:
		//line sql.y:69
		{
			sqlVAL.fields = sqlS[sqlpt-0].fields
		}
	case 5:
		//line sql.y:75
		{
			sqlVAL.fields = []Expr{sqlS[sqlpt-0].expr}
		}
	case 6:
		//line sql.y:79
		{
			sqlVAL.fields = append(sqlS[sqlpt-2].fields, sqlS[sqlpt-0].expr)
		}
	case 7:
		//line sql.y:85
		{
			sqlVAL.expr = sqlS[sqlpt-0].expr
		}
	case 8:
		//line sql.y:89
		{
			sqlVAL.expr = AliasedExpr{Expr: sqlS[sqlpt-2].expr, Alias: sqlS[sqlpt-0].src}
		}
	case 9:
		//line sql.y:93
		{
			sqlVAL.expr = AliasedExpr{Expr: sqlS[sqlpt-1].expr, Alias: sqlS[sqlpt-0].src}
		}
	case 10:
		//line sql.y:99
		{
			sqlVAL.expr = ColumnRef{Column: sqlS[sqlpt-0].src}
		}
	case 11:
		//line sql.y:103
		{
			sqlVAL.expr = ColumnRef{Table: sqlS[sqlpt-2].src, Column: sqlS[sqlpt-0].src}
		}
	case 12:
		//line sql.y:107
		{
			sqlVAL.expr = StringLiteral(sqlS[sqlpt-0].src)
		}
	case 13:
		//line sql.y:111
		{
			sqlVAL.expr = IntegerLiteral(sqlS[sqlpt-0].src)
		}
	case 14:
		//line sql.y:115
		{
			sqlVAL.expr = BinaryExpr{Left: sqlS[sqlpt-2].expr, Operator: sqlS[sqlpt-1].src, Right: sqlS[sqlpt-0].expr}
		}
	case 15:
		//line sql.y:119
		{
			sqlVAL.expr = ParenExpr{Expr: sqlS[sqlpt-1].expr}
		}
	case 16:
		//line sql.y:123
		{
			sqlVAL.expr = ParenExpr{Expr: sqlS[sqlpt-1].sqlSelect}
		}
	case 17:
		//line sql.y:129
		{
			sqlVAL.identifiers = []string{sqlS[sqlpt-0].src}
		}
	case 18:
		//line sql.y:133
		{
			sqlVAL.identifiers = append(sqlS[sqlpt-2].identifiers, sqlS[sqlpt-0].src)
		}
	case 19:
		//line sql.y:139
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-2].expr, Join: sqlS[sqlpt-1].src, Right: sqlS[sqlpt-0].expr}
		}
	case 20:
		//line sql.y:143
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-3].expr, Join: "cross join", Right: sqlS[sqlpt-0].expr}
		}
	case 21:
		//line sql.y:147
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-6].expr, Join: "join", Right: sqlS[sqlpt-4].expr, Using: sqlS[sqlpt-1].identifiers}
		}
	case 22:
		//line sql.y:151
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-4].expr, Join: "join", Right: sqlS[sqlpt-2].expr, On: sqlS[sqlpt-0].expr}
		}
	case 23:
		//line sql.y:157
		{
			sqlVAL.fromClause = &FromClause{Expr: sqlS[sqlpt-0].expr}
		}
	case 24:
		//line sql.y:161
		{
			sqlVAL.fromClause = &FromClause{Expr: sqlS[sqlpt-0].expr}
		}
	}
	goto sqlstack /* stack new state and value */
}
