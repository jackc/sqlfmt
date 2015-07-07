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

%token  <src> COMMA
%token  <src> PERIOD
%token  <src> SELECT
%token  <src> AS
%token  <src> FROM
%token  <src> CROSS
%token  <src> JOIN
%token  <src> USING
%token  <src> IDENTIFIER
%token  <src> STRING_LITERAL
%token  <src> NUMBER_LITERAL
%token  <src> OPERATOR
%token  <src> LPAREN
%token  <src> RPAREN

%left OPERATOR

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
| selectClause fromClause
  {
    $$ = &SelectStmt{}
    $$.Fields = $1
    $$.FromClause = $2
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
| selectExprSeq COMMA aliasableExpr
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
| IDENTIFIER PERIOD IDENTIFIER
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
| LPAREN expr RPAREN
  {
    $$ = ParenExpr{Expr: $2}
  }
| LPAREN selectStatement RPAREN
  {
    $$ = ParenExpr{Expr: $2}
  }

identifierSeq:
  IDENTIFIER
  {
    $$ = []string{$1}
  }
| identifierSeq COMMA IDENTIFIER
  {
    $$ = append($1, $3)
  }

joinExpr:
  aliasableExpr COMMA aliasableExpr
  {
    $$ = JoinExpr{Left: $1, Join: $2, Right: $3}
  }
| aliasableExpr CROSS JOIN aliasableExpr
  {
    $$ = JoinExpr{Left: $1, Join: "cross join", Right: $4}
  }
| aliasableExpr JOIN aliasableExpr USING LPAREN identifierSeq RPAREN
  {
    $$ = JoinExpr{Left: $1, Join: "join", Right: $3, Using: $6}
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
%%

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0

