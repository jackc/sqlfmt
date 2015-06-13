%{

package main

import (
  "fmt"
)

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
    fmt.Println("=", $$)
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


%%

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0

