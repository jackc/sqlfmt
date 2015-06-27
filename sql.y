%{

package main

%}

%union {
  sqlSelect *SelectStmt
  fields []string
  src string
}

%type <sqlSelect> top
%type <fields> selectClause
%type <fields> selectExprSeq
%type <src> selectExpr
%type <src> fromClause
%type <src> tableExpr

%token  <src> COMMA
%token  <src> SELECT
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
    $$ = []string{$1}
  }
| selectExprSeq COMMA selectExpr
  {
    $$ = append($1, $3)
  }

selectExpr:
  IDENTIFIER
  {
    $$ = $1
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

