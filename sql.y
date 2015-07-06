%{

package main

%}

%union {
  sqlSelect *SelectStmt
  fields []SelectExpr
  field SelectExpr
  src string
}

%type <sqlSelect> top
%type <fields> selectClause
%type <fields> selectExprSeq
%type <field> selectExpr
%type <src> fromClause
%type <src> tableExpr

%token  <src> COMMA
%token  <src> SELECT
%token  <src> AS
%token  <src> FROM
%token  <src> IDENTIFIER

%%

top:
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
    $$.FromTable = $2
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
    $$ = []SelectExpr{$1}
  }
| selectExprSeq COMMA selectExpr
  {
    $$ = append($1, $3)
  }

selectExpr:
  IDENTIFIER
  {
    $$ = SelectExpr{Expr: $1}
  }
| IDENTIFIER AS IDENTIFIER
  {
    $$ = SelectExpr{Expr: $1, Alias: $3}
  }
| IDENTIFIER IDENTIFIER
  {
    $$ = SelectExpr{Expr: $1, Alias: $2}
  }

fromClause:
  FROM tableExpr
  {
    $$ = $2
  }

tableExpr:
  IDENTIFIER
  {
    $$ = $1
  }

%%

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0

