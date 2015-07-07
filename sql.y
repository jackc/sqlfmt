%{

package main

%}

%union {
  sqlSelect *SelectStmt
  fields []SelectExpr
  field SelectExpr
  expr Expr
  src string
}

%type <sqlSelect> top
%type <fields> selectClause
%type <fields> selectExprSeq
%type <field> selectExpr
%type <expr> expr
%type <src> fromClause
%type <src> tableExpr

%token  <src> COMMA
%token  <src> PERIOD
%token  <src> SELECT
%token  <src> AS
%token  <src> FROM
%token  <src> IDENTIFIER
%token  <src> STRING_LITERAL
%token  <src> NUMBER_LITERAL
%token  <src> OPERATOR

%left OPERATOR

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
  expr
  {
    $$ = SelectExpr{Expr: $1}
  }
| expr AS IDENTIFIER
  {
    $$ = SelectExpr{Expr: $1, Alias: $3}
  }
| expr IDENTIFIER
  {
    $$ = SelectExpr{Expr: $1, Alias: $2}
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

