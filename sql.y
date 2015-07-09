%{

package main

%}

%union {
  sqlSelect *SelectStmt
  fields []Expr
  expr Expr
  src string
  keyword string
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

%token  <src> IDENT
%token  <src> STRING_LITERAL
%token  <src> NUMBER_LITERAL
%token  <src> OP


/* ordinary key words in alphabetical order */
%token <keyword> ABORT_P ABSOLUTE_P ACCESS ACTION ADD_P ADMIN AFTER
  AGGREGATE ALL ALSO ALTER ALWAYS ANALYSE ANALYZE AND ANY ARRAY AS ASC
  ASSERTION ASSIGNMENT ASYMMETRIC AT ATTRIBUTE AUTHORIZATION

  BACKWARD BEFORE BEGIN_P BETWEEN BIGINT BINARY BIT
  BOOLEAN_P BOTH BY

  CACHE CALLED CASCADE CASCADED CASE CAST CATALOG_P CHAIN CHAR_P
  CHARACTER CHARACTERISTICS CHECK CHECKPOINT CLASS CLOSE
  CLUSTER COALESCE COLLATE COLLATION COLUMN COMMENT COMMENTS COMMIT
  COMMITTED CONCURRENTLY CONFIGURATION CONFLICT CONNECTION CONSTRAINT
  CONSTRAINTS CONTENT_P CONTINUE_P CONVERSION_P COPY COST CREATE
  CROSS CSV CUBE CURRENT_P
  CURRENT_CATALOG CURRENT_DATE CURRENT_ROLE CURRENT_SCHEMA
  CURRENT_TIME CURRENT_TIMESTAMP CURRENT_USER CURSOR CYCLE

  DATA_P DATABASE DAY_P DEALLOCATE DEC DECIMAL_P DECLARE DEFAULT DEFAULTS
  DEFERRABLE DEFERRED DEFINER DELETE_P DELIMITER DELIMITERS DESC
  DICTIONARY DISABLE_P DISCARD DISTINCT DO DOCUMENT_P DOMAIN_P DOUBLE_P DROP

  EACH ELSE ENABLE_P ENCODING ENCRYPTED END_P ENUM_P ESCAPE EVENT EXCEPT
  EXCLUDE EXCLUDING EXCLUSIVE EXECUTE EXISTS EXPLAIN
  EXTENSION EXTERNAL EXTRACT

  FALSE_P FAMILY FETCH FILTER FIRST_P FLOAT_P FOLLOWING FOR
  FORCE FOREIGN FORWARD FREEZE FROM FULL FUNCTION FUNCTIONS

  GLOBAL GRANT GRANTED GREATEST GROUP_P GROUPING

  HANDLER HAVING HEADER_P HOLD HOUR_P

  IDENTITY_P IF_P ILIKE IMMEDIATE IMMUTABLE IMPLICIT_P IMPORT_P IN_P
  INCLUDING INCREMENT INDEX INDEXES INHERIT INHERITS INITIALLY INLINE_P
  INNER_P INOUT INPUT_P INSENSITIVE INSERT INSTEAD INT_P INTEGER
  INTERSECT INTERVAL INTO INVOKER IS ISNULL ISOLATION

  JOIN

  KEY

  LABEL LANGUAGE LARGE_P LAST_P LATERAL_P
  LEADING LEAKPROOF LEAST LEFT LEVEL LIKE LIMIT LISTEN LOAD LOCAL
  LOCALTIME LOCALTIMESTAMP LOCATION LOCK_P LOCKED LOGGED

  MAPPING MATCH MATERIALIZED MAXVALUE MINUTE_P MINVALUE MODE MONTH_P MOVE

  NAME_P NAMES NATIONAL NATURAL NCHAR NEXT NO NONE
  NOT NOTHING NOTIFY NOTNULL NOWAIT NULL_P NULLIF
  NULLS_P NUMERIC

  OBJECT_P OF OFF OFFSET OIDS ON ONLY OPERATOR OPTION OPTIONS OR
  ORDER ORDINALITY OUT_P OUTER_P OVER OVERLAPS OVERLAY OWNED OWNER

  PARSER PARTIAL PARTITION PASSING PASSWORD PLACING PLANS POLICY POSITION
  PRECEDING PRECISION PRESERVE PREPARE PREPARED PRIMARY
  PRIOR PRIVILEGES PROCEDURAL PROCEDURE PROGRAM

  QUOTE

  RANGE READ REAL REASSIGN RECHECK RECURSIVE REF REFERENCES REFRESH REINDEX
  RELATIVE_P RELEASE RENAME REPEATABLE REPLACE REPLICA
  RESET RESTART RESTRICT RETURNING RETURNS REVOKE RIGHT ROLE ROLLBACK ROLLUP
  ROW ROWS RULE

  SAVEPOINT SCHEMA SCROLL SEARCH SECOND_P SECURITY SELECT SEQUENCE SEQUENCES
  SERIALIZABLE SERVER SESSION SESSION_USER SET SETS SETOF SHARE SHOW
  SIMILAR SIMPLE SKIP SMALLINT SNAPSHOT SOME SQL_P STABLE STANDALONE_P START
  STATEMENT STATISTICS STDIN STDOUT STORAGE STRICT_P STRIP_P SUBSTRING
  SYMMETRIC SYSID SYSTEM_P

  TABLE TABLES TABLESAMPLE TABLESPACE TEMP TEMPLATE TEMPORARY TEXT_P THEN
  TIME TIMESTAMP TO TRAILING TRANSACTION TRANSFORM TREAT TRIGGER TRIM TRUE_P
  TRUNCATE TRUSTED TYPE_P TYPES_P

  UNBOUNDED UNCOMMITTED UNENCRYPTED UNION UNIQUE UNKNOWN UNLISTEN UNLOGGED
  UNTIL UPDATE USER USING

  VACUUM VALID VALIDATE VALIDATOR VALUE_P VALUES VARCHAR VARIADIC VARYING
  VERBOSE VERSION_P VIEW VIEWS VOLATILE

  WHEN WHERE WHITESPACE_P WINDOW WITH WITHIN WITHOUT WORK WRAPPER WRITE

  XML_P XMLATTRIBUTES XMLCONCAT XMLELEMENT XMLEXISTS XMLFOREST XMLPARSE
  XMLPI XMLROOT XMLSERIALIZE

  YEAR_P YES_P

  ZONE



%left OP
%left AND
%left OR
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
| expr AS IDENT
  {
    $$ = AliasedExpr{Expr: $1, Alias: $3}
  }
| expr IDENT
  {
    $$ = AliasedExpr{Expr: $1, Alias: $2}
  }

expr:
  IDENT
  {
    $$ = ColumnRef{Column: $1}
  }
| IDENT '.' IDENT
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
| TRUE_P /* temp hack while integrating PostgreSQL keywords */
  {
    $$ = ColumnRef{Column: "true"}
  }
| expr OP expr
  {
    $$ = BinaryExpr{Left: $1, Operator: $2, Right: $3}
  }
| expr AND expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "and", Right: $3}
  }
| expr OR expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "or", Right: $3}
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
  IDENT
  {
    $$ = []string{$1}
  }
| identifierSeq ',' IDENT
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
    $$ = OrderExpr{Expr: $1, Order: "asc"}
  }
| expr DESC
  {
    $$ = OrderExpr{Expr: $1, Order: "desc"}
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

