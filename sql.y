%{

package main

%}

%union {
  sqlSelect *SelectStmt
  fields []Expr
  expr Expr
  src string
  fromClause *FromClause
}

%type <sqlSelect> top
%type <sqlSelect> selectStatement
%type <fields> selectClause
%type <fields> selectExprSeq
%type <expr> selectExpr
%type <expr> expr
%type <fromClause> fromClause
%type <expr> joinExpr

%token  <src> COMMA
%token  <src> PERIOD
%token  <src> SELECT
%token  <src> AS
%token  <src> FROM
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
  selectExpr
  {
    $$ = []Expr{$1}
  }
| selectExprSeq COMMA selectExpr
  {
    $$ = append($1, $3)
  }

selectExpr:
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

joinExpr:
  expr COMMA expr
  {
    $$ = JoinExpr{Left: $1, Join: $2, Right: $3}
  }

fromClause:
  FROM expr
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

