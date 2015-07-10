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
%type <fields> opt_target_list target_list
%type <expr> aliasableExpr
%type <expr> expr target_el a_expr
%type <fromClause> fromClause
%type <identifiers> identifierSeq
%type <expr> joinExpr
%type <whereClause> where_clause
%type <orderExpr> orderExpr
%type <orderClause> optOrderClause
%type <orderClause> orderClause

%type <src>  Iconst SignedIconst Sconst

%type <src>
  ColLabel
  unreserved_keyword
  col_name_keyword
  type_func_name_keyword
  reserved_keyword



%token  <src> IDENT ICONST SCONST
%token  <src> OP


/* ordinary key words in alphabetical order */
%token <src> ABORT_P ABSOLUTE_P ACCESS ACTION ADD_P ADMIN AFTER
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
  selectClause fromClause where_clause optOrderClause
  {
    $$ = &SelectStmt{}
    $$.TargetList = $1
    $$.FromClause = $2
    $$.WhereClause = $3
    $$.OrderClause = $4
    sqllex.(*sqlLex).stmt = $$
  }
| selectClause where_clause
  {
    $$ = &SelectStmt{}
    $$.TargetList = $1
    $$.WhereClause = $2
    sqllex.(*sqlLex).stmt = $$
  }

selectClause:
  SELECT opt_target_list
  {
    $$ = $2
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
| Sconst
  {
    $$ = StringLiteral($1)
  }
| Iconst
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

a_expr: expr

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



where_clause:
  WHERE a_expr   { $$ = &WhereClause{Expr: $2} }
| /*EMPTY*/      { $$ = nil }


/*****************************************************************************
 *
 *  target list for SELECT
 *
 *****************************************************************************/

opt_target_list:
  target_list  { $$ = $1 }
| /* EMPTY */  { $$ = nil }

target_list:
  target_el    { $$ = []Expr{$1} }
| target_list ',' target_el
  {
    $$ = append($1, $3)
  }

target_el:
  a_expr AS ColLabel
  {
    $$ = AliasedExpr{Expr: $1, Alias: $3}
  }
| a_expr IDENT
  {
    $$ = AliasedExpr{Expr: $1, Alias: $2}
  }
| a_expr
/* TODO
      | '*'
*/

Iconst:   ICONST
Sconst:   SCONST

SignedIconst:
  Iconst      { $$ = $1 }
/* TODO - determine what to do with numbers
| '+' Iconst  { $$ = + $2 }
| '-' Iconst  { $$ = - $2 }
*/

/*
 * Name classification hierarchy.
 *
 * IDENT is the lexeme returned by the lexer for identifiers that match
 * no known keyword.  In most cases, we can accept certain keywords as
 * names, not only IDENTs.  We prefer to accept as many such keywords
 * as possible to minimize the impact of "reserved words" on programmers.
 * So, we divide names into several possible classes.  The classification
 * is chosen in part to make keywords acceptable as names wherever possible.
 */

/* Column identifier --- names that can be column, table, etc names.
 */
ColId:
  IDENT
| unreserved_keyword
| col_name_keyword

/* Type/function identifier --- names that can be type or function names.
 */
type_function_name:
  IDENT
| unreserved_keyword
| type_func_name_keyword

/* Any not-fully-reserved word --- these names can be, eg, role names.
 */
NonReservedWord:
  IDENT
| unreserved_keyword
| col_name_keyword
| type_func_name_keyword

/* Column label --- allowed labels in "AS" clauses.
 * This presently includes *all* Postgres keywords.
 */
ColLabel:
  IDENT                   { $$ = $1 }
| unreserved_keyword      { $$ = $1 }
| col_name_keyword        { $$ = $1 }
| type_func_name_keyword  { $$ = $1 }
| reserved_keyword        { $$ = $1 }

/*
 * Keyword category lists.  Generally, every keyword present in
 * the Postgres grammar should appear in exactly one of these lists.
 */

/* "Unreserved" keywords --- available for use as any kind of name.
 */
unreserved_keyword:
  ABORT_P
| ABSOLUTE_P
| ACCESS
| ACTION
| ADD_P
| ADMIN
| AFTER
| AGGREGATE
| ALSO
| ALTER
| ALWAYS
| ASSERTION
| ASSIGNMENT
| AT
| ATTRIBUTE
| BACKWARD
| BEFORE
| BEGIN_P
| BY
| CACHE
| CALLED
| CASCADE
| CASCADED
| CATALOG_P
| CHAIN
| CHARACTERISTICS
| CHECKPOINT
| CLASS
| CLOSE
| CLUSTER
| COMMENT
| COMMENTS
| COMMIT
| COMMITTED
| CONFIGURATION
| CONFLICT
| CONNECTION
| CONSTRAINTS
| CONTENT_P
| CONTINUE_P
| CONVERSION_P
| COPY
| COST
| CSV
| CUBE
| CURRENT_P
| CURSOR
| CYCLE
| DATA_P
| DATABASE
| DAY_P
| DEALLOCATE
| DECLARE
| DEFAULTS
| DEFERRED
| DEFINER
| DELETE_P
| DELIMITER
| DELIMITERS
| DICTIONARY
| DISABLE_P
| DISCARD
| DOCUMENT_P
| DOMAIN_P
| DOUBLE_P
| DROP
| EACH
| ENABLE_P
| ENCODING
| ENCRYPTED
| ENUM_P
| ESCAPE
| EVENT
| EXCLUDE
| EXCLUDING
| EXCLUSIVE
| EXECUTE
| EXPLAIN
| EXTENSION
| EXTERNAL
| FAMILY
| FILTER
| FIRST_P
| FOLLOWING
| FORCE
| FORWARD
| FUNCTION
| FUNCTIONS
| GLOBAL
| GRANTED
| HANDLER
| HEADER_P
| HOLD
| HOUR_P
| IDENTITY_P
| IF_P
| IMMEDIATE
| IMMUTABLE
| IMPLICIT_P
| IMPORT_P
| INCLUDING
| INCREMENT
| INDEX
| INDEXES
| INHERIT
| INHERITS
| INLINE_P
| INPUT_P
| INSENSITIVE
| INSERT
| INSTEAD
| INVOKER
| ISOLATION
| KEY
| LABEL
| LANGUAGE
| LARGE_P
| LAST_P
| LEAKPROOF
| LEVEL
| LISTEN
| LOAD
| LOCAL
| LOCATION
| LOCK_P
| LOCKED
| LOGGED
| MAPPING
| MATCH
| MATERIALIZED
| MAXVALUE
| MINUTE_P
| MINVALUE
| MODE
| MONTH_P
| MOVE
| NAME_P
| NAMES
| NEXT
| NO
| NOTHING
| NOTIFY
| NOWAIT
| NULLS_P
| OBJECT_P
| OF
| OFF
| OIDS
| OPERATOR
| OPTION
| OPTIONS
| ORDINALITY
| OVER
| OWNED
| OWNER
| PARSER
| PARTIAL
| PARTITION
| PASSING
| PASSWORD
| PLANS
| POLICY
| PRECEDING
| PREPARE
| PREPARED
| PRESERVE
| PRIOR
| PRIVILEGES
| PROCEDURAL
| PROCEDURE
| PROGRAM
| QUOTE
| RANGE
| READ
| REASSIGN
| RECHECK
| RECURSIVE
| REF
| REFRESH
| REINDEX
| RELATIVE_P
| RELEASE
| RENAME
| REPEATABLE
| REPLACE
| REPLICA
| RESET
| RESTART
| RESTRICT
| RETURNS
| REVOKE
| ROLE
| ROLLBACK
| ROLLUP
| ROWS
| RULE
| SAVEPOINT
| SCHEMA
| SCROLL
| SEARCH
| SECOND_P
| SECURITY
| SEQUENCE
| SEQUENCES
| SERIALIZABLE
| SERVER
| SESSION
| SET
| SETS
| SHARE
| SHOW
| SIMPLE
| SKIP
| SNAPSHOT
| SQL_P
| STABLE
| STANDALONE_P
| START
| STATEMENT
| STATISTICS
| STDIN
| STDOUT
| STORAGE
| STRICT_P
| STRIP_P
| SYSID
| SYSTEM_P
| TABLES
| TABLESPACE
| TEMP
| TEMPLATE
| TEMPORARY
| TEXT_P
| TRANSACTION
| TRANSFORM
| TRIGGER
| TRUNCATE
| TRUSTED
| TYPE_P
| TYPES_P
| UNBOUNDED
| UNCOMMITTED
| UNENCRYPTED
| UNKNOWN
| UNLISTEN
| UNLOGGED
| UNTIL
| UPDATE
| VACUUM
| VALID
| VALIDATE
| VALIDATOR
| VALUE_P
| VARYING
| VERSION_P
| VIEW
| VIEWS
| VOLATILE
| WHITESPACE_P
| WITHIN
| WITHOUT
| WORK
| WRAPPER
| WRITE
| XML_P
| YEAR_P
| YES_P
| ZONE


/* Column identifier --- keywords that can be column, table, etc names.
 *
 * Many of these keywords will in fact be recognized as type or function
 * names too; but they have special productions for the purpose, and so
 * can't be treated as "generic" type or function names.
 *
 * The type names appearing here are not usable as function names
 * because they can be followed by '(' in typename productions, which
 * looks too much like a function call for an LR(1) parser.
 */
col_name_keyword:
  BETWEEN
| BIGINT
| BIT
| BOOLEAN_P
| CHAR_P
| CHARACTER
| COALESCE
| DEC
| DECIMAL_P
| EXISTS
| EXTRACT
| FLOAT_P
| GREATEST
| GROUPING
| INOUT
| INT_P
| INTEGER
| INTERVAL
| LEAST
| NATIONAL
| NCHAR
| NONE
| NULLIF
| NUMERIC
| OUT_P
| OVERLAY
| POSITION
| PRECISION
| REAL
| ROW
| SETOF
| SMALLINT
| SUBSTRING
| TIME
| TIMESTAMP
| TREAT
| TRIM
| VALUES
| VARCHAR
| XMLATTRIBUTES
| XMLCONCAT
| XMLELEMENT
| XMLEXISTS
| XMLFOREST
| XMLPARSE
| XMLPI
| XMLROOT
| XMLSERIALIZE

/* Type/function identifier --- keywords that can be type or function names.
 *
 * Most of these are keywords that are used as operators in expressions;
 * in general such keywords can't be column names because they would be
 * ambiguous with variables, but they are unambiguous as function identifiers.
 *
 * Do not include POSITION, SUBSTRING, etc here since they have explicit
 * productions in a_expr to support the goofy SQL9x argument syntax.
 * - thomas 2000-11-28
 */
type_func_name_keyword:
  AUTHORIZATION
| BINARY
| COLLATION
| CONCURRENTLY
| CROSS
| CURRENT_SCHEMA
| FREEZE
| FULL
| ILIKE
| INNER_P
| IS
| ISNULL
| JOIN
| LEFT
| LIKE
| NATURAL
| NOTNULL
| OUTER_P
| OVERLAPS
| RIGHT
| SIMILAR
| TABLESAMPLE
| VERBOSE

/* Reserved keyword --- these keywords are usable only as a ColLabel.
 *
 * Keywords appear here if they could not be distinguished from variable,
 * type, or function names in some contexts.  Don't put things here unless
 * forced to.
 */
reserved_keyword:
  ALL
| ANALYSE
| ANALYZE
| AND
| ANY
| ARRAY
| AS
| ASC
| ASYMMETRIC
| BOTH
| CASE
| CAST
| CHECK
| COLLATE
| COLUMN
| CONSTRAINT
| CREATE
| CURRENT_CATALOG
| CURRENT_DATE
| CURRENT_ROLE
| CURRENT_TIME
| CURRENT_TIMESTAMP
| CURRENT_USER
| DEFAULT
| DEFERRABLE
| DESC
| DISTINCT
| DO
| ELSE
| END_P
| EXCEPT
| FALSE_P
| FETCH
| FOR
| FOREIGN
| FROM
| GRANT
| GROUP_P
| HAVING
| IN_P
| INITIALLY
| INTERSECT
| INTO
| LATERAL_P
| LEADING
| LIMIT
| LOCALTIME
| LOCALTIMESTAMP
| NOT
| NULL_P
| OFFSET
| ON
| ONLY
| OR
| ORDER
| PLACING
| PRIMARY
| REFERENCES
| RETURNING
| SELECT
| SESSION_USER
| SOME
| SYMMETRIC
| TABLE
| THEN
| TO
| TRAILING
| TRUE_P
| UNION
| UNIQUE
| USER
| USING
| VARIADIC
| WHEN
| WHERE
| WINDOW
| WITH

%%

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0

