%{

package main

%}

%union {
  sqlSelect *SelectStmt
  fields []Expr
  expr Expr
  src string
  identifiers []string
  fromClause *FromClause
  whereClause *WhereClause
  orderExpr OrderExpr
  orderClause *OrderClause
}

%type <sqlSelect> top
%type <sqlSelect> selectStatement
%type <fields> selectClause
%type <fields> selectExprSeq
%type <expr> aliasableExpr
%type <expr> expr
%type <fromClause> fromClause
%type <identifiers> identifierSeq
%type <expr> joinExpr
%type <whereClause> whereClause
%type <orderExpr> orderExpr
%type <orderClause> optOrderClause
%type <orderClause> orderClause

%token  <src> SELECT
%token  <src> AS
%token  <src> FROM
%token  <src> CROSS
%token  <src> NATURAL
%token  <src> JOIN
%token  <src> USING
%token  <src> ON
%token  <src> IDENTIFIER
%token  <src> STRING_LITERAL
%token  <src> NUMBER_LITERAL
%token  <src> OPERATOR
%token  <src> NOT
%token  <src> WHERE
%token  <src> ORDER
%token  <src> BY
%token  <src> ASC
%token  <src> DESC

%left OPERATOR
%left NOT

%%

top:
  selectStatement
  {
    $$ = $1
  }

selectStatement:
  selectClause
  {
    $$ = &SelectStmt{}
    $$.Fields = $1
    sqllex.(*sqlLex).stmt = $$
  }
| selectClause fromClause optOrderClause
  {
    $$ = &SelectStmt{}
    $$.Fields = $1
    $$.FromClause = $2
    $$.OrderClause = $3
    sqllex.(*sqlLex).stmt = $$
  }
| selectClause fromClause whereClause optOrderClause
  {
    $$ = &SelectStmt{}
    $$.Fields = $1
    $$.FromClause = $2
    $$.WhereClause = $3
    $$.OrderClause = $4
    sqllex.(*sqlLex).stmt = $$
  }
| selectClause whereClause
  {
    $$ = &SelectStmt{}
    $$.Fields = $1
    $$.WhereClause = $2
    sqllex.(*sqlLex).stmt = $$
  }

selectClause:
  SELECT selectExprSeq
  {
    $$ = $2
  }

selectExprSeq:
  aliasableExpr
  {
    $$ = []Expr{$1}
  }
| selectExprSeq ',' aliasableExpr
  {
    $$ = append($1, $3)
  }

aliasableExpr:
  expr
  {
    $$ = $1
  }
| expr AS IDENTIFIER
  {
    $$ = AliasedExpr{Expr: $1, Alias: $3}
  }
| expr IDENTIFIER
  {
    $$ = AliasedExpr{Expr: $1, Alias: $2}
  }

expr:
  IDENTIFIER
  {
    $$ = ColumnRef{Column: $1}
  }
| IDENTIFIER '.' IDENTIFIER
  {
    $$ = ColumnRef{Table: $1, Column: $3}
  }
| STRING_LITERAL
  {
    $$ = StringLiteral($1)
  }
| NUMBER_LITERAL
  {
    $$ = IntegerLiteral($1)
  }
| expr OPERATOR expr
  {
    $$ = BinaryExpr{Left: $1, Operator: $2, Right: $3}
  }
| NOT expr
  {
    $$ = NotExpr{Expr: $2}
  }
| '(' expr ')'
  {
    $$ = ParenExpr{Expr: $2}
  }
| '(' selectStatement ')'
  {
    $$ = ParenExpr{Expr: $2}
  }

identifierSeq:
  IDENTIFIER
  {
    $$ = []string{$1}
  }
| identifierSeq ',' IDENTIFIER
  {
    $$ = append($1, $3)
  }

joinExpr:
  aliasableExpr ',' aliasableExpr
  {
    $$ = JoinExpr{Left: $1, Join: ",", Right: $3}
  }
| aliasableExpr CROSS JOIN aliasableExpr
  {
    $$ = JoinExpr{Left: $1, Join: "cross join", Right: $4}
  }
| aliasableExpr NATURAL JOIN aliasableExpr
  {
    $$ = JoinExpr{Left: $1, Join: "natural join", Right: $4}
  }
| aliasableExpr JOIN aliasableExpr USING '(' identifierSeq ')'
  {
    $$ = JoinExpr{Left: $1, Join: "join", Right: $3, Using: $6}
  }
| aliasableExpr JOIN aliasableExpr ON expr
  {
    $$ = JoinExpr{Left: $1, Join: "join", Right: $3, On: $5}
  }

fromClause:
  FROM aliasableExpr
  {
    $$ = &FromClause{Expr: $2}
  }
| FROM joinExpr
  {
    $$ = &FromClause{Expr: $2}
  }

whereClause:
  WHERE expr
  {
    $$ = &WhereClause{Expr: $2}
  }

orderExpr:
  expr
  {
    $$ = OrderExpr{Expr: $1}
  }
| expr ASC
  {
    $$ = OrderExpr{Expr: $1, Order: $2}
  }
| expr DESC
  {
    $$ = OrderExpr{Expr: $1, Order: $2}
  }

optOrderClause:
  /*EMPTY*/
  {
    $$ = nil
  }
| orderClause

orderClause:
  ORDER BY orderExpr
  {
    $$ = &OrderClause{Exprs: []OrderExpr{$3}}
  }
| orderClause ',' orderExpr
  {
    $1.Exprs = append($1.Exprs, $3)
    $$ = $1
  }
%%

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0

