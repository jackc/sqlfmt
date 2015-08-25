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
  anyNames []AnyName
  intoClause *IntoClause
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
  row Row
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
  iconst IntegerConst
  optArrayBounds []IntegerConst
  optInterval *OptInterval
  intervalSecond *IntervalSecond
  subqueryOp SubqueryOp
  extractList *ExtractList
  overlayList OverlayList
  positionList *PositionList
  substrList SubstrList
  trimList TrimList
  xmlAttributes XmlAttributes
  xmlAttributeEls []XmlAttributeEl
  xmlAttributeEl XmlAttributeEl
  xmlExistsArgument XmlExistsArgument
  xmlRootVersion XmlRootVersion
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
%type <intoClause> OptTempTableName into_clause
%type <boolean> opt_table
%type <whereClause> where_clause
%type <orderExpr> sortby
%type <orderClause> opt_sort_clause sort_clause sortby_list
%type <str> opt_asc_desc opt_nulls_order opt_charset
%type <placeholder> row_or_rows
  first_or_next
  within_group_clause
  opt_asymmetric

%type <fields> opt_type_modifiers

%type <filterClause> filter_clause

%type <relationExpr> relation_expr

%type <extractList> extract_list
%type <expr> extract_arg

%type <overlayList> overlay_list
%type <expr> overlay_placing substr_from substr_for

%type <positionList> position_list

%type <placeholder> substr_list

%type <trimList> trim_list

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

%type <identifiers> name_list
%type <anyNames> qualified_name_list locked_rels_list

%type <indirectionEl> indirection_el
%type <indirection> indirection opt_indirection
%type <str> attr_name ColId name param_name
%type <anyName> qualified_name

%type <str> MathOp all_Op sub_type
%type <anyName> qual_Op qual_all_Op

%type <groupByClause> group_clause
%type <fields>  group_by_list
%type <expr> group_by_item
%type <expr> in_expr

%type <expr> having_clause

%type <boolean> all_or_distinct opt_varying opt_timezone

%type <expr>  SignedIconst Sconst AexprConst
%type <iconst> Iconst opt_float

%type <expr> case_expr case_arg case_default
%type <whenClauses> when_clause_list
%type <whenClause> when_clause

%type <expr> ctext_expr
%type <valuesRow> ctext_expr_list ctext_row
%type <row> row explicit_row implicit_row

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

%type <optInterval> opt_interval
%type <intervalSecond> interval_second

%type <arrayExpr> array_expr array_expr_list

%type <columnRef> columnref
%type <anyName> any_name attrs any_operator

%type <subqueryOp> subquery_Op

%type <xmlAttributes> xml_attributes
%type <xmlAttributeEls> xml_attribute_list
%type <xmlAttributeEl> xml_attribute_el
%type <xmlExistsArgument> xmlexists_argument
%type <str> document_or_content xml_whitespace_option
%type <xmlRootVersion> xml_root_version
%type <str> opt_xml_root_standalone


%type <str>
  ColLabel
  unreserved_keyword
  col_name_keyword
  type_function_name
  type_func_name_keyword
  reserved_keyword

%type <anyName> func_name

%type <pgType>
  GenericType
  Numeric
  Typename
  SimpleTypename
  Character
  CharacterWithLength
  CharacterWithoutLength
  character
  BitWithLength
  BitWithoutLength
  Bit
  ConstTypename
  ConstBit
  ConstCharacter
  ConstDatetime

%type <pgTypes> type_list
%type <optArrayBounds> opt_array_bounds

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
| SelectStmt ';'
  {
    $$ = $1
    $$.Semicolon = true
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

any_operator:
  all_Op
  {
    $$ = AnyName{$1}
  }
| ColId '.' any_operator
  {
    $$ = append(AnyName{$1}, $3...)
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
| SimpleTypename ARRAY '[' Iconst ']'
  {
    $$ = $1
    $$.ArrayWord = true
    $$.ArrayBounds = []IntegerConst{$4}
  }
| SETOF SimpleTypename ARRAY '[' Iconst ']'
  {
    $$ = $2
    $$.Setof = true
    $$.ArrayWord = true
    $$.ArrayBounds = []IntegerConst{$5}
  }
| SimpleTypename ARRAY
  {
    $$ = $1
    $$.ArrayWord = true
  }
| SETOF SimpleTypename ARRAY
  {
    $$ = $2
    $$.Setof = true
    $$.ArrayWord = true
  }

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
  GenericType
| Numeric
| Bit
| Character
| ConstDatetime
| ConstInterval opt_interval
  {
    $$ = PgType{Name: AnyName{"interval"}, OptInterval: $2}
  }
| ConstInterval '(' Iconst ')'
  {
    $$ = PgType{Name: AnyName{"interval"}, TypeMods: []Expr{$3}}
  }


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
ConstTypename:
  Numeric
  | ConstBit
  | ConstCharacter
  | ConstDatetime

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
    $$ = PgType{Name: AnyName{$1}, TypeMods: $2}
  }
| type_function_name attrs opt_type_modifiers
  {
    $$ = PgType{Name: append(AnyName{$1}, $2...), TypeMods: $3}
  }

opt_type_modifiers:
  '(' expr_list ')'   { $$ = $2 }
| /* EMPTY */         { $$ = nil }

/*
 * SQL numeric data types
 */
Numeric:
  INT_P
  {
    $$ = PgType{Name: AnyName{"int"}}
  }
| INTEGER
  {
    $$ = PgType{Name: AnyName{"integer"}}
  }
| SMALLINT
  {
    $$ = PgType{Name: AnyName{"smallint"}}
  }
| BIGINT
  {
    $$ = PgType{Name: AnyName{"bigint"}}
  }
| REAL
  {
    $$ = PgType{Name: AnyName{"real"}}
  }
| FLOAT_P opt_float
  {
    $$ = PgType{Name: AnyName{"float"}}
    if $2 != IntegerConst("") {
      $$.TypeMods = []Expr{$2}
    }
  }
| DOUBLE_P PRECISION
  {
    $$ = PgType{Name: AnyName{"double precision"}}
  }
| DECIMAL_P opt_type_modifiers
  {
    $$ = PgType{Name: AnyName{"decimal"}, TypeMods: $2}
  }
| DEC opt_type_modifiers
  {
    $$ = PgType{Name: AnyName{"dec"}, TypeMods: $2}
  }
| NUMERIC opt_type_modifiers
  {
    $$ = PgType{Name: AnyName{"numeric"}, TypeMods: $2}
  }
| BOOLEAN_P
  {
    $$ = PgType{Name: AnyName{"bool"}}
  }

opt_float:
  '(' Iconst ')'
  {
    $$ = $2
  }
| /*EMPTY*/
  {
    $$ = IntegerConst("")
  }

Bit:
  BitWithLength
| BitWithoutLength

ConstBit:
  BitWithLength
| BitWithoutLength

BitWithLength:
  BIT opt_varying '(' expr_list ')'
  {
    $$ = PgType{}
    if $2 {
      $$.Name = AnyName{"varbit"}
    } else {
      $$.Name = AnyName{"bit"}
    }
    $$.TypeMods = $4
  }

BitWithoutLength:
  BIT opt_varying
  {
    $$ = PgType{}
    if $2 {
      $$ = PgType{Name: AnyName{"varbit"}}
    } else {
      $$ = PgType{Name: AnyName{"bit"}}
    }
  }

/*
 * SQL character data types
 * The following implements CHAR() and VARCHAR().
 */
Character:
  CharacterWithLength
| CharacterWithoutLength

ConstCharacter:
  CharacterWithLength
| CharacterWithoutLength

CharacterWithLength:
  character '(' Iconst ')' opt_charset
  {
    $$ = $1
    $$.TypeMods = []Expr{$3}
    $$.CharSet = $5
  }

CharacterWithoutLength:
  character opt_charset
  {
    $$ = $1
    $$.CharSet = $2
  }

character:
  CHARACTER opt_varying
  {
    if $2 {
      $$ = PgType{Name: AnyName{"varchar"}}
    } else {
      $$ = PgType{Name: AnyName{"char"}}
    }
  }
| CHAR_P opt_varying
  {
    if $2 {
      $$ = PgType{Name: AnyName{"varchar"}}
    } else {
      $$ = PgType{Name: AnyName{"char"}}
    }
  }
| VARCHAR
  {
    $$ = PgType{Name: AnyName{"varchar"}}
  }
| NATIONAL CHARACTER opt_varying
  {
    if $3 {
      $$ = PgType{Name: AnyName{"varchar"}}
    } else {
      $$ = PgType{Name: AnyName{"char"}}
    }
  }
| NATIONAL CHAR_P opt_varying
  {
    if $3 {
      $$ = PgType{Name: AnyName{"varchar"}}
    } else {
      $$ = PgType{Name: AnyName{"char"}}
    }
  }
| NCHAR opt_varying
  {
    if $2 {
      $$ = PgType{Name: AnyName{"varchar"}}
    } else {
      $$ = PgType{Name: AnyName{"char"}}
    }
  }

opt_varying:
  VARYING
  {
    $$ = true
  }
| /*EMPTY*/
  {
    $$ = false
  }

opt_charset:
  CHARACTER SET ColId
  {
    $$ = $3
  }
| /*EMPTY*/
  {
    $$ = ""
  }

/*
 * SQL date/time types
 */
ConstDatetime:
  TIMESTAMP '(' Iconst ')' opt_timezone
  {
    $$ = PgType{Name: AnyName{"timestamp"}, TypeMods: []Expr{$3}, WithTimeZone: $5}
  }
| TIMESTAMP opt_timezone
  {
    $$ = PgType{Name: AnyName{"timestamp"}, WithTimeZone: $2}
  }
| TIME '(' Iconst ')' opt_timezone
  {
    $$ = PgType{Name: AnyName{"time"}, TypeMods: []Expr{$3}, WithTimeZone: $5}
  }
| TIME opt_timezone
  {
    $$ = PgType{Name: AnyName{"time"}, WithTimeZone: $2}
  }

ConstInterval:
  INTERVAL

opt_timezone:
  WITH_LA TIME ZONE
  {
    $$ = true
  }
| WITHOUT TIME ZONE
  {
    $$ = false
  }
| /*EMPTY*/
  {
    $$ = false
  }

opt_interval:
  YEAR_P
  {
    $$ = &OptInterval{Left: "year"}
  }
| MONTH_P
  {
    $$ = &OptInterval{Left: "month"}
  }
| DAY_P
  {
    $$ = &OptInterval{Left: "day"}
  }
| HOUR_P
  {
    $$ = &OptInterval{Left: "hour"}
  }
| MINUTE_P
  {
    $$ = &OptInterval{Left: "minute"}
  }
| interval_second
  {
    $$ = &OptInterval{Second: $1}
  }
| YEAR_P TO MONTH_P
  {
    $$ = &OptInterval{Left: "year", Right: "month"}
  }
| DAY_P TO HOUR_P
  {
    $$ = &OptInterval{Left: "day", Right: "hour"}
  }
| DAY_P TO MINUTE_P
  {
    $$ = &OptInterval{Left: "day", Right: "minute"}
  }
| DAY_P TO interval_second
  {
    $$ = &OptInterval{Left: "day", Second: $3}
  }
| HOUR_P TO MINUTE_P
  {
    $$ = &OptInterval{Left: "hour", Right: "minute"}
  }
| HOUR_P TO interval_second
  {
    $$ = &OptInterval{Left: "hour", Second: $3}
  }
| MINUTE_P TO interval_second
  {
    $$ = &OptInterval{Left: "minute", Second: $3}
  }
| /*EMPTY*/
  {
    $$ = nil
  }

interval_second:
  SECOND_P
  {
    $$ = &IntervalSecond{}
  }
| SECOND_P '(' Iconst ')'
  {
    $$ = &IntervalSecond{Precision: $3}
  }


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
    $$ = UnaryExpr{Operator: AnyName{"+"}, Expr: $2}
  }
| '-' a_expr          %prec UMINUS
  {
    $$ = UnaryExpr{Operator: AnyName{"-"}, Expr: $2}
  }
| a_expr '+' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"+"}, Right: $3}
  }
| a_expr '-' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"-"}, Right: $3}
  }
| a_expr '*' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"*"}, Right: $3}
  }
| a_expr '/' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"/"}, Right: $3}
  }
| a_expr '%' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"%"}, Right: $3}
  }
| a_expr '^' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"^"}, Right: $3}
  }
| a_expr '<' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"<"}, Right: $3}
  }
| a_expr '>' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{">"}, Right: $3}
  }
| a_expr '=' a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"="}, Right: $3}
  }
| a_expr LESS_EQUALS a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"<="}, Right: $3}
  }
| a_expr GREATER_EQUALS a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{">="}, Right: $3}
  }
| a_expr NOT_EQUALS a_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"!="}, Right: $3}
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
    $$ = IsExpr{Expr: $1, Op: "null"}
  }
| a_expr ISNULL
  {
    $$ = IsExpr{Expr: $1, Op: "null"}
  }
| a_expr IS NOT NULL_P            %prec IS
  {
    $$ = IsExpr{Expr: $1, Not: true, Op: "null"}
  }
| a_expr NOTNULL
  {
    $$ = IsExpr{Expr: $1, Not: true, Op: "null"}
  }
| row OVERLAPS row
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"overlaps"}, Right: $3}
  }
| a_expr IS TRUE_P              %prec IS
  {
    $$ = IsExpr{Expr: $1, Op: "true"}
  }
| a_expr IS NOT TRUE_P            %prec IS
  {
    $$ = IsExpr{Expr: $1, Not: true, Op: "true"}
  }
| a_expr IS FALSE_P             %prec IS
  {
    $$ = IsExpr{Expr: $1, Op: "false"}
  }
| a_expr IS NOT FALSE_P           %prec IS
  {
    $$ = IsExpr{Expr: $1, Not: true, Op: "false"}
  }
| a_expr IS UNKNOWN             %prec IS
  {
    $$ = IsExpr{Expr: $1, Op: "unknown"}
  }
| a_expr IS NOT UNKNOWN           %prec IS
  {
    $$ = IsExpr{Expr: $1, Not: true, Op: "unknown"}
  }
| a_expr IS DISTINCT FROM a_expr      %prec IS
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"is distinct from"}, Right: $5}
  }
| a_expr IS NOT DISTINCT FROM a_expr    %prec IS
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"is not distinct from"}, Right: $6}
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
| a_expr IN_P in_expr
  {
    $$ = InExpr{Value: $1, In: $3}
  }
| a_expr NOT_LA IN_P in_expr            %prec NOT_LA
  {
    $$ = InExpr{Value: $1, Not: true, In: $4}
  }
| a_expr subquery_Op sub_type select_with_parens  %prec Op
  {
    $$ = SubqueryOpExpr{Value: $1, Op: $2, Type: $3, Query: $4}
  }
| a_expr subquery_Op sub_type '(' a_expr ')'    %prec Op
  {
    $$ = SubqueryOpExpr{Value: $1, Op: $2, Type: $3, Query: ParenExpr{Expr: $5}}
  }
| a_expr IS DOCUMENT_P          %prec IS
  {
    $$ = IsExpr{Expr: $1, Op: "document"}
  }
| a_expr IS NOT DOCUMENT_P        %prec IS
  {
    $$ = IsExpr{Expr: $1, Not: true, Op: "document"}
  }

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
    $$ = UnaryExpr{Operator: AnyName{"+"}, Expr: $2}
  }
| '-' b_expr          %prec UMINUS
  {
    $$ = UnaryExpr{Operator: AnyName{"-"}, Expr: $2}
  }
| b_expr '+' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"+"}, Right: $3}
  }
| b_expr '-' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"-"}, Right: $3}
  }
| b_expr '*' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"*"}, Right: $3}
  }
| b_expr '/' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"/"}, Right: $3}
  }
| b_expr '%' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"%"}, Right: $3}
  }
| b_expr '^' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"^"}, Right: $3}
  }
| b_expr '<' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"<"}, Right: $3}
  }
| b_expr '>' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{">"}, Right: $3}
  }
| b_expr '=' b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"="}, Right: $3}
  }
| b_expr LESS_EQUALS b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"<="}, Right: $3}
  }
| b_expr GREATER_EQUALS b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{">="}, Right: $3}
  }
| b_expr NOT_EQUALS b_expr
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"!="}, Right: $3}
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
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"is distinct from"}, Right: $5}
  }
| b_expr IS NOT DISTINCT FROM b_expr  %prec IS
  {
    $$ = BinaryExpr{Left: $1, Operator: AnyName{"is not distinct from"}, Right: $6}
  }
| b_expr IS OF '(' type_list ')'    %prec IS
  {
    $$ = IsOfExpr{Expr: $1, Types: $5}
  }
| b_expr IS NOT OF '(' type_list ')'  %prec IS
  {
    $$ = IsOfExpr{Expr: $1, Not: true, Types: $6}
  }
| b_expr IS DOCUMENT_P          %prec IS
  {
    $$ = IsExpr{Expr: $1, Op: "document"}
  }
| b_expr IS NOT DOCUMENT_P        %prec IS
  {
    $$ = IsExpr{Expr: $1, Not: true, Op: "document"}
  }

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
/*
  PARAM does not occur directly is the grammar -- it is used for output
  params of functions. See outfuncs.c in PostgreSQL source.
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
| explicit_row
  {
    $$ = $1
  }
| implicit_row
  {
    $$ = $1
  }
/* TODO
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
  COLLATION FOR '(' a_expr ')'
  {
    $$ = FuncApplication{Name: AnyName{"collation for"}, Args: []FuncArg{{Expr: $4}}}
  }
| CURRENT_DATE
  {
    $$ = FuncExprNoParens("current_date")
  }
| CURRENT_TIME
  {
    $$ = FuncExprNoParens("current_time")
  }
| CURRENT_TIME '(' Iconst ')'
  {
    $$ = FuncApplication{Name: AnyName{"current_time"}, Args: []FuncArg{{Expr: $3}}}
  }
| CURRENT_TIMESTAMP
  {
    $$ = FuncExprNoParens("current_timestamp")
  }
| CURRENT_TIMESTAMP '(' Iconst ')'
  {
    $$ = FuncApplication{Name: AnyName{"current_timestamp"}, Args: []FuncArg{{Expr: $3}}}
  }
| LOCALTIME
  {
    $$ = FuncExprNoParens("localtime")
  }
| LOCALTIME '(' Iconst ')'
  {
    $$ = FuncApplication{Name: AnyName{"localtime"}, Args: []FuncArg{{Expr: $3}}}
  }
| LOCALTIMESTAMP
  {
    $$ = FuncExprNoParens("localtimestamp")
  }
| LOCALTIMESTAMP '(' Iconst ')'
  {
    $$ = FuncApplication{Name: AnyName{"localtimestamp"}, Args: []FuncArg{{Expr: $3}}}
  }
| CURRENT_ROLE
  {
    $$ = FuncExprNoParens("current_role")
  }
| CURRENT_USER
  {
    $$ = FuncExprNoParens("current_user")
  }
| SESSION_USER
  {
    $$ = FuncExprNoParens("session_user")
  }
| USER
  {
    $$ = FuncExprNoParens("user")
  }
| CURRENT_CATALOG
  {
    $$ = FuncExprNoParens("current_catalog")
  }
| CURRENT_SCHEMA
  {
    $$ = FuncExprNoParens("current_schema")
  }
| CAST '(' a_expr AS Typename ')'
  {
    $$ = CastFunc{Name: "cast", Expr: $3, Type: $5}
  }
| EXTRACT '(' extract_list ')'
  {
    $$ = ExtractExpr(*$3)
  }
| OVERLAY '(' overlay_list ')'
  {
    $$ = OverlayExpr($3)
  }
| POSITION '(' position_list ')'
  {
    $$ = PositionExpr(*$3)
  }
| SUBSTRING '(' substr_list ')'
  {
    if $3 == nil {
      $$ = FuncApplication{Name: AnyName{"substring"}}
    } else if se, ok := $3.(SubstrList); ok {
      $$ = SubstrExpr(se)
    } else {
      var args []FuncArg
      for _, a := range $3.([]Expr) {
        args = append(args, FuncArg{Expr: a})
      }
      $$ = FuncApplication{Name: AnyName{"substring"}, Args: args}
    }
  }
| TREAT '(' a_expr AS Typename ')'
  {
    $$ = CastFunc{Name: "treat", Expr: $3, Type: $5}
  }
| TRIM '(' BOTH trim_list ')'
  {
    $$ = TrimExpr{Direction: "both", TrimList: $4}
  }
| TRIM '(' LEADING trim_list ')'
  {
    $$ = TrimExpr{Direction: "leading", TrimList: $4}
  }
| TRIM '(' TRAILING trim_list ')'
  {
    $$ = TrimExpr{Direction: "trailing", TrimList: $4}
  }
| TRIM '(' trim_list ')'
  {
    $$ = TrimExpr{TrimList: $3}
  }
| NULLIF '(' a_expr ',' a_expr ')'
  {
    $$ = FuncApplication{Name: AnyName{"nullif"}, Args: []FuncArg{{Expr: $3}, {Expr: $5}}}
  }
| COALESCE '(' expr_list ')'
{
  fa := FuncApplication{Name: AnyName{"coalesce"}}
  for _, e := range $3 {
    fa.Args = append(fa.Args, FuncArg{Expr: e})
  }
  $$ = fa
}
| GREATEST '(' expr_list ')'
{
  fa := FuncApplication{Name: AnyName{"greatest"}}
  for _, e := range $3 {
    fa.Args = append(fa.Args, FuncArg{Expr: e})
  }
  $$ = fa
}
| LEAST '(' expr_list ')'
{
  fa := FuncApplication{Name: AnyName{"least"}}
  for _, e := range $3 {
    fa.Args = append(fa.Args, FuncArg{Expr: e})
  }
  $$ = fa
}
| XMLCONCAT '(' expr_list ')'
{
  fa := FuncApplication{Name: AnyName{"xmlconcat"}}
  for _, e := range $3 {
    fa.Args = append(fa.Args, FuncArg{Expr: e})
  }
  $$ = fa
}
| XMLELEMENT '(' NAME_P ColLabel ')'
  {
    $$ = XmlElement{Name: $4}
  }
| XMLELEMENT '(' NAME_P ColLabel ',' xml_attributes ')'
  {
    $$ = XmlElement{Name: $4, Attributes: $6}
  }
| XMLELEMENT '(' NAME_P ColLabel ',' expr_list ')'
  {
    $$ = XmlElement{Name: $4, Body: $6}
  }
| XMLELEMENT '(' NAME_P ColLabel ',' xml_attributes ',' expr_list ')'
  {
    $$ = XmlElement{Name: $4, Attributes: $6, Body: $8}
  }
| XMLEXISTS '(' c_expr xmlexists_argument ')'
  {
    $$ = XmlExists{Path: $3, Body: $4}
  }
| XMLFOREST '(' xml_attribute_list ')'
  {
    $$ = XmlForest($3)
  }
| XMLPARSE '(' document_or_content a_expr xml_whitespace_option ')'
  {
    $$ = XmlParse{Type: $3, Content: $4, WhitespaceOption: $5}
  }
| XMLPI '(' NAME_P ColLabel ')'
  {
    $$ = XmlPi{Name: $4}
  }
| XMLPI '(' NAME_P ColLabel ',' a_expr ')'
  {
    $$ = XmlPi{Name: $4, Content: $6}
  }
| XMLROOT '(' a_expr ',' xml_root_version opt_xml_root_standalone ')'
  {
    $$ = XmlRoot{Xml: $3, Version: $5, Standalone: $6}
  }
| XMLSERIALIZE '(' document_or_content a_expr AS SimpleTypename ')'
  {
    $$ = XmlSerialize{XmlType: $3, Content: $4, Type: $6}
  }

/*
 * SQL/XML support
 */
xml_root_version:
  VERSION_P a_expr
  {
    $$ = XmlRootVersion{Expr: $2}
  }
| VERSION_P NO VALUE_P
  {
    $$ = XmlRootVersion{}
  }

opt_xml_root_standalone:
  ',' STANDALONE_P YES_P      { $$ = "yes" }
| ',' STANDALONE_P NO         { $$ = "no" }
| ',' STANDALONE_P NO VALUE_P { $$ = "no value"}
| /*EMPTY*/                   { $$ = "" }

xml_attributes:
  XMLATTRIBUTES '(' xml_attribute_list ')'
  {
    $$ = XmlAttributes($3)
  }

xml_attribute_list:
  xml_attribute_el
  {
    $$ = []XmlAttributeEl{$1}
  }
| xml_attribute_list ',' xml_attribute_el
  {
    $$ = append($1, $3)
  }

xml_attribute_el:
  a_expr AS ColLabel
  {
    $$ = XmlAttributeEl{Value: $1, Name: $3}
  }
| a_expr
  {
    $$ = XmlAttributeEl{Value: $1}
  }

document_or_content:
  DOCUMENT_P { $$ = "document" }
| CONTENT_P { $$ = "content" }

xml_whitespace_option:
  PRESERVE WHITESPACE_P { $$ = "preserve whitespace" }
| STRIP_P WHITESPACE_P  { $$ = "strip whitespace" }
| /*EMPTY*/             { $$ = "" }

/* We allow several variants for SQL and other compatibility. */
xmlexists_argument:
  PASSING c_expr
  {
    $$ = XmlExistsArgument{Arg: $2}
  }
| PASSING c_expr BY REF
  {
    $$ = XmlExistsArgument{Arg: $2, RightByRef: true}
  }
| PASSING BY REF c_expr
  {
    $$ = XmlExistsArgument{LeftByRef: true, Arg: $4}
  }
| PASSING BY REF c_expr BY REF
  {
    $$ = XmlExistsArgument{LeftByRef: true, Arg: $4, RightByRef: true}
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

subquery_Op:
  all_Op
  {
    $$ = SubqueryOp{Name: AnyName{$1}}
  }
| OPERATOR '(' any_operator ')'
  {
    $$ = SubqueryOp{Operator: true, Name: $3}
  }
| LIKE
  {
    $$ = SubqueryOp{Name: AnyName{"like"}}
  }
| NOT_LA LIKE
  {
    $$ = SubqueryOp{Name: AnyName{"not like"}}
  }
| ILIKE
  {
    $$ = SubqueryOp{Name: AnyName{"ilike"}}
  }
| NOT_LA ILIKE
  {
    $$ = SubqueryOp{Name: AnyName{"not ilike"}}
  }
/* cannot put SIMILAR TO here, because SIMILAR TO is a hack.
 * the regular expression is preprocessed by a function (similar_escape),
 * and the ~ operator for posix regular expressions is used.
 *        x SIMILAR TO y     ->    x ~ similar_escape(y)
 * this transformation is made on the fly by the parser upwards.
 * however the SubLink structure which handles any/some/all stuff
 * is not ready for such a thing.
 */

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

extract_list:
  extract_arg FROM a_expr
  {
    $$ = &ExtractList{Extract: $1, Time: $3}
  }
| /*EMPTY*/
  {
    $$ = nil
  }

/* Allow delimited string Sconst in extract_arg as an SQL extension.
 * - thomas 2001-04-12
 */
extract_arg:
  IDENT     { $$ = AnyName{$1} }
| YEAR_P    { $$ = AnyName{"year"} }
| MONTH_P   { $$ = AnyName{"month"} }
| DAY_P     { $$ = AnyName{"day"} }
| HOUR_P    { $$ = AnyName{"hour"} }
| MINUTE_P  { $$ = AnyName{"minute"} }
| SECOND_P  { $$ = AnyName{"second"} }
| Sconst    { $$ = $1 }

/* OVERLAY() arguments
 * SQL99 defines the OVERLAY() function:
 * o overlay(text placing text from int for int)
 * o overlay(text placing text from int)
 * and similarly for binary strings
 */
overlay_list:
  a_expr overlay_placing substr_from substr_for
  {
    $$ = OverlayList{Dest: $1, Placing: $2, From: $3, For: $4}
  }
| a_expr overlay_placing substr_from
  {
    $$ = OverlayList{Dest: $1, Placing: $2, From: $3}
  }

overlay_placing:
  PLACING a_expr
  {
    $$ = $2
  }

/* position_list uses b_expr not a_expr to avoid conflict with general IN */

position_list:
  b_expr IN_P b_expr
  {
    $$ = &PositionList{Substring: $1, String: $3}
  }
| /*EMPTY*/  { $$ = nil }

/* SUBSTRING() arguments
 * SQL9x defines a specific syntax for arguments to SUBSTRING():
 * o substring(text from int for int)
 * o substring(text from int) get entire string from starting point "int"
 * o substring(text for int) get first "int" characters of string
 * o substring(text from pattern) get entire string matching pattern
 * o substring(text from pattern for escape) same with specified escape char
 * We also want to support generic substring functions which accept
 * the usual generic list of arguments. So we will accept both styles
 * here, and convert the SQL9x style to the generic list for further
 * processing. - thomas 2000-11-28
 */
substr_list:
a_expr substr_from substr_for
  {
    $$ = SubstrList{Source: $1, From: $2, For: $3}
  }
| a_expr substr_for substr_from
  {
    /* not legal per SQL99, but might as well allow it */
    $$ = SubstrList{Source: $1, From: $3, For: $2}
  }
| a_expr substr_from
  {
    $$ = SubstrList{Source: $1, From: $2}
  }
| a_expr substr_for
  {
    $$ = SubstrList{Source: $1, For: $2}
  }
| expr_list
  {
    $$ = $1
  }
| /*EMPTY*/
  {
    $$ = nil
  }

substr_from:
  FROM a_expr
  {
    $$ = $2
  }

substr_for:
  FOR a_expr
  {
    $$ = $2
  }

trim_list:
  a_expr FROM expr_list
  {
    $$ = TrimList{Left: $1, From: true, Right: $3}
  }
| FROM expr_list
  {
    $$ = TrimList{From: true, Right: $2}
  }
| expr_list
  {
    $$ = TrimList{Right: $1}
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
    ss.IntoClause = $4
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
    ss.IntoClause = $4
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
  INTO OptTempTableName
  {
    $$ = $2
  }
| /*EMPTY*/
  {
    $$ = nil
  }

/*
 * Redundancy here is needed to avoid shift/reduce conflicts,
 * since TEMP is not a reserved word.  See also OptTemp.
 */
OptTempTableName:
  TEMPORARY opt_table qualified_name
  {
    $$ = &IntoClause{Options: "temporary", OptTable: $2, Target: $3}
  }
| TEMP opt_table qualified_name
  {
    $$ = &IntoClause{Options: "temp", OptTable: $2, Target: $3}
  }
| LOCAL TEMPORARY opt_table qualified_name
  {
    $$ = &IntoClause{Options: "local temporary", OptTable: $3, Target: $4}
  }
| LOCAL TEMP opt_table qualified_name
  {
    $$ = &IntoClause{Options: "local temp", OptTable: $3, Target: $4}
  }
| GLOBAL TEMPORARY opt_table qualified_name
  {
    $$ = &IntoClause{Options: "global temporary", OptTable: $3, Target: $4}
  }
| GLOBAL TEMP opt_table qualified_name
  {
    $$ = &IntoClause{Options: "global temp", OptTable: $3, Target: $4}
  }
| UNLOGGED opt_table qualified_name
  {
    $$ = &IntoClause{Options: "unlogged", OptTable: $2, Target: $3}
  }
| TABLE qualified_name
  {
    $$ = &IntoClause{OptTable: true, Target: $2}
  }
| qualified_name
  {
    $$ = &IntoClause{Target: $1}
  }

opt_table:
  TABLE      { $$ = true }
| /*EMPTY*/  { $$ = false }

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
| /*EMPTY*/          { $$ = IntegerConst("1") }

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



/* Explicit row production.
 *
 * SQL99 allows an optional ROW keyword, so we can now do single-element rows
 * without conflicting with the parenthesized a_expr production.  Without the
 * ROW keyword, there must be more than one a_expr inside the parens.
 */
row:
  ROW '(' expr_list ')'
  {
    $$ = Row{RowWord: true, Exprs: $3}
  }
| ROW '(' ')'
  {
    $$ = Row{RowWord: true, Exprs: nil}
  }
| '(' expr_list ',' a_expr ')'
  {
    $$ = Row{Exprs: append($2, $4)}
  }

explicit_row:
  ROW '(' expr_list ')'
  {
    $$ = Row{RowWord: true, Exprs: $3}
  }
| ROW '(' ')'
  {
    $$ = Row{RowWord: true, Exprs: nil}
  }

implicit_row:
  '(' expr_list ',' a_expr ')'
  {
    $$ = Row{Exprs: append($2, $4)}
  }

sub_type:
  ANY  { $$ = "any" }
| SOME { $$ = "some" }
| ALL  { $$ = "all" }

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
  Op { $$ = AnyName{$1} }
| OPERATOR '(' any_operator ')' { $$ = $3 }

qual_all_Op:
  all_Op { $$ = AnyName{$1} }
| OPERATOR '(' any_operator ')' { $$ = $3 }

in_expr:
  select_with_parens
  {
    $$ = $1
  }
| '(' expr_list ')'
  {
    $$ = ValuesRow($2)
  }

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
    $$ = []AnyName{$1}
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
    $$ = AnyName{$1}
  }
| ColId indirection
  {
    $$ = AnyName{$1}
    for _, s := range $2 {
      $$ = append($$, s.Name)
    }
  }

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
    $$ = AnyName{$1}
  }
| ColId indirection
  {
    $$ = AnyName{$1}
    for _, s := range $2 {
      $$ = append($$, s.Name)
    }
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
| BCONST
  {
    $$ = BitConst($1)
  }
| XCONST
  {
    $$ = BitConst($1)
  }
  | func_name Sconst
  {
    $$ = ConstTypeExpr{Typename: PgType{Name: $1}, Expr: $2}
  }
| func_name '(' func_arg_list opt_sort_clause ')' Sconst
  {
    pgType := PgType{Name: $1}

    /*
     * We must use func_arg_list and opt_sort_clause in the
     * production to avoid reduce/reduce conflicts, but we
     * don't actually wish to allow NamedArgExpr in this
     * context, nor ORDER BY.
     */

    for _, arg := range $3 {
      if arg.Name != "" {
        yylex.Error("type modifier cannot have parameter name")
      }

      pgType.TypeMods = append(pgType.TypeMods, Expr(arg))
    }

    if $4 != nil {
      yylex.Error("type modifier cannot have ORDER BY")
    }

    $$ = ConstTypeExpr{Typename: pgType, Expr: $6}
  }
| ConstTypename Sconst
  {
    $$ = ConstTypeExpr{Typename: $1, Expr: $2}
  }
| ConstInterval Sconst opt_interval
  {
    $$ = ConstIntervalExpr{Value: $2, OptInterval: $3}
  }
| ConstInterval '(' Iconst ')' Sconst
  {
    $$ = ConstIntervalExpr{Precision: $3, Value: $5}
  }
| TRUE_P
  {
    $$ = BoolConst(true)
  }
| FALSE_P
  {
    $$ = BoolConst(false)
  }
| NULL_P
  {
    $$ = NullConst{}
  }

Iconst:   ICONST { $$ = IntegerConst($1) }
Sconst:   SCONST { $$ = StringConst($1) }

SignedIconst:
  Iconst      { $$ = $1 }
| '+' Iconst  { $$ = "+" + $2 }
| '-' Iconst  { $$ = "-" + $2 }

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

