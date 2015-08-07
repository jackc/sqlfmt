%{

package sqlfmt

%}

%union {
  sqlSelect *SelectStmt
  simpleSelect *SimpleSelect
  fields []Expr
  expr Expr
  str string
  identifiers []string
  fromClause *FromClause
  whereClause *WhereClause
  orderExpr OrderExpr
  orderClause *OrderClause
  groupByClause *GroupByClause
  limitClause *LimitClause
  lockingClause *LockingClause
  lockingItem LockingItem
  boolean bool
  placeholder interface{}
  columnRef ColumnRef
  whenClauses []WhenClause
  whenClause WhenClause
  pgType PgType
  pgTypes []PgType
  valuesRow ValuesRow
  valuesClause ValuesClause
  funcApplication FuncApplication
  funcArgs []FuncArg
  funcArg FuncArg
  filterClause *FilterClause
  relationExpr *RelationExpr
  windowDefinitions []WindowDefinition
  windowDefinition WindowDefinition
  windowSpecification WindowSpecification
  overClause *OverClause
  partitionClause PartitionClause
  frameClause *FrameClause
  frameBound *FrameBound
  arrayExpr ArrayExpr
  anyName AnyName
  indirectionEl IndirectionEl
  indirection Indirection
  iconst IntegerLiteral
  optArrayBounds []IntegerLiteral
}

%type <sqlSelect> top
%type <sqlSelect> SelectStmt
%type <sqlSelect> select_no_parens
%type <sqlSelect> select_with_parens select_clause
%type <simpleSelect> simple_select
%type <valuesClause> values_clause
%type <fields> opt_target_list target_list distinct_clause expr_list
%type <placeholder> opt_all_clause

%type <expr> aliasableExpr
%type <expr> target_el a_expr b_expr c_expr
%type <fromClause> from_clause
%type <identifiers> identifierSeq
%type <expr> joinExpr
%type <whereClause> where_clause
%type <orderExpr> sortby
%type <orderClause> opt_sort_clause sort_clause sortby_list
%type <str> opt_asc_desc opt_nulls_order
%type <placeholder> into_clause
  row_or_rows
  first_or_next
  within_group_clause
  opt_asymmetric

%type <fields> opt_type_modifiers

%type <filterClause> filter_clause

%type <relationExpr> relation_expr


%type <limitClause> select_limit opt_select_limit

%type <expr>
  limit_clause
  offset_clause
  select_limit_value
  select_offset_value
  opt_select_fetch_first_value
  select_offset_value2

%type <lockingClause> opt_for_locking_clause for_locking_clause for_locking_items
%type <lockingItem> for_locking_item
%type <str> for_locking_strength opt_nowait_or_skip

%type <identifiers> locked_rels_list qualified_name_list name_list

%type <indirectionEl> indirection_el
%type <indirection> indirection opt_indirection
%type <str> attr_name qualified_name ColId name param_name

%type <str> MathOp qual_Op qual_all_Op all_Op

%type <groupByClause> group_clause
%type <fields>  group_by_list
%type <expr> group_by_item

%type <expr> having_clause

%type <boolean> all_or_distinct

%type <expr>  SignedIconst Sconst AexprConst
%type <iconst> Iconst opt_float

%type <expr> case_expr case_arg case_default
%type <whenClauses> when_clause_list
%type <whenClause> when_clause

%type <expr> ctext_expr
%type <valuesRow> ctext_expr_list ctext_row

%type <funcApplication> func_application
%type <funcArgs> func_arg_list
%type <funcArg> func_arg_expr
%type <expr> func_expr func_expr_common_subexpr

%type <windowDefinitions> window_definition_list window_clause
%type <windowDefinition> window_definition
%type <windowSpecification> window_specification
%type <overClause> over_clause
%type <str> opt_existing_window_name
%type <partitionClause> opt_partition_clause
%type <frameClause> opt_frame_clause frame_extent
%type <frameBound> frame_bound

%type <arrayExpr> array_expr array_expr_list

%type <columnRef> columnref
%type <anyName> any_name attrs

%type <str>
  ColLabel
  unreserved_keyword
  col_name_keyword
  type_function_name
  type_func_name_keyword
  reserved_keyword
  func_name

%type <pgType> GenericType Numeric Typename SimpleTypename
%type <pgTypes> type_list
%type <optArrayBounds> opt_array_bounds


%token  <str> any_operator


/*
 * Non-keyword token types.  These are hard-wired into the "flex" lexer.
 * They must be listed first so that their numeric codes do not depend on
 * the set of keywords.  PL/pgsql depends on this so that it can share the
 * same lexer.  If you add/change tokens here, fix PL/pgsql to match!
 *
 * DOT_DOT is unused in the core SQL grammar, and so will always provoke
 * parse errors.  It is needed by PL/pgsql.
 */
%token <str>  IDENT FCONST SCONST BCONST XCONST Op
/* ival in PostgreSQL */
%token <str> ICONST PARAM
%token      TYPECAST DOT_DOT COLON_EQUALS EQUALS_GREATER
%token      LESS_EQUALS GREATER_EQUALS NOT_EQUALS


/* ordinary key words in alphabetical order */
%token <str> ABORT_P ABSOLUTE_P ACCESS ACTION ADD_P ADMIN AFTER
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

/*
 * The grammar thinks these are keywords, but they are not in the kwlist.h
 * list and so can never be entered directly.  The filter in parser.c
 * creates these tokens when required (based on looking one token ahead).
 *
 * NOT_LA exists so that productions such as NOT LIKE can be given the same
 * precedence as LIKE; otherwise they'd effectively have the same precedence
 * as NOT, at least with respect to their left-hand subexpression.
 * NULLS_LA and WITH_LA are needed to make the grammar LALR(1).
 */
%token    NOT_LA NULLS_LA WITH_LA

%left OP


/* Precedence: lowest to highest */
%nonassoc SET       /* see relation_expr_opt_alias */
%left   UNION EXCEPT
%left   INTERSECT
%left   OR
%left   AND
%right    NOT
%nonassoc IS ISNULL NOTNULL /* IS sets precedence for IS NULL, etc */
%nonassoc '<' '>' '=' LESS_EQUALS GREATER_EQUALS NOT_EQUALS
%nonassoc BETWEEN IN_P LIKE ILIKE SIMILAR NOT_LA
%nonassoc ESCAPE      /* ESCAPE must be just above LIKE/ILIKE/SIMILAR */
%nonassoc OVERLAPS
%left   POSTFIXOP   /* dummy for postfix Op rules */
/*
 * To support target_el without AS, we must give IDENT an explicit priority
 * between POSTFIXOP and Op.  We can safely assign the same priority to
 * various unreserved keywords as needed to resolve ambiguities (this can't
 * have any bad effects since obviously the keywords will still behave the
 * same as if they weren't keywords).  We need to do this for PARTITION,
 * RANGE, ROWS to support opt_existing_window_name; and for RANGE, ROWS
 * so that they can follow a_expr without creating postfix-operator problems;
 * and for NULL so that it can follow b_expr in ColQualList without creating
 * postfix-operator problems.
 *
 * To support CUBE and ROLLUP in GROUP BY without reserving them, we give them
 * an explicit priority lower than '(', so that a rule with CUBE '(' will shift
 * rather than reducing a conflicting rule that takes CUBE as a function name.
 * Using the same precedence as IDENT seems right for the reasons given above.
 *
 * The frame_bound productions UNBOUNDED PRECEDING and UNBOUNDED FOLLOWING
 * are even messier: since UNBOUNDED is an unreserved keyword (per spec!),
 * there is no principled way to distinguish these from the productions
 * a_expr PRECEDING/FOLLOWING.  We hack this up by giving UNBOUNDED slightly
 * lower precedence than PRECEDING and FOLLOWING.  At present this doesn't
 * appear to cause UNBOUNDED to be treated differently from other unreserved
 * keywords anywhere else in the grammar, but it's definitely risky.  We can
 * blame any funny behavior of UNBOUNDED on the SQL standard, though.
 */
%nonassoc UNBOUNDED   /* ideally should have same precedence as IDENT */
%nonassoc IDENT NULL_P PARTITION RANGE ROWS PRECEDING FOLLOWING CUBE ROLLUP
%left   Op OPERATOR   /* multi-character ops and user-defined operators */
%left   '+' '-'
%left   '*' '/' '%'
%left   '^'
/* Unary Operators */
%left   AT        /* sets precedence for AT TIME ZONE */
%left   COLLATE
%right    UMINUS
%left   '[' ']'
%left   '(' ')'
%left   TYPECAST
%left   '.'
/*
 * These might seem to be low-precedence, but actually they are not part
 * of the arithmetic hierarchy at all in their use as JOIN operators.
 * We make them high-precedence to support their use as function names.
 * They wouldn't be given a precedence at all, were it not that we need
 * left-associativity among the JOIN rules themselves.
 */
%left   JOIN CROSS LEFT FULL RIGHT INNER_P NATURAL
/* kluge to keep xml_whitespace_option from causing shift/reduce conflicts */
%right    PRESERVE STRIP_P


%%

top:
  SelectStmt
  {
    $$ = $1
    yylex.(*sqlLex).stmt = $1
  }

opt_asc_desc:
  ASC          { $$ = "asc" }
| DESC         { $$ = "desc" }
| /*EMPTY*/    { $$ = "" }

opt_nulls_order:
  NULLS_LA FIRST_P    { $$ = "first" }
| NULLS_LA LAST_P     { $$ = "last" }
| /*EMPTY*/           { $$ = "" }

/*
 * Ideally param_name should be ColId, but that causes too many conflicts.
 */
param_name: type_function_name

aliasableExpr:
  a_expr
  {
    $$ = $1
  }
| a_expr AS IDENT
  {
    $$ = AliasedExpr{Expr: $1, Alias: $3}
  }
| a_expr IDENT
  {
    $$ = AliasedExpr{Expr: $1, Alias: $2}
  }


any_name:
  ColId
  {
    $$ = AnyName{$1}
  }
| ColId attrs
  {
    $$ = AnyName{$1}
    $$ = append($$, $2...)
  }

attrs:
  '.' attr_name
  {
    $$ = AnyName{$2}
  }
| attrs '.' attr_name
  {
    $$ = append($1, $3)
  }


/*****************************************************************************
 *
 *  Type syntax
 *    SQL introduces a large amount of type-specific syntax.
 *    Define individual clauses to handle these cases, and use
 *     the generic case to handle regular type-extensible Postgres syntax.
 *    - thomas 1997-10-10
 *
 *****************************************************************************/

Typename:
  SimpleTypename opt_array_bounds
  {
    $$ = $1
    $$.ArrayBounds = $2
  }
| SETOF SimpleTypename opt_array_bounds
  {
    $$ = $2
    $$.Setof = true
    $$.ArrayBounds = $3
  }
      /* SQL standard syntax, currently only one-dimensional */
/* TODO
      | SimpleTypename ARRAY '[' Iconst ']'
      | SETOF SimpleTypename ARRAY '[' Iconst ']'
      | SimpleTypename ARRAY
      | SETOF SimpleTypename ARRAY
*/

opt_array_bounds:
  opt_array_bounds '[' ']'
  {
    $$ = append($1, "")
  }
| opt_array_bounds '[' Iconst ']'
  {
    $$ = append($1, $3)
  }
| /*EMPTY*/
  {  $$ = nil }

SimpleTypename:
  GenericType      { $$ = $1 }
| Numeric          { $$ = $1 }
/* TODO
| Bit              { $$ = $1 }
| Character        { $$ = $1 }
| ConstDatetime    { $$ = $1 }
| ConstInterval opt_interval
  {
    $$ = $1;
    $$->typmods = $2;
  }
| ConstInterval '(' Iconst ')'
  {
    $$ = $1;
    $$->typmods = list_make2(makeIntConst(INTERVAL_FULL_RANGE, -1),
                 makeIntConst($3, @3));
  }
*/

/* We have a separate ConstTypename to allow defaulting fixed-length
 * types such as CHAR() and BIT() to an unspecified length.
 * SQL9x requires that these default to a length of one, but this
 * makes no sense for constructs like CHAR 'hi' and BIT '0101',
 * where there is an obvious better choice to make.
 * Note that ConstInterval is not included here since it must
 * be pushed up higher in the rules to accommodate the postfix
 * options (e.g. INTERVAL '1' YEAR). Likewise, we have to handle
 * the generic-type-name case in AExprConst to avoid premature
 * reduce/reduce conflicts against function names.
 */
/* TODO
ConstTypename:
      Numeric                 { $$ = $1; }
      | ConstBit                { $$ = $1; }
      | ConstCharacter            { $$ = $1; }
      | ConstDatetime             { $$ = $1; }
*/

/*
 * GenericType covers all type names that don't have special syntax mandated
 * by the standard, including qualified names.  We also allow type modifiers.
 * To avoid parsing conflicts against function invocations, the modifiers
 * have to be shown as expr_list here, but parse analysis will only accept
 * constants for them.
 */
GenericType:
  type_function_name opt_type_modifiers
  {
    $$ = PgType{Name: $1, TypeMods: $2}
  }
/* TODO
| type_function_name attrs opt_type_modifiers
  {
    panic("TODO")
  }
*/

opt_type_modifiers:
  '(' expr_list ')'   { $$ = $2 }
| /* EMPTY */         { $$ = nil }

/*
 * SQL numeric data types
 */
Numeric:
  INT_P
  {
    $$ = PgType{Name: "int"}
  }
| INTEGER
  {
    $$ = PgType{Name: "integer"}
  }
| SMALLINT
  {
    $$ = PgType{Name: "smallint"}
  }
| BIGINT
  {
    $$ = PgType{Name: "bigint"}
  }
| REAL
  {
    $$ = PgType{Name: "real"}
  }
| FLOAT_P opt_float
  {
    $$ = PgType{Name: "float"}
    if $2 != IntegerLiteral("") {
      $$.TypeMods = []Expr{$2}
    }
  }
| DOUBLE_P PRECISION
  {
    $$ = PgType{Name: "double precision"}
  }
| DECIMAL_P opt_type_modifiers
  {
    $$ = PgType{Name: "decimal", TypeMods: $2}
  }
| DEC opt_type_modifiers
  {
    $$ = PgType{Name: "dec", TypeMods: $2}
  }
| NUMERIC opt_type_modifiers
  {
    $$ = PgType{Name: "numeric", TypeMods: $2}
  }
| BOOLEAN_P
  {
    $$ = PgType{Name: "bool"}
  }

opt_float:
  '(' Iconst ')'
  {
    $$ = $2
  }
| /*EMPTY*/
  {
    $$ = IntegerLiteral("")
  }
/* TODO
Bit
ConstBit
BitWithLength
BitWithoutLength
Character
ConstCharacter
CharacterWithLength
CharacterWithoutLength
character
opt_varying
opt_charset
ConstDatetime
ConstInterval
opt_timezone
opt_interval
interval_second
*/


/*****************************************************************************
 *
 *  expression grammar
 *
 *****************************************************************************/

/*
 * General expressions
 * This is the heart of the expression syntax.
 *
 * We have two expression types: a_expr is the unrestricted kind, and
 * b_expr is a subset that must be used in some places to avoid shift/reduce
 * conflicts.  For example, we can't do BETWEEN as "BETWEEN a_expr AND a_expr"
 * because that use of AND conflicts with AND as a boolean operator.  So,
 * b_expr is used in BETWEEN and we remove boolean keywords from b_expr.
 *
 * Note that '(' a_expr ')' is a b_expr, so an unrestricted expression can
 * always be used by surrounding it with parens.
 *
 * c_expr is all the productions that are common to a_expr and b_expr;
 * it's factored out just to eliminate redundant coding.
 *
 * Be careful of productions involving more than one terminal token.
 * By default, bison will assign such productions the precedence of their
 * last terminal, but in nearly all cases you want it to be the precedence
 * of the first terminal instead; otherwise you will not get the behavior
 * you expect!  So we use %prec annotations freely to set precedences.
 */
a_expr:
  c_expr
| a_expr TYPECAST Typename
  {
    $$ = TypecastExpr{Expr: $1, Typename: $3}
  }
| a_expr COLLATE any_name
  {
    $$ = CollateExpr{Expr: $1, Collation: $3}
  }
| a_expr AT TIME ZONE a_expr      %prec AT
  {
    $$ = AtTimeZoneExpr{Expr: $1, TimeZone: $5}
  }
/*
* These operators must be called out explicitly in order to make use
* of bison's automatic operator-precedence handling.  All other
* operator names are handled by the generic productions using "Op",
* below; and all those operators will have the same precedence.
*
* If you add more explicitly-known operators, be sure to add them
* also to b_expr and to the MathOp list below.
*/
| '+' a_expr          %prec UMINUS
  {
    $$ = UnaryExpr{Operator: "+", Expr: $2}
  }
| '-' a_expr          %prec UMINUS
  {
    $$ = UnaryExpr{Operator: "-", Expr: $2}
  }
| a_expr '+' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "+", Right: $3}
  }
| a_expr '-' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "-", Right: $3}
  }
| a_expr '*' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "*", Right: $3}
  }
| a_expr '/' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "/", Right: $3}
  }
| a_expr '%' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "%", Right: $3}
  }
| a_expr '^' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "^", Right: $3}
  }
| a_expr '<' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "<", Right: $3}
  }
| a_expr '>' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: ">", Right: $3}
  }
| a_expr '=' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "=", Right: $3}
  }
| a_expr LESS_EQUALS a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "<=", Right: $3}
  }
| a_expr GREATER_EQUALS a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: ">=", Right: $3}
  }
| a_expr NOT_EQUALS a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "!=", Right: $3}
  }
| a_expr qual_Op a_expr       %prec Op
  {
    $$ = BinaryExpr{Left: $1, Operator: $2, Right: $3}
  }
| qual_Op a_expr          %prec Op
  {
    $$ = UnaryExpr{Operator: $1, Expr: $2}
  }
| a_expr qual_Op          %prec POSTFIXOP
  {
    $$ = PostfixExpr{Expr: $1, Operator: $2}
  }
| a_expr AND a_expr
  {
    $$ = BooleanExpr{Left: $1, Operator: "and", Right: $3}
  }
| a_expr OR a_expr
  {
    $$ = BooleanExpr{Left: $1, Operator: "or", Right: $3}
  }
| NOT a_expr
  {
    $$ = NotExpr{Expr: $2}
  }
| NOT_LA a_expr           %prec NOT
  {
    $$ = NotExpr{Expr: $2}
  }
| a_expr LIKE a_expr
  {
    $$ = TextOpWithEscapeExpr{Left: $1, Operator: "like", Right: $3}
  }
| a_expr LIKE a_expr ESCAPE a_expr          %prec LIKE
  {
    $$ = TextOpWithEscapeExpr{Left: $1, Operator: "like", Right: $3, Escape: $5}
  }
| a_expr NOT_LA LIKE a_expr             %prec NOT_LA
  {
    $$ = TextOpWithEscapeExpr{Left: $1, Operator: "not like", Right: $4}
  }
| a_expr NOT_LA LIKE a_expr ESCAPE a_expr     %prec NOT_LA
  {
    $$ = TextOpWithEscapeExpr{Left: $1, Operator: "not like", Right: $4, Escape: $6}
  }
| a_expr ILIKE a_expr
  {
    $$ = TextOpWithEscapeExpr{Left: $1, Operator: "ilike", Right: $3}
  }
| a_expr ILIKE a_expr ESCAPE a_expr         %prec ILIKE
  {
    $$ = TextOpWithEscapeExpr{Left: $1, Operator: "ilike", Right: $3, Escape: $5}
  }
| a_expr NOT_LA ILIKE a_expr            %prec NOT_LA
  {
    $$ = TextOpWithEscapeExpr{Left: $1, Operator: "not ilike", Right: $4}
  }
| a_expr NOT_LA ILIKE a_expr ESCAPE a_expr      %prec NOT_LA
  {
    $$ = TextOpWithEscapeExpr{Left: $1, Operator: "not ilike", Right: $4, Escape: $6}
  }

| a_expr SIMILAR TO a_expr              %prec SIMILAR
  {
    $$ = TextOpWithEscapeExpr{Left: $1, Operator: "similar to", Right: $4}
  }
| a_expr SIMILAR TO a_expr ESCAPE a_expr      %prec SIMILAR
  {
    $$ = TextOpWithEscapeExpr{Left: $1, Operator: "similar to", Right: $4, Escape: $6}
  }
| a_expr NOT_LA SIMILAR TO a_expr         %prec NOT_LA
  {
    $$ = TextOpWithEscapeExpr{Left: $1, Operator: "not similar to", Right: $5}
  }
| a_expr NOT_LA SIMILAR TO a_expr ESCAPE a_expr   %prec NOT_LA
  {
    $$ = TextOpWithEscapeExpr{Left: $1, Operator: "not similar to", Right: $5, Escape: $7}
  }
/* NullTest clause
 * Define SQL-style Null test clause.
 * Allow two forms described in the standard:
 *  a IS NULL
 *  a IS NOT NULL
 * Allow two SQL extensions
 *  a ISNULL
 *  a NOTNULL
 */
| a_expr IS NULL_P              %prec IS
  {
    $$ = IsNullExpr{Expr: $1}
  }
| a_expr ISNULL
  {
    $$ = IsNullExpr{Expr: $1}
  }
| a_expr IS NOT NULL_P            %prec IS
  {
    $$ = IsNullExpr{Expr: $1, Not: true}
  }
| a_expr NOTNULL
  {
    $$ = IsNullExpr{Expr: $1, Not: true}
  }
/* TODO
      | row OVERLAPS row
*/
| a_expr IS TRUE_P              %prec IS
  {
    $$ = IsBoolOpExpr{Expr: $1, Op: "true"}
  }
| a_expr IS NOT TRUE_P            %prec IS
  {
    $$ = IsBoolOpExpr{Expr: $1, Not: true, Op: "true"}
  }
| a_expr IS FALSE_P             %prec IS
  {
    $$ = IsBoolOpExpr{Expr: $1, Op: "false"}
  }
| a_expr IS NOT FALSE_P           %prec IS
  {
    $$ = IsBoolOpExpr{Expr: $1, Not: true, Op: "false"}
  }
| a_expr IS UNKNOWN             %prec IS
  {
    $$ = IsBoolOpExpr{Expr: $1, Op: "unknown"}
  }
| a_expr IS NOT UNKNOWN           %prec IS
  {
    $$ = IsBoolOpExpr{Expr: $1, Not: true, Op: "unknown"}
  }
| a_expr IS DISTINCT FROM a_expr      %prec IS
  {
    $$ = BinaryExpr{Left: $1, Operator: "is distinct from", Right: $5}
  }
| a_expr IS NOT DISTINCT FROM a_expr    %prec IS
  {
    $$ = BinaryExpr{Left: $1, Operator: "is not distinct from", Right: $6}
  }
| a_expr IS OF '(' type_list ')'      %prec IS
  {
    $$ = IsOfExpr{Expr: $1, Types: $5}
  }
| a_expr IS NOT OF '(' type_list ')'    %prec IS
  {
    $$ = IsOfExpr{Expr: $1, Not: true, Types: $6}
  }
| a_expr BETWEEN opt_asymmetric b_expr AND a_expr   %prec BETWEEN
  {
    $$ = BetweenExpr{Expr: $1, Left: $4, Right: $6}
  }
| a_expr NOT_LA BETWEEN opt_asymmetric b_expr AND a_expr %prec NOT_LA
  {
    $$ = BetweenExpr{Expr: $1, Not: true, Left: $5, Right: $7}
  }
| a_expr BETWEEN SYMMETRIC b_expr AND a_expr      %prec BETWEEN
  {
    $$ = BetweenExpr{Expr: $1, Symmetric: true, Left: $4, Right: $6}
  }
| a_expr NOT_LA BETWEEN SYMMETRIC b_expr AND a_expr   %prec NOT_LA
  {
    $$ = BetweenExpr{Expr: $1, Not: true, Symmetric: true, Left: $5, Right: $7}
  }
/* TODO
      | a_expr IN_P in_expr
      | a_expr NOT_LA IN_P in_expr            %prec NOT_LA
      | a_expr subquery_Op sub_type select_with_parens  %prec Op
      | a_expr subquery_Op sub_type '(' a_expr ')'    %prec Op
      | a_expr IS DOCUMENT_P          %prec IS
      | a_expr IS NOT DOCUMENT_P        %prec IS
*/

/*
 * Restricted expressions
 *
 * b_expr is a subset of the complete expression syntax defined by a_expr.
 *
 * Presently, AND, NOT, IS, and IN are the a_expr keywords that would
 * cause trouble in the places where b_expr is used.  For simplicity, we
 * just eliminate all the boolean-keyword-operator productions from b_expr.
 */
b_expr:
  c_expr
  {
    $$ = $1
  }
| b_expr TYPECAST Typename
  {
    $$ = TypecastExpr{Expr: $1, Typename: $3}
  }
| '+' b_expr          %prec UMINUS
  {
    $$ = UnaryExpr{Operator: "+", Expr: $2}
  }
| '-' b_expr          %prec UMINUS
  {
    $$ = UnaryExpr{Operator: "-", Expr: $2}
  }
| b_expr '+' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "+", Right: $3}
  }
| b_expr '-' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "-", Right: $3}
  }
| b_expr '*' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "*", Right: $3}
  }
| b_expr '/' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "/", Right: $3}
  }
| b_expr '%' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "%", Right: $3}
  }
| b_expr '^' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "^", Right: $3}
  }
| b_expr '<' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "<", Right: $3}
  }
| b_expr '>' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: ">", Right: $3}
  }
| b_expr '=' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "=", Right: $3}
  }
| b_expr LESS_EQUALS b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "<=", Right: $3}
  }
| b_expr GREATER_EQUALS b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: ">=", Right: $3}
  }
| b_expr NOT_EQUALS b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: "!=", Right: $3}
  }
| b_expr qual_Op b_expr       %prec Op
  {
    $$ = BinaryExpr{Left: $1, Operator: $2, Right: $3}
  }
| qual_Op b_expr          %prec Op
  {
    $$ = UnaryExpr{Operator: $1, Expr: $2}
  }
| b_expr qual_Op          %prec POSTFIXOP
  {
    $$ = PostfixExpr{Expr: $1, Operator: $2}
  }
| b_expr IS DISTINCT FROM b_expr    %prec IS
  {
    $$ = BinaryExpr{Left: $1, Operator: "is distinct from", Right: $5}
  }
| b_expr IS NOT DISTINCT FROM b_expr  %prec IS
  {
    $$ = BinaryExpr{Left: $1, Operator: "is not distinct from", Right: $6}
  }
| b_expr IS OF '(' type_list ')'    %prec IS
  {
    $$ = IsOfExpr{Expr: $1, Types: $5}
  }
| b_expr IS NOT OF '(' type_list ')'  %prec IS
  {
    $$ = IsOfExpr{Expr: $1, Not: true, Types: $6}
  }
/* TODO
| b_expr IS DOCUMENT_P          %prec IS
  {
    $$ = makeXmlExpr(IS_DOCUMENT, NULL, NIL,
             list_make1($1), @2);
  }
| b_expr IS NOT DOCUMENT_P        %prec IS
  {
    $$ = makeNotExpr(makeXmlExpr(IS_DOCUMENT, NULL, NIL,
                   list_make1($1), @2),
             @2);
  }
*/

/*
 * Productions that can be used in both a_expr and b_expr.
 *
 * Note: productions that refer recursively to a_expr or b_expr mostly
 * cannot appear here.  However, it's OK to refer to a_exprs that occur
 * inside parentheses, such as function arguments; that cannot introduce
 * ambiguity to the b_expr syntax.
 */
c_expr:
  columnref   { $$ = $1 }
| AexprConst  { $$ = $1 }
/* TODO
| PARAM opt_indirection
*/

| '(' a_expr ')' opt_indirection
  {
    $$ = ParenExpr{Expr: $2, Indirection: $4}
  }
| case_expr { $$ = $1 }
| func_expr { $$ = $1 }
| select_with_parens      %prec UMINUS
  {
    $$ = $1
  }
| select_with_parens indirection
  {
    $1.ParenWrapped = false
    $$ = ParenExpr{Expr: $1, Indirection: $2}
  }
| EXISTS select_with_parens
  {
    $$ = ExistsExpr(*$2)
  }
| ARRAY select_with_parens
  {
    $$ = ArraySubselect(*$2)
  }
| ARRAY array_expr {
  $$ = ArrayConstructorExpr($2)
}
/* TODO
| explicit_row
| implicit_row
| GROUPING '(' expr_list ')'
*/



func_application:
  func_name '(' ')'
  {
    $$ = FuncApplication{Name: $1}
  }
| func_name '(' func_arg_list opt_sort_clause ')'
  {
    $$ = FuncApplication{Name: $1, Args: $3, OrderClause: $4}
  }
/* TODO
      | func_name '(' VARIADIC func_arg_expr opt_sort_clause ')'
      | func_name '(' func_arg_list ',' VARIADIC func_arg_expr opt_sort_clause ')'
*/
| func_name '(' ALL func_arg_list opt_sort_clause ')'
  {
    $$ = FuncApplication{Name: $1, Args: $4, OrderClause: $5}
  }
| func_name '(' DISTINCT func_arg_list opt_sort_clause ')'
  {
    $$ = FuncApplication{Name: $1, Distinct: true, Args: $4, OrderClause: $5}
  }
| func_name '(' '*' ')'
  {
    $$ = FuncApplication{Name: $1, Star: true}
  }


/*
 * func_expr and its cousin func_expr_windowless are split out from c_expr just
 * so that we have classifications for "everything that is a function call or
 * looks like one".  This isn't very important, but it saves us having to
 * document which variants are legal in places like "FROM function()" or the
 * backwards-compatible functional-index syntax for CREATE INDEX.
 * (Note that many of the special SQL functions wouldn't actually make any
 * sense as functional index entries, but we ignore that consideration here.)
 */
func_expr:
  func_application /* TODO within_group_clause */ filter_clause over_clause
  {
    $$ = &FuncExpr{FuncApplication: $1, FilterClause: $2, OverClause: $3}
  }
| func_expr_common_subexpr
  {
    $$ = $1
  }





/*
 * Special expressions that are considered to be functions.
 */
func_expr_common_subexpr:
/* TODO
COLLATION FOR '(' a_expr ')'
| CURRENT_DATE
| CURRENT_TIME
| CURRENT_TIME '(' Iconst ')'
| CURRENT_TIMESTAMP
| CURRENT_TIMESTAMP '(' Iconst ')'
| LOCALTIME
| LOCALTIME '(' Iconst ')'
| LOCALTIMESTAMP
| LOCALTIMESTAMP '(' Iconst ')'
| CURRENT_ROLE
| CURRENT_USER
| SESSION_USER
| USER
| CURRENT_CATALOG
| CURRENT_SCHEMA
| */ CAST '(' a_expr AS Typename ')'
  {
    $$ = CastFunc{Expr: $3, Type: $5}
  }
/* TODO
| EXTRACT '(' extract_list ')'
| OVERLAY '(' overlay_list ')'
| POSITION '(' position_list ')'
| SUBSTRING '(' substr_list ')'
| TREAT '(' a_expr AS Typename ')'
| TRIM '(' BOTH trim_list ')'
| TRIM '(' LEADING trim_list ')'
| TRIM '(' TRAILING trim_list ')'
| TRIM '(' trim_list ')'
| NULLIF '(' a_expr ',' a_expr ')'
*/
| COALESCE '(' expr_list ')'
{
  fa := FuncApplication{Name: "coalesce"}
  for _, e := range $3 {
    fa.Args = append(fa.Args, FuncArg{Expr: e})
  }
  $$ = fa
}
| GREATEST '(' expr_list ')'
{
  fa := FuncApplication{Name: "greatest"}
  for _, e := range $3 {
    fa.Args = append(fa.Args, FuncArg{Expr: e})
  }
  $$ = fa
}
| LEAST '(' expr_list ')'
{
  fa := FuncApplication{Name: "least"}
  for _, e := range $3 {
    fa.Args = append(fa.Args, FuncArg{Expr: e})
  }
  $$ = fa
}
| XMLCONCAT '(' expr_list ')'
{
  fa := FuncApplication{Name: "xmlconcat"}
  for _, e := range $3 {
    fa.Args = append(fa.Args, FuncArg{Expr: e})
  }
  $$ = fa
}
/* TODO
| XMLELEMENT '(' NAME_P ColLabel ')'
| XMLELEMENT '(' NAME_P ColLabel ',' xml_attributes ')'
| XMLELEMENT '(' NAME_P ColLabel ',' expr_list ')'
| XMLELEMENT '(' NAME_P ColLabel ',' xml_attributes ',' expr_list ')'
| XMLEXISTS '(' c_expr xmlexists_argument ')'
| XMLFOREST '(' xml_attribute_list ')'
| XMLPARSE '(' document_or_content a_expr xml_whitespace_option ')'
| XMLPI '(' NAME_P ColLabel ')'
| XMLPI '(' NAME_P ColLabel ',' a_expr ')'
| XMLROOT '(' a_expr ',' xml_root_version opt_xml_root_standalone ')'
| XMLSERIALIZE '(' document_or_content a_expr AS SimpleTypename ')'
*/




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
| aliasableExpr JOIN aliasableExpr ON a_expr
  {
    $$ = JoinExpr{Left: $1, Join: "join", Right: $3, On: $5}
  }

from_clause:
  FROM aliasableExpr
  {
    $$ = &FromClause{Expr: $2}
  }
| FROM joinExpr
  {
    $$ = &FromClause{Expr: $2}
  }
| /*EMPTY*/  { $$ = nil }


opt_nowait_or_skip:
  NOWAIT        { $$ = "nowait" }
| SKIP LOCKED   { $$ = "skip locked" }
| /*EMPTY*/     { $$ = "" }


expr_list:
  a_expr
  {
    $$ = []Expr{$1}
  }
| expr_list ',' a_expr
  {
    $$ = append($1, $3)
  }

/* function arguments can have names */
func_arg_list:
  func_arg_expr
  {
    $$ = []FuncArg{$1}
  }
  | func_arg_list ',' func_arg_expr
  {
    $$ = append($1, $3)
  }

func_arg_expr:
  a_expr
  {
    $$ = FuncArg{Expr: $1}
  }
| param_name COLON_EQUALS a_expr
  {
    $$ = FuncArg{Name: $1, NameOp: ":=", Expr: $3}
  }
| param_name EQUALS_GREATER a_expr
  {
    $$ = FuncArg{Name: $1, NameOp: "=>", Expr: $3}
  }

type_list:
  Typename
  {
    $$ = []PgType{$1}
  }
| type_list ',' Typename
  {
    $$ = append($1, $3)
  }

array_expr:
  '[' expr_list ']'
  {
    $$ = ArrayExpr($2)
  }
| '[' array_expr_list ']'
  {
    $$ = $2
  }
| '[' ']'
  {
    $$ = ArrayExpr{}
  }

array_expr_list:
  array_expr
  {
    $$ = ArrayExpr{$1}
  }
| array_expr_list ',' array_expr
  {
    $$ = append($1, $3)
  }

/*****************************************************************************
 *
 *    QUERY:
 *        SELECT STATEMENTS
 *
 *****************************************************************************/

/* A complete SELECT statement looks like this.
 *
 * The rule returns either a single SelectStmt node or a tree of them,
 * representing a set-operation tree.
 *
 * There is an ambiguity when a sub-SELECT is within an a_expr and there
 * are excess parentheses: do the parentheses belong to the sub-SELECT or
 * to the surrounding a_expr?  We don't really care, but bison wants to know.
 * To resolve the ambiguity, we are careful to define the grammar so that
 * the decision is staved off as long as possible: as long as we can keep
 * absorbing parentheses into the sub-SELECT, we will do so, and only when
 * it's no longer possible to do that will we decide that parens belong to
 * the expression.  For example, in "SELECT (((SELECT 2)) + 3)" the extra
 * parentheses are treated as part of the sub-select.  The necessity of doing
 * it that way is shown by "SELECT (((SELECT 2)) UNION SELECT 2)".  Had we
 * parsed "((SELECT 2))" as an a_expr, it'd be too late to go back to the
 * SELECT viewpoint when we see the UNION.
 *
 * This approach is implemented by defining a nonterminal select_with_parens,
 * which represents a SELECT with at least one outer layer of parentheses,
 * and being careful to use select_with_parens, never '(' SelectStmt ')',
 * in the expression grammar.  We will then have shift-reduce conflicts
 * which we can resolve in favor of always treating '(' <select> ')' as
 * a select_with_parens.  To resolve the conflicts, the productions that
 * conflict with the select_with_parens productions are manually given
 * precedences lower than the precedence of ')', thereby ensuring that we
 * shift ')' (and then reduce to select_with_parens) rather than trying to
 * reduce the inner <select> nonterminal to something else.  We use UMINUS
 * precedence for this, which is a fairly arbitrary choice.
 *
 * To be able to define select_with_parens itself without ambiguity, we need
 * a nonterminal select_no_parens that represents a SELECT structure with no
 * outermost parentheses.  This is a little bit tedious, but it works.
 *
 * In non-expression contexts, we use SelectStmt which can represent a SELECT
 * with or without outer parentheses.
 */

SelectStmt:
  select_no_parens   %prec UMINUS
| select_with_parens %prec UMINUS


select_with_parens:
  '(' select_no_parens ')'
  {
    $2.ParenWrapped = true
    $$ = $2
  }
| '(' select_with_parens ')'      { $$ = $2 }

select_no_parens:
  simple_select
  {
    ss := &SelectStmt{}
    ss.SimpleSelect = *$1
    $$ = ss
  }
| select_clause sort_clause
  {
    $1.OrderClause = $2
    $$ = $1
  }
| select_clause opt_sort_clause for_locking_clause opt_select_limit
  {
    $1.OrderClause = $2
    $1.LockingClause = $3
    $1.LimitClause = $4
    $$ = $1;
  }
| select_clause opt_sort_clause select_limit opt_for_locking_clause
  {
    $1.OrderClause = $2
    $1.LimitClause = $3
    $1.LockingClause = $4
    $$ = $1;
  }

select_clause:
  simple_select
  {
    ss := &SelectStmt{}
    ss.SimpleSelect = *$1
    $$ = ss
  }
| select_with_parens

/*
 * This rule parses SELECT statements that can appear within set operations,
 * including UNION, INTERSECT and EXCEPT.  '(' and ')' can be used to specify
 * the ordering of the set operations.  Without '(' and ')' we want the
 * operations to be ordered per the precedence specs at the head of this file.
 *
 * As with select_no_parens, simple_select cannot have outer parentheses,
 * but can have parenthesized subclauses.
 *
 * Note that sort clauses cannot be included at this level --- SQL requires
 *    SELECT foo UNION SELECT bar ORDER BY baz
 * to be parsed as
 *    (SELECT foo UNION SELECT bar) ORDER BY baz
 * not
 *    SELECT foo UNION (SELECT bar ORDER BY baz)
 * Likewise for WITH, FOR UPDATE and LIMIT.  Therefore, those clauses are
 * described as part of the select_no_parens production, not simple_select.
 * This does not limit functionality, because you can reintroduce these
 * clauses inside parentheses.
 *
 * NOTE: only the leftmost component SelectStmt should have INTO.
 * However, this is not checked by the grammar; parse analysis must check it.
 */
simple_select:
  SELECT opt_all_clause opt_target_list
  into_clause from_clause where_clause
  group_clause having_clause window_clause
    {
      ss := &SimpleSelect{}
      ss.TargetList = $3
      ss.FromClause = $5
      ss.WhereClause = $6
      ss.GroupByClause = $7
      ss.HavingClause = $8
      ss.WindowClause = $9
      $$ = ss
    }
| SELECT distinct_clause target_list
  into_clause from_clause where_clause
  group_clause having_clause window_clause
  {
    ss := &SimpleSelect{}
    ss.DistinctList = $2
    ss.TargetList = $3
    ss.FromClause = $5
    ss.WhereClause = $6
    ss.GroupByClause = $7
    ss.HavingClause = $8
    ss.WindowClause = $9
    $$ = ss
  }
| values_clause
  {
    ss := &SimpleSelect{}
    ss.ValuesClause = $1
    $$ = ss
  }
| TABLE relation_expr
  {
    ss := &SimpleSelect{}
    ss.Table = $2
    $$ = ss
  }
| select_clause UNION all_or_distinct select_clause
  {
    ss := &SimpleSelect{}
    ss.LeftSelect = $1
    ss.SetOp = "union"
    ss.SetAll = $3
    ss.RightSelect = $4
    $$ = ss
  }
| select_clause INTERSECT all_or_distinct select_clause
  {
    ss := &SimpleSelect{}
    ss.LeftSelect = $1
    ss.SetOp = "intersect"
    ss.SetAll = $3
    ss.RightSelect = $4
    $$ = ss
  }
| select_clause EXCEPT all_or_distinct select_clause
  {
    ss := &SimpleSelect{}
    ss.LeftSelect = $1
    ss.SetOp = "except"
    ss.SetAll = $3
    ss.RightSelect = $4
    $$ = ss
  }



into_clause:
/* TODO      INTO OptTempTableName
        {
          $$ = makeNode(IntoClause);
          $$->rel = $2;
          $$->colNames = NIL;
          $$->options = NIL;
          $$->onCommit = ONCOMMIT_NOOP;
          $$->tableSpaceName = NULL;
          $$->viewQuery = NULL;
          $$->skipData = false;
        }
      | */ /*EMPTY*/
        { $$ = nil }



all_or_distinct:
  ALL                   { $$ = true }
      | DISTINCT                { $$ = false }
      | /*EMPTY*/               { $$ = false }
    ;


distinct_clause:
  DISTINCT { $$ = make([]Expr, 0) }
| DISTINCT ON '(' expr_list ')'     { $$ = $4; }

opt_all_clause:
  ALL        { $$ = nil }
| /*EMPTY*/  { $$ = nil }

opt_sort_clause:
  sort_clause  { $$ = $1 }
| /*EMPTY*/    { $$ = nil }

sort_clause:
  ORDER BY sortby_list  { $$ = $3 }

sortby_list:
  sortby
  {
    $$ = &OrderClause{Exprs: []OrderExpr{$1}}
  }
| sortby_list ',' sortby
  {
    $1.Exprs = append($1.Exprs, $3)
    $$ = $1
  }


sortby:
a_expr USING qual_all_Op opt_nulls_order
  {
    $$ = OrderExpr{Expr: $1, Using: $3, Nulls: $4}
  }
| a_expr opt_asc_desc opt_nulls_order
  {
    $$ = OrderExpr{Expr: $1, Order: $2, Nulls: $3}
  }

/*
 * This syntax for group_clause tries to follow the spec quite closely.
 * However, the spec allows only column references, not expressions,
 * which introduces an ambiguity between implicit row constructors
 * (a,b) and lists of column references.
 *
 * We handle this by using the a_expr production for what the spec calls
 * <ordinary grouping set>, which in the spec represents either one column
 * reference or a parenthesized list of column references. Then, we check the
 * top node of the a_expr to see if it's an implicit RowExpr, and if so, just
 * grab and use the list, discarding the node. (this is done in parse analysis,
 * not here)
 *
 * (we abuse the row_format field of RowExpr to distinguish implicit and
 * explicit row constructors; it's debatable if anyone sanely wants to use them
 * in a group clause, but if they have a reason to, we make it possible.)
 *
 * Each item in the group_clause list is either an expression tree or a
 * GroupingSet node of some type.
 */
group_clause:
  GROUP_P BY group_by_list  { $$ = &GroupByClause{Exprs: $3} }
| /*EMPTY*/               { $$ = nil }

group_by_list:
  group_by_item
  {
    $$ = []Expr{$1}
  }
| group_by_list ',' group_by_item
  {
    $$ = append($1, $3)
  }

group_by_item:
  a_expr
/* TODO      | empty_grouping_set          { $$ = $1; }
      | cube_clause             { $$ = $1; }
      | rollup_clause             { $$ = $1; }
      | grouping_sets_clause          { $$ = $1; }
*/

having_clause:
  HAVING a_expr  { $$ = $2 }
| /*EMPTY*/      { $$ = nil }

for_locking_clause:
  for_locking_items   { $$ = $1 }
| FOR READ ONLY       { $$ = nil }

opt_for_locking_clause:
  for_locking_clause  { $$ = $1 }
| /* EMPTY */         { $$ = nil }

for_locking_items:
  for_locking_item
  {
    $$ = &LockingClause{Locks: []LockingItem{$1}}
  }
| for_locking_items for_locking_item
  {
    $1.Locks = append($1.Locks, $2)
    $$ = $1
  }

for_locking_item:
  for_locking_strength locked_rels_list opt_nowait_or_skip
  {
    $$ = LockingItem{Strength: $1, LockedRels: $2, WaitPolicy: $3}
  }

for_locking_strength:
FOR UPDATE            { $$ = "update" }
| FOR NO KEY UPDATE   { $$ = "no key update" }
| FOR SHARE           { $$ = "share" }
| FOR KEY SHARE       { $$ = "key share" }

locked_rels_list:
  OF qualified_name_list  { $$ = $2 }
| /* EMPTY */             { $$ = nil }


/*
 * Window Definitions
 */
window_clause:
  WINDOW window_definition_list
  {
    $$ = $2
  }
| /*EMPTY*/               { $$ = nil }

window_definition_list:
  window_definition
  {
    $$ = []WindowDefinition{$1}
  }
| window_definition_list ',' window_definition
  {
    $$ = append($1, $3)
  }

window_definition:
  ColId AS window_specification
  {
    $$ = WindowDefinition{Name: $1, Specification: $3}
  }

over_clause:
  OVER window_specification
  {
    spec := $2
    $$ = &OverClause{Specification: &spec}
  }
| OVER ColId
  {
    $$ = &OverClause{Name: $2}
  }
| /*EMPTY*/ { $$ = nil }

window_specification:
  '(' opt_existing_window_name opt_partition_clause opt_sort_clause opt_frame_clause ')'
  {
    $$ = WindowSpecification{ExistingName: $2, PartitionClause: $3, OrderClause: $4, FrameClause: $5}
  }

/*
 * If we see PARTITION, RANGE, or ROWS as the first token after the '('
 * of a window_specification, we want the assumption to be that there is
 * no existing_window_name; but those keywords are unreserved and so could
 * be ColIds.  We fix this by making them have the same precedence as IDENT
 * and giving the empty production here a slightly higher precedence, so
 * that the shift/reduce conflict is resolved in favor of reducing the rule.
 * These keywords are thus precluded from being an existing_window_name but
 * are not reserved for any other purpose.
 */
opt_existing_window_name:
  ColId           { $$ = $1 }
| /*EMPTY*/       %prec Op    { $$ = "" }

opt_partition_clause:
  PARTITION BY expr_list    { $$ = PartitionClause($3) }
| /*EMPTY*/               { $$ = nil }

/*
 * For frame clauses, we return a WindowDef, but only some fields are used:
 * frameOptions, startOffset, and endOffset.
 *
 * This is only a subset of the full SQL:2008 frame_clause grammar.
 * We don't support <window frame exclusion> yet.
 */
opt_frame_clause:
  RANGE frame_extent
  {
    $2.Mode = "range"
    $$ = $2
  }
| ROWS frame_extent
  {
    $2.Mode = "rows"
    $$ = $2
  }
| /*EMPTY*/
  {
    $$ = nil
  }

frame_extent:
  frame_bound
  {
    $$ = &FrameClause{Start: $1}
  }
| BETWEEN frame_bound AND frame_bound
  {
    $$ = &FrameClause{Start: $2, End: $4}
  }

/*
 * This is used for both frame start and frame end, with output set up on
 * the assumption it's frame start; the frame_extent productions must reject
 * invalid cases.
 */
frame_bound:
  UNBOUNDED PRECEDING
  {
    $$ = &FrameBound{Direction: "preceding"}
  }
| UNBOUNDED FOLLOWING
  {
    $$ = &FrameBound{Direction: "following"}
  }
| CURRENT_P ROW
  {
    $$ = &FrameBound{CurrentRow: true}
  }
| a_expr PRECEDING
  {
    $$ = &FrameBound{BoundExpr: $1, Direction: "preceding"}
  }
| a_expr FOLLOWING
  {
    $$ = &FrameBound{BoundExpr: $1, Direction: "following"}
  }



relation_expr:
  qualified_name
  {
    $$ = &RelationExpr{Name: $1}
  }
| qualified_name '*'
  {
    $$ = &RelationExpr{Name: $1, Star: true}
  }
| ONLY qualified_name
  {
    $$ = &RelationExpr{Name: $2, Only: true}
  }
| ONLY '(' qualified_name ')'
  {
    $$ = &RelationExpr{Name: $3, Only: true}
  }



select_limit:
  limit_clause offset_clause
  {
    $$ = &LimitClause{Limit: $1, Offset: $2}
  }
| offset_clause limit_clause
  {
    $$ = &LimitClause{Limit: $2, Offset: $1}
  }
| limit_clause
  {
    $$ = &LimitClause{Limit: $1}
  }
| offset_clause
  {
    $$ = &LimitClause{Offset: $1}
  }


opt_select_limit:
  select_limit
| /* EMPTY */   { $$ = nil }

limit_clause:
  LIMIT select_limit_value
  {
    $$ = $2
  }
  /* SQL:2008 syntax */
| FETCH first_or_next opt_select_fetch_first_value row_or_rows ONLY
  {
    $$ = $3
  }

offset_clause:
  OFFSET select_offset_value
  {
    $$ = $2
  }
  /* SQL:2008 syntax */
| OFFSET select_offset_value2 row_or_rows
  {
    $$ = $2
  }

select_limit_value:
  a_expr      { $$ = $1; }
| ALL
  {
    $$ = nil
  }

select_offset_value:
  a_expr   { $$ = $1 }

/*
 * Allowing full expressions without parentheses causes various parsing
 * problems with the trailing ROW/ROWS key words.  SQL only calls for
 * constants, so we allow the rest only with parentheses.  If omitted,
 * default to 1.
 */
opt_select_fetch_first_value:
  SignedIconst       { $$ = $1 }
| '(' a_expr ')'     { $$ = $2 }
| /*EMPTY*/          { $$ = IntegerLiteral("1") }

/*
 * Again, the trailing ROW/ROWS in this case prevent the full expression
 * syntax.  c_expr is the best we can do.
 */
select_offset_value2:
  c_expr  { $$ = $1 }

/* noise words */
row_or_rows:
  ROW   { $$ = 0 }
| ROWS  { $$ = 0 }

first_or_next:
  FIRST_P  { $$ = 0 }
| NEXT     { $$ = 0 }

values_clause:
VALUES ctext_row
  {
    $$ = ValuesClause{$2}
  }
| values_clause ',' ctext_row
  {
    $$ = append($1, $3)
  }


/*****************************************************************************
 *
 *  clauses common to all Optimizable Stmts:
 *    from_clause   - allow list of both JOIN expressions and table names
 *    where_clause  - qualifications for joins or restrictions
 *
 *****************************************************************************/




where_clause:
  WHERE a_expr   { $$ = &WhereClause{Expr: $2} }
| /*EMPTY*/      { $$ = nil }




/*
 * Aggregate decoration clauses
 */
within_group_clause:
      WITHIN GROUP_P '(' sort_clause ')'    { panic("TODO") }
      | /*EMPTY*/               { $$ = nil }
    ;

filter_clause:
  FILTER '(' WHERE a_expr ')'
  {
    $$ = &FilterClause{Expr: $4}
  }
| /*EMPTY*/ { $$ = nil }


all_Op:
  Op { $$ = string($1) }
| MathOp { $$ = string($1) }

MathOp:
  '+'             { $$ = "+" }
| '-'             { $$ = "-" }
| '*'             { $$ = "*" }
| '/'             { $$ = "/" }
| '%'             { $$ = "%" }
| '^'             { $$ = "^" }
| '<'             { $$ = "<" }
| '>'             { $$ = ">" }
| '='             { $$ = "=" }
| LESS_EQUALS     { $$ = "<=" }
| GREATER_EQUALS  { $$ = ">=" }
| NOT_EQUALS      { $$ = "<>" }

qual_Op:
  Op { $$ = string($1) }
| OPERATOR '(' any_operator ')' { $$ = $3 }

qual_all_Op:
  all_Op { $$ = string($1) }
| OPERATOR '(' any_operator ')' { $$ = $3 }


/*
 * Define SQL-style CASE clause.
 * - Full specification
 *  CASE WHEN a = b THEN c ... ELSE d END
 * - Implicit argument
 *  CASE a WHEN b THEN c ... ELSE d END
 */
case_expr:
  CASE case_arg when_clause_list case_default END_P
  {
    $$ = CaseExpr{CaseArg: $2, WhenClauses: $3, Default: $4}
  }

when_clause_list:
  /* There must be at least one */
  when_clause
  {
    $$ = []WhenClause{$1}
  }
| when_clause_list when_clause
  {
    $$ = append($1, $2)
  }

when_clause:
  WHEN a_expr THEN a_expr
  {
    $$ = WhenClause{When: $2, Then: $4}
  }

case_default:
  ELSE a_expr   { $$ = $2 }
| /*EMPTY*/     { $$ = nil }

case_arg:
  a_expr        { $$ = $1 }
| /*EMPTY*/     { $$ = nil }

columnref:
ColId
  {
    $$ = ColumnRef{Name: $1}
  }
| ColId indirection
  {
    $$ = ColumnRef{Name: $1, Indirection: $2}
  }



indirection_el:
  '.' attr_name
  {
    $$ = IndirectionEl{Name: $2}
  }
| '.' '*'
  {
    $$ = IndirectionEl{Name: "*"}
  }
| '[' a_expr ']'
  {
    $$ = IndirectionEl{LowerSubscript: $2}
  }
| '[' a_expr ':' a_expr ']'
  {
    $$ = IndirectionEl{LowerSubscript: $2, UpperSubscript:$4}
  }

indirection:
  indirection_el              { $$ = Indirection{$1} }
| indirection indirection_el  { $$ = append($1, $2) }

opt_indirection:
  /*EMPTY*/               { $$ = nil }
| opt_indirection indirection_el
  {
    if $1 != nil {
      $$ = append($1, $2)
    } else {
      $$ = Indirection{$2}
    }
  }

opt_asymmetric:
  ASYMMETRIC
  {
    $$ = nil
  }
| /*EMPTY*/
  {
    $$ = nil
  }

/*
 * The SQL spec defines "contextually typed value expressions" and
 * "contextually typed row value constructors", which for our purposes
 * are the same as "a_expr" and "row" except that DEFAULT can appear at
 * the top level.
 */

ctext_expr:
  a_expr    { $$ = $1 }
| DEFAULT   { $$ = DefaultExpr(true) }

ctext_expr_list:
  ctext_expr
  {
    $$ = ValuesRow{$1}
  }
| ctext_expr_list ',' ctext_expr
  {
    $$ = append($1, $3)
  }

/*
 * We should allow ROW '(' ctext_expr_list ')' too, but that seems to require
 * making VALUES a fully reserved word, which will probably break more apps
 * than allowing the noise-word is worth.
 */
ctext_row:
  '(' ctext_expr_list ')'
  {
    $$ = $2
  }

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
| '*'
  {
    $$ = ColumnRef{Name: "*"}
  }



/*****************************************************************************
 *
 *  Names and constants
 *
 *****************************************************************************/

qualified_name_list:
  qualified_name
  {
    $$ = []string{$1}
  }
| qualified_name_list ',' qualified_name
  {
    $$ = append($1, $3)
  }

/*
 * The production for a qualified relation name has to exactly match the
 * production for a qualified func_name, because in a FROM clause we cannot
 * tell which we are parsing until we see what comes after it ('(' for a
 * func_name, something else for a relation). Therefore we allow 'indirection'
 * which may contain subscripts, and reject that case in the C code.
 */
qualified_name:
  ColId
    {
      $$ = $1
    }
/* TODO
| ColId indirection
  {
    check_qualified_name($2, yyscanner);
    $$ = makeRangeVar(NULL, NULL, @1);
    switch (list_length($2))
    {
      case 1:
        $$->catalogname = NULL;
        $$->schemaname = $1;
        $$->relname = strVal(linitial($2));
        break;
      case 2:
        $$->catalogname = $1;
        $$->schemaname = strVal(linitial($2));
        $$->relname = strVal(lsecond($2));
        break;
      default:
        ereport(ERROR,
            (errcode(ERRCODE_SYNTAX_ERROR),
             errmsg("improper qualified name (too many dotted names): %s",
                NameListToString(lcons(makeString($1), $2))),
             parser_errposition(@1)));
        break;
    }
  }
*/

name_list:
  name { $$ = []string{$1} }
| name_list ',' name
  {
    $$ = append($1, $3)
  }

name:
  ColId { $$ = $1 }

attr_name:
  ColLabel { $$ = $1 }


/*
 * The production for a qualified func_name has to exactly match the
 * production for a qualified columnref, because we cannot tell which we
 * are parsing until we see what comes after it ('(' or Sconst for a func_name,
 * anything else for a columnref).  Therefore we allow 'indirection' which
 * may contain subscripts, and reject that case in the C code.  (If we
 * ever implement SQL99-like methods, such syntax may actually become legal!)
 */
func_name:
  type_function_name
  {
    $$ = $1
  }
| ColId indirection
  {
    panic("TODO")
  }


/*
 * Constants
 */
AexprConst:
Iconst
  {
    $$ = $1
  }
| FCONST
  {
    $$ = FloatConst($1)
  }
| Sconst
  {
    $$ = $1
  }
/* TODO
| BCONST
| XCONST
| func_name Sconst
| func_name '(' func_arg_list opt_sort_clause ')' Sconst
| ConstTypename Sconst
| ConstInterval Sconst opt_interval
| ConstInterval '(' Iconst ')' Sconst
*/
| TRUE_P
  {
    $$ = BoolLiteral(true)
  }
| FALSE_P
  {
    $$ = BoolLiteral(false)
  }
| NULL_P
  {
    $$ = NullLiteral{}
  }

Iconst:   ICONST { $$ = IntegerLiteral($1) }
Sconst:   SCONST { $$ = StringLiteral($1) }

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

