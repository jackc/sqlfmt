//line sql.y:2
package main

import __yyfmt__ "fmt"

//line sql.y:3
//line sql.y:7
type sqlSymType struct {
	yys         int
	sqlSelect   *SelectStmt
	fields      []Expr
	expr        Expr
	src         string
	identifiers []string
	fromClause  *FromClause
	whereClause *WhereClause
	orderExpr   OrderExpr
	orderClause *OrderClause
}

const IDENT = 57346
const ICONST = 57347
const SCONST = 57348
const OP = 57349
const ABORT_P = 57350
const ABSOLUTE_P = 57351
const ACCESS = 57352
const ACTION = 57353
const ADD_P = 57354
const ADMIN = 57355
const AFTER = 57356
const AGGREGATE = 57357
const ALL = 57358
const ALSO = 57359
const ALTER = 57360
const ALWAYS = 57361
const ANALYSE = 57362
const ANALYZE = 57363
const AND = 57364
const ANY = 57365
const ARRAY = 57366
const AS = 57367
const ASC = 57368
const ASSERTION = 57369
const ASSIGNMENT = 57370
const ASYMMETRIC = 57371
const AT = 57372
const ATTRIBUTE = 57373
const AUTHORIZATION = 57374
const BACKWARD = 57375
const BEFORE = 57376
const BEGIN_P = 57377
const BETWEEN = 57378
const BIGINT = 57379
const BINARY = 57380
const BIT = 57381
const BOOLEAN_P = 57382
const BOTH = 57383
const BY = 57384
const CACHE = 57385
const CALLED = 57386
const CASCADE = 57387
const CASCADED = 57388
const CASE = 57389
const CAST = 57390
const CATALOG_P = 57391
const CHAIN = 57392
const CHAR_P = 57393
const CHARACTER = 57394
const CHARACTERISTICS = 57395
const CHECK = 57396
const CHECKPOINT = 57397
const CLASS = 57398
const CLOSE = 57399
const CLUSTER = 57400
const COALESCE = 57401
const COLLATE = 57402
const COLLATION = 57403
const COLUMN = 57404
const COMMENT = 57405
const COMMENTS = 57406
const COMMIT = 57407
const COMMITTED = 57408
const CONCURRENTLY = 57409
const CONFIGURATION = 57410
const CONFLICT = 57411
const CONNECTION = 57412
const CONSTRAINT = 57413
const CONSTRAINTS = 57414
const CONTENT_P = 57415
const CONTINUE_P = 57416
const CONVERSION_P = 57417
const COPY = 57418
const COST = 57419
const CREATE = 57420
const CROSS = 57421
const CSV = 57422
const CUBE = 57423
const CURRENT_P = 57424
const CURRENT_CATALOG = 57425
const CURRENT_DATE = 57426
const CURRENT_ROLE = 57427
const CURRENT_SCHEMA = 57428
const CURRENT_TIME = 57429
const CURRENT_TIMESTAMP = 57430
const CURRENT_USER = 57431
const CURSOR = 57432
const CYCLE = 57433
const DATA_P = 57434
const DATABASE = 57435
const DAY_P = 57436
const DEALLOCATE = 57437
const DEC = 57438
const DECIMAL_P = 57439
const DECLARE = 57440
const DEFAULT = 57441
const DEFAULTS = 57442
const DEFERRABLE = 57443
const DEFERRED = 57444
const DEFINER = 57445
const DELETE_P = 57446
const DELIMITER = 57447
const DELIMITERS = 57448
const DESC = 57449
const DICTIONARY = 57450
const DISABLE_P = 57451
const DISCARD = 57452
const DISTINCT = 57453
const DO = 57454
const DOCUMENT_P = 57455
const DOMAIN_P = 57456
const DOUBLE_P = 57457
const DROP = 57458
const EACH = 57459
const ELSE = 57460
const ENABLE_P = 57461
const ENCODING = 57462
const ENCRYPTED = 57463
const END_P = 57464
const ENUM_P = 57465
const ESCAPE = 57466
const EVENT = 57467
const EXCEPT = 57468
const EXCLUDE = 57469
const EXCLUDING = 57470
const EXCLUSIVE = 57471
const EXECUTE = 57472
const EXISTS = 57473
const EXPLAIN = 57474
const EXTENSION = 57475
const EXTERNAL = 57476
const EXTRACT = 57477
const FALSE_P = 57478
const FAMILY = 57479
const FETCH = 57480
const FILTER = 57481
const FIRST_P = 57482
const FLOAT_P = 57483
const FOLLOWING = 57484
const FOR = 57485
const FORCE = 57486
const FOREIGN = 57487
const FORWARD = 57488
const FREEZE = 57489
const FROM = 57490
const FULL = 57491
const FUNCTION = 57492
const FUNCTIONS = 57493
const GLOBAL = 57494
const GRANT = 57495
const GRANTED = 57496
const GREATEST = 57497
const GROUP_P = 57498
const GROUPING = 57499
const HANDLER = 57500
const HAVING = 57501
const HEADER_P = 57502
const HOLD = 57503
const HOUR_P = 57504
const IDENTITY_P = 57505
const IF_P = 57506
const ILIKE = 57507
const IMMEDIATE = 57508
const IMMUTABLE = 57509
const IMPLICIT_P = 57510
const IMPORT_P = 57511
const IN_P = 57512
const INCLUDING = 57513
const INCREMENT = 57514
const INDEX = 57515
const INDEXES = 57516
const INHERIT = 57517
const INHERITS = 57518
const INITIALLY = 57519
const INLINE_P = 57520
const INNER_P = 57521
const INOUT = 57522
const INPUT_P = 57523
const INSENSITIVE = 57524
const INSERT = 57525
const INSTEAD = 57526
const INT_P = 57527
const INTEGER = 57528
const INTERSECT = 57529
const INTERVAL = 57530
const INTO = 57531
const INVOKER = 57532
const IS = 57533
const ISNULL = 57534
const ISOLATION = 57535
const JOIN = 57536
const KEY = 57537
const LABEL = 57538
const LANGUAGE = 57539
const LARGE_P = 57540
const LAST_P = 57541
const LATERAL_P = 57542
const LEADING = 57543
const LEAKPROOF = 57544
const LEAST = 57545
const LEFT = 57546
const LEVEL = 57547
const LIKE = 57548
const LIMIT = 57549
const LISTEN = 57550
const LOAD = 57551
const LOCAL = 57552
const LOCALTIME = 57553
const LOCALTIMESTAMP = 57554
const LOCATION = 57555
const LOCK_P = 57556
const LOCKED = 57557
const LOGGED = 57558
const MAPPING = 57559
const MATCH = 57560
const MATERIALIZED = 57561
const MAXVALUE = 57562
const MINUTE_P = 57563
const MINVALUE = 57564
const MODE = 57565
const MONTH_P = 57566
const MOVE = 57567
const NAME_P = 57568
const NAMES = 57569
const NATIONAL = 57570
const NATURAL = 57571
const NCHAR = 57572
const NEXT = 57573
const NO = 57574
const NONE = 57575
const NOT = 57576
const NOTHING = 57577
const NOTIFY = 57578
const NOTNULL = 57579
const NOWAIT = 57580
const NULL_P = 57581
const NULLIF = 57582
const NULLS_P = 57583
const NUMERIC = 57584
const OBJECT_P = 57585
const OF = 57586
const OFF = 57587
const OFFSET = 57588
const OIDS = 57589
const ON = 57590
const ONLY = 57591
const OPERATOR = 57592
const OPTION = 57593
const OPTIONS = 57594
const OR = 57595
const ORDER = 57596
const ORDINALITY = 57597
const OUT_P = 57598
const OUTER_P = 57599
const OVER = 57600
const OVERLAPS = 57601
const OVERLAY = 57602
const OWNED = 57603
const OWNER = 57604
const PARSER = 57605
const PARTIAL = 57606
const PARTITION = 57607
const PASSING = 57608
const PASSWORD = 57609
const PLACING = 57610
const PLANS = 57611
const POLICY = 57612
const POSITION = 57613
const PRECEDING = 57614
const PRECISION = 57615
const PRESERVE = 57616
const PREPARE = 57617
const PREPARED = 57618
const PRIMARY = 57619
const PRIOR = 57620
const PRIVILEGES = 57621
const PROCEDURAL = 57622
const PROCEDURE = 57623
const PROGRAM = 57624
const QUOTE = 57625
const RANGE = 57626
const READ = 57627
const REAL = 57628
const REASSIGN = 57629
const RECHECK = 57630
const RECURSIVE = 57631
const REF = 57632
const REFERENCES = 57633
const REFRESH = 57634
const REINDEX = 57635
const RELATIVE_P = 57636
const RELEASE = 57637
const RENAME = 57638
const REPEATABLE = 57639
const REPLACE = 57640
const REPLICA = 57641
const RESET = 57642
const RESTART = 57643
const RESTRICT = 57644
const RETURNING = 57645
const RETURNS = 57646
const REVOKE = 57647
const RIGHT = 57648
const ROLE = 57649
const ROLLBACK = 57650
const ROLLUP = 57651
const ROW = 57652
const ROWS = 57653
const RULE = 57654
const SAVEPOINT = 57655
const SCHEMA = 57656
const SCROLL = 57657
const SEARCH = 57658
const SECOND_P = 57659
const SECURITY = 57660
const SELECT = 57661
const SEQUENCE = 57662
const SEQUENCES = 57663
const SERIALIZABLE = 57664
const SERVER = 57665
const SESSION = 57666
const SESSION_USER = 57667
const SET = 57668
const SETS = 57669
const SETOF = 57670
const SHARE = 57671
const SHOW = 57672
const SIMILAR = 57673
const SIMPLE = 57674
const SKIP = 57675
const SMALLINT = 57676
const SNAPSHOT = 57677
const SOME = 57678
const SQL_P = 57679
const STABLE = 57680
const STANDALONE_P = 57681
const START = 57682
const STATEMENT = 57683
const STATISTICS = 57684
const STDIN = 57685
const STDOUT = 57686
const STORAGE = 57687
const STRICT_P = 57688
const STRIP_P = 57689
const SUBSTRING = 57690
const SYMMETRIC = 57691
const SYSID = 57692
const SYSTEM_P = 57693
const TABLE = 57694
const TABLES = 57695
const TABLESAMPLE = 57696
const TABLESPACE = 57697
const TEMP = 57698
const TEMPLATE = 57699
const TEMPORARY = 57700
const TEXT_P = 57701
const THEN = 57702
const TIME = 57703
const TIMESTAMP = 57704
const TO = 57705
const TRAILING = 57706
const TRANSACTION = 57707
const TRANSFORM = 57708
const TREAT = 57709
const TRIGGER = 57710
const TRIM = 57711
const TRUE_P = 57712
const TRUNCATE = 57713
const TRUSTED = 57714
const TYPE_P = 57715
const TYPES_P = 57716
const UNBOUNDED = 57717
const UNCOMMITTED = 57718
const UNENCRYPTED = 57719
const UNION = 57720
const UNIQUE = 57721
const UNKNOWN = 57722
const UNLISTEN = 57723
const UNLOGGED = 57724
const UNTIL = 57725
const UPDATE = 57726
const USER = 57727
const USING = 57728
const VACUUM = 57729
const VALID = 57730
const VALIDATE = 57731
const VALIDATOR = 57732
const VALUE_P = 57733
const VALUES = 57734
const VARCHAR = 57735
const VARIADIC = 57736
const VARYING = 57737
const VERBOSE = 57738
const VERSION_P = 57739
const VIEW = 57740
const VIEWS = 57741
const VOLATILE = 57742
const WHEN = 57743
const WHERE = 57744
const WHITESPACE_P = 57745
const WINDOW = 57746
const WITH = 57747
const WITHIN = 57748
const WITHOUT = 57749
const WORK = 57750
const WRAPPER = 57751
const WRITE = 57752
const XML_P = 57753
const XMLATTRIBUTES = 57754
const XMLCONCAT = 57755
const XMLELEMENT = 57756
const XMLEXISTS = 57757
const XMLFOREST = 57758
const XMLPARSE = 57759
const XMLPI = 57760
const XMLROOT = 57761
const XMLSERIALIZE = 57762
const YEAR_P = 57763
const YES_P = 57764
const ZONE = 57765

var sqlToknames = []string{
	"IDENT",
	"ICONST",
	"SCONST",
	"OP",
	"ABORT_P",
	"ABSOLUTE_P",
	"ACCESS",
	"ACTION",
	"ADD_P",
	"ADMIN",
	"AFTER",
	"AGGREGATE",
	"ALL",
	"ALSO",
	"ALTER",
	"ALWAYS",
	"ANALYSE",
	"ANALYZE",
	"AND",
	"ANY",
	"ARRAY",
	"AS",
	"ASC",
	"ASSERTION",
	"ASSIGNMENT",
	"ASYMMETRIC",
	"AT",
	"ATTRIBUTE",
	"AUTHORIZATION",
	"BACKWARD",
	"BEFORE",
	"BEGIN_P",
	"BETWEEN",
	"BIGINT",
	"BINARY",
	"BIT",
	"BOOLEAN_P",
	"BOTH",
	"BY",
	"CACHE",
	"CALLED",
	"CASCADE",
	"CASCADED",
	"CASE",
	"CAST",
	"CATALOG_P",
	"CHAIN",
	"CHAR_P",
	"CHARACTER",
	"CHARACTERISTICS",
	"CHECK",
	"CHECKPOINT",
	"CLASS",
	"CLOSE",
	"CLUSTER",
	"COALESCE",
	"COLLATE",
	"COLLATION",
	"COLUMN",
	"COMMENT",
	"COMMENTS",
	"COMMIT",
	"COMMITTED",
	"CONCURRENTLY",
	"CONFIGURATION",
	"CONFLICT",
	"CONNECTION",
	"CONSTRAINT",
	"CONSTRAINTS",
	"CONTENT_P",
	"CONTINUE_P",
	"CONVERSION_P",
	"COPY",
	"COST",
	"CREATE",
	"CROSS",
	"CSV",
	"CUBE",
	"CURRENT_P",
	"CURRENT_CATALOG",
	"CURRENT_DATE",
	"CURRENT_ROLE",
	"CURRENT_SCHEMA",
	"CURRENT_TIME",
	"CURRENT_TIMESTAMP",
	"CURRENT_USER",
	"CURSOR",
	"CYCLE",
	"DATA_P",
	"DATABASE",
	"DAY_P",
	"DEALLOCATE",
	"DEC",
	"DECIMAL_P",
	"DECLARE",
	"DEFAULT",
	"DEFAULTS",
	"DEFERRABLE",
	"DEFERRED",
	"DEFINER",
	"DELETE_P",
	"DELIMITER",
	"DELIMITERS",
	"DESC",
	"DICTIONARY",
	"DISABLE_P",
	"DISCARD",
	"DISTINCT",
	"DO",
	"DOCUMENT_P",
	"DOMAIN_P",
	"DOUBLE_P",
	"DROP",
	"EACH",
	"ELSE",
	"ENABLE_P",
	"ENCODING",
	"ENCRYPTED",
	"END_P",
	"ENUM_P",
	"ESCAPE",
	"EVENT",
	"EXCEPT",
	"EXCLUDE",
	"EXCLUDING",
	"EXCLUSIVE",
	"EXECUTE",
	"EXISTS",
	"EXPLAIN",
	"EXTENSION",
	"EXTERNAL",
	"EXTRACT",
	"FALSE_P",
	"FAMILY",
	"FETCH",
	"FILTER",
	"FIRST_P",
	"FLOAT_P",
	"FOLLOWING",
	"FOR",
	"FORCE",
	"FOREIGN",
	"FORWARD",
	"FREEZE",
	"FROM",
	"FULL",
	"FUNCTION",
	"FUNCTIONS",
	"GLOBAL",
	"GRANT",
	"GRANTED",
	"GREATEST",
	"GROUP_P",
	"GROUPING",
	"HANDLER",
	"HAVING",
	"HEADER_P",
	"HOLD",
	"HOUR_P",
	"IDENTITY_P",
	"IF_P",
	"ILIKE",
	"IMMEDIATE",
	"IMMUTABLE",
	"IMPLICIT_P",
	"IMPORT_P",
	"IN_P",
	"INCLUDING",
	"INCREMENT",
	"INDEX",
	"INDEXES",
	"INHERIT",
	"INHERITS",
	"INITIALLY",
	"INLINE_P",
	"INNER_P",
	"INOUT",
	"INPUT_P",
	"INSENSITIVE",
	"INSERT",
	"INSTEAD",
	"INT_P",
	"INTEGER",
	"INTERSECT",
	"INTERVAL",
	"INTO",
	"INVOKER",
	"IS",
	"ISNULL",
	"ISOLATION",
	"JOIN",
	"KEY",
	"LABEL",
	"LANGUAGE",
	"LARGE_P",
	"LAST_P",
	"LATERAL_P",
	"LEADING",
	"LEAKPROOF",
	"LEAST",
	"LEFT",
	"LEVEL",
	"LIKE",
	"LIMIT",
	"LISTEN",
	"LOAD",
	"LOCAL",
	"LOCALTIME",
	"LOCALTIMESTAMP",
	"LOCATION",
	"LOCK_P",
	"LOCKED",
	"LOGGED",
	"MAPPING",
	"MATCH",
	"MATERIALIZED",
	"MAXVALUE",
	"MINUTE_P",
	"MINVALUE",
	"MODE",
	"MONTH_P",
	"MOVE",
	"NAME_P",
	"NAMES",
	"NATIONAL",
	"NATURAL",
	"NCHAR",
	"NEXT",
	"NO",
	"NONE",
	"NOT",
	"NOTHING",
	"NOTIFY",
	"NOTNULL",
	"NOWAIT",
	"NULL_P",
	"NULLIF",
	"NULLS_P",
	"NUMERIC",
	"OBJECT_P",
	"OF",
	"OFF",
	"OFFSET",
	"OIDS",
	"ON",
	"ONLY",
	"OPERATOR",
	"OPTION",
	"OPTIONS",
	"OR",
	"ORDER",
	"ORDINALITY",
	"OUT_P",
	"OUTER_P",
	"OVER",
	"OVERLAPS",
	"OVERLAY",
	"OWNED",
	"OWNER",
	"PARSER",
	"PARTIAL",
	"PARTITION",
	"PASSING",
	"PASSWORD",
	"PLACING",
	"PLANS",
	"POLICY",
	"POSITION",
	"PRECEDING",
	"PRECISION",
	"PRESERVE",
	"PREPARE",
	"PREPARED",
	"PRIMARY",
	"PRIOR",
	"PRIVILEGES",
	"PROCEDURAL",
	"PROCEDURE",
	"PROGRAM",
	"QUOTE",
	"RANGE",
	"READ",
	"REAL",
	"REASSIGN",
	"RECHECK",
	"RECURSIVE",
	"REF",
	"REFERENCES",
	"REFRESH",
	"REINDEX",
	"RELATIVE_P",
	"RELEASE",
	"RENAME",
	"REPEATABLE",
	"REPLACE",
	"REPLICA",
	"RESET",
	"RESTART",
	"RESTRICT",
	"RETURNING",
	"RETURNS",
	"REVOKE",
	"RIGHT",
	"ROLE",
	"ROLLBACK",
	"ROLLUP",
	"ROW",
	"ROWS",
	"RULE",
	"SAVEPOINT",
	"SCHEMA",
	"SCROLL",
	"SEARCH",
	"SECOND_P",
	"SECURITY",
	"SELECT",
	"SEQUENCE",
	"SEQUENCES",
	"SERIALIZABLE",
	"SERVER",
	"SESSION",
	"SESSION_USER",
	"SET",
	"SETS",
	"SETOF",
	"SHARE",
	"SHOW",
	"SIMILAR",
	"SIMPLE",
	"SKIP",
	"SMALLINT",
	"SNAPSHOT",
	"SOME",
	"SQL_P",
	"STABLE",
	"STANDALONE_P",
	"START",
	"STATEMENT",
	"STATISTICS",
	"STDIN",
	"STDOUT",
	"STORAGE",
	"STRICT_P",
	"STRIP_P",
	"SUBSTRING",
	"SYMMETRIC",
	"SYSID",
	"SYSTEM_P",
	"TABLE",
	"TABLES",
	"TABLESAMPLE",
	"TABLESPACE",
	"TEMP",
	"TEMPLATE",
	"TEMPORARY",
	"TEXT_P",
	"THEN",
	"TIME",
	"TIMESTAMP",
	"TO",
	"TRAILING",
	"TRANSACTION",
	"TRANSFORM",
	"TREAT",
	"TRIGGER",
	"TRIM",
	"TRUE_P",
	"TRUNCATE",
	"TRUSTED",
	"TYPE_P",
	"TYPES_P",
	"UNBOUNDED",
	"UNCOMMITTED",
	"UNENCRYPTED",
	"UNION",
	"UNIQUE",
	"UNKNOWN",
	"UNLISTEN",
	"UNLOGGED",
	"UNTIL",
	"UPDATE",
	"USER",
	"USING",
	"VACUUM",
	"VALID",
	"VALIDATE",
	"VALIDATOR",
	"VALUE_P",
	"VALUES",
	"VARCHAR",
	"VARIADIC",
	"VARYING",
	"VERBOSE",
	"VERSION_P",
	"VIEW",
	"VIEWS",
	"VOLATILE",
	"WHEN",
	"WHERE",
	"WHITESPACE_P",
	"WINDOW",
	"WITH",
	"WITHIN",
	"WITHOUT",
	"WORK",
	"WRAPPER",
	"WRITE",
	"XML_P",
	"XMLATTRIBUTES",
	"XMLCONCAT",
	"XMLELEMENT",
	"XMLEXISTS",
	"XMLFOREST",
	"XMLPARSE",
	"XMLPI",
	"XMLROOT",
	"XMLSERIALIZE",
	"YEAR_P",
	"YES_P",
	"ZONE",
}
var sqlStatenames = []string{}

const sqlEofCode = 1
const sqlErrCode = 2
const sqlMaxDepth = 200

//line sql.y:862
// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0

//line yacctab:1
var sqlExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const sqlNprod = 479
const sqlPrivate = 57344

var sqlTokenNames []string
var sqlStates []string

const sqlLast = 910

var sqlAct = []int{

	48, 482, 496, 495, 53, 54, 55, 56, 57, 58,
	59, 60, 392, 61, 62, 63, 393, 394, 395, 396,
	397, 398, 399, 64, 65, 400, 66, 67, 369, 68,
	69, 70, 321, 322, 370, 323, 324, 401, 71, 72,
	73, 74, 75, 402, 403, 76, 77, 325, 326, 78,
	404, 79, 80, 81, 82, 327, 405, 371, 406, 83,
	84, 85, 86, 372, 87, 88, 89, 407, 90, 91,
	92, 93, 94, 95, 408, 373, 96, 97, 98, 409,
	410, 411, 374, 412, 413, 414, 99, 100, 101, 102,
	103, 104, 328, 329, 105, 415, 106, 416, 107, 108,
	109, 110, 111, 417, 112, 113, 114, 418, 419, 115,
	116, 117, 118, 119, 420, 120, 121, 122, 421, 123,
	124, 125, 422, 126, 127, 128, 129, 330, 130, 131,
	132, 331, 423, 133, 424, 134, 135, 332, 136, 425,
	137, 426, 138, 375, 427, 376, 139, 140, 141, 428,
	142, 333, 429, 334, 143, 430, 144, 145, 146, 147,
	148, 377, 149, 150, 151, 152, 431, 153, 154, 155,
	156, 157, 158, 432, 159, 378, 335, 160, 161, 162,
	163, 336, 337, 433, 338, 434, 164, 379, 380, 165,
	381, 166, 167, 168, 169, 170, 435, 436, 171, 339,
	382, 172, 383, 437, 173, 174, 175, 438, 439, 176,
	177, 178, 179, 180, 181, 182, 183, 184, 185, 186,
	187, 188, 189, 190, 340, 384, 341, 191, 192, 342,
	440, 193, 194, 385, 195, 441, 343, 196, 344, 197,
	198, 199, 442, 200, 443, 444, 201, 202, 203, 445,
	446, 204, 345, 386, 205, 387, 346, 206, 207, 208,
	209, 210, 211, 212, 447, 213, 214, 347, 215, 348,
	218, 216, 217, 448, 219, 220, 221, 222, 223, 224,
	225, 226, 349, 227, 228, 229, 230, 449, 231, 232,
	233, 234, 235, 236, 237, 238, 239, 240, 241, 450,
	242, 243, 388, 244, 245, 246, 350, 247, 248, 249,
	250, 251, 252, 253, 254, 451, 255, 256, 257, 258,
	259, 452, 260, 261, 351, 262, 263, 389, 264, 265,
	352, 266, 453, 267, 268, 269, 270, 271, 272, 273,
	274, 275, 276, 277, 353, 454, 278, 279, 455, 280,
	390, 281, 282, 283, 284, 285, 456, 354, 355, 457,
	458, 286, 287, 356, 288, 357, 459, 289, 290, 291,
	292, 293, 294, 295, 460, 461, 296, 297, 298, 299,
	300, 462, 463, 301, 302, 303, 304, 305, 358, 359,
	464, 306, 391, 307, 308, 309, 310, 465, 466, 311,
	467, 468, 312, 313, 314, 315, 316, 317, 360, 361,
	362, 363, 364, 365, 366, 367, 368, 318, 319, 320,
	25, 23, 14, 21, 20, 13, 14, 21, 20, 13,
	41, 30, 475, 27, 474, 491, 33, 7, 8, 34,
	35, 488, 4, 30, 30, 31, 31, 32, 13, 39,
	479, 469, 470, 471, 45, 478, 476, 30, 31, 31,
	11, 29, 477, 489, 2, 480, 497, 12, 6, 494,
	481, 472, 31, 52, 22, 44, 26, 51, 484, 50,
	49, 47, 28, 15, 36, 16, 38, 37, 46, 24,
	493, 5, 10, 9, 3, 1, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 490, 43, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 487,
	42, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 18, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 32, 32, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 32,
	32, 8, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 32, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 4, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 40, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 17, 0,
	0, 0, 17, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 19, 0, 0, 0, 19, 0, 0,
	473, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 483, 483, 0, 0,
	485, 486, 0, 0, 0, 0, 0, 0, 0, 492,
}
var sqlPact = []int{

	123, -1000, -1000, 289, 422, 36, -1000, 422, 422, -1000,
	6, -1000, 457, 436, 12, -1000, -1000, -1000, 422, 418,
	-1000, -1000, 195, 351, -1000, 450, -1000, 422, -4, -1000,
	422, 422, 422, 467, -1000, 424, 8, -1000, 5, 414,
	422, 261, 256, 422, 466, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 423,
	194, -1000, -1000, -1000, -1000, 422, 422, -1000, 422, 422,
	193, -1000, -1000, 437, -1000, -1000, -1000, 10, 422, -1000,
	-1000, 465, 436, -424, -1000, 462, -1000, -1000,
}
var sqlPgo = []int{

	0, 495, 464, 494, 493, 492, 421, 420, 460, 467,
	491, 490, 489, 468, 1, 487, 486, 485, 485, 483,
	481, 480, 479, 477, 473, 473, 473, 473,
}
var sqlR1 = []int{

	0, 1, 2, 2, 3, 6, 6, 6, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 9,
	11, 11, 12, 12, 12, 12, 12, 10, 10, 14,
	14, 14, 15, 15, 16, 16, 13, 13, 4, 4,
	5, 5, 8, 8, 8, 17, 19, 18, 25, 25,
	25, 26, 26, 26, 27, 27, 27, 27, 20, 20,
	20, 20, 20, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 22, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 22, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 22, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 22, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 22, 22, 22, 22, 22, 22, 22, 22, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 24, 24, 24, 24, 24, 24, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 24, 24,
}
var sqlR2 = []int{

	0, 1, 4, 2, 2, 1, 3, 2, 1, 3,
	1, 1, 1, 3, 3, 3, 2, 3, 3, 1,
	1, 3, 3, 4, 4, 7, 5, 2, 2, 1,
	2, 2, 0, 1, 3, 3, 2, 0, 1, 0,
	1, 3, 3, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1,
}
var sqlChk = []int{

	-1000, -1, -2, -3, 319, -10, -13, 148, 402, -4,
	-5, -8, -9, -7, 4, -19, -17, 370, 234, 425,
	6, 5, -13, -6, -12, -7, -9, 427, 25, 4,
	7, 22, 253, 424, -7, -7, -2, -15, -16, 254,
	427, 79, 229, 194, 25, 4, -8, -20, 4, -21,
	-22, -23, -24, 8, 9, 10, 11, 12, 13, 14,
	15, 17, 18, 19, 27, 28, 30, 31, 33, 34,
	35, 42, 43, 44, 45, 46, 49, 50, 53, 55,
	56, 57, 58, 63, 64, 65, 66, 68, 69, 70,
	72, 73, 74, 75, 76, 77, 80, 81, 82, 90,
	91, 92, 93, 94, 95, 98, 100, 102, 103, 104,
	105, 106, 108, 109, 110, 113, 114, 115, 116, 117,
	119, 120, 121, 123, 124, 125, 127, 128, 129, 130,
	132, 133, 134, 137, 139, 140, 142, 144, 146, 150,
	151, 152, 154, 158, 160, 161, 162, 163, 164, 166,
	167, 168, 169, 171, 172, 173, 174, 175, 176, 178,
	181, 182, 183, 184, 190, 193, 195, 196, 197, 198,
	199, 202, 205, 208, 209, 210, 213, 214, 215, 216,
	217, 218, 219, 220, 221, 222, 223, 224, 225, 226,
	227, 231, 232, 235, 236, 238, 241, 243, 244, 245,
	247, 250, 251, 252, 255, 258, 261, 262, 263, 264,
	265, 266, 267, 269, 270, 272, 275, 276, 274, 278,
	279, 280, 281, 282, 283, 284, 285, 287, 288, 289,
	290, 292, 293, 294, 295, 296, 297, 298, 299, 300,
	301, 302, 304, 305, 307, 308, 309, 311, 312, 313,
	314, 315, 316, 317, 318, 320, 321, 322, 323, 324,
	326, 327, 329, 330, 332, 333, 335, 337, 338, 339,
	340, 341, 342, 343, 344, 345, 346, 347, 350, 351,
	353, 355, 356, 357, 358, 359, 365, 366, 368, 371,
	372, 373, 374, 375, 376, 377, 380, 381, 382, 383,
	384, 387, 388, 389, 390, 391, 395, 397, 398, 399,
	400, 403, 406, 407, 408, 409, 410, 411, 421, 422,
	423, 36, 37, 39, 40, 51, 52, 59, 96, 97,
	131, 135, 141, 155, 157, 180, 185, 186, 188, 203,
	228, 230, 233, 240, 242, 256, 260, 271, 273, 286,
	310, 328, 334, 348, 361, 362, 367, 369, 392, 393,
	412, 413, 414, 415, 416, 417, 418, 419, 420, 32,
	38, 61, 67, 79, 86, 147, 149, 165, 179, 191,
	192, 194, 204, 206, 229, 237, 257, 259, 306, 331,
	354, 396, 16, 20, 21, 22, 23, 24, 25, 26,
	29, 41, 47, 48, 54, 60, 62, 71, 78, 83,
	84, 85, 87, 88, 89, 99, 101, 107, 111, 112,
	118, 122, 126, 136, 138, 143, 145, 148, 153, 156,
	159, 170, 177, 187, 189, 200, 201, 207, 211, 212,
	234, 239, 246, 248, 249, 253, 254, 268, 277, 291,
	303, 319, 325, 336, 349, 352, 360, 363, 364, 370,
	378, 379, 385, 386, 394, 401, 402, 404, 405, -7,
	-7, -7, 4, 426, 426, 427, 42, -6, 194, 194,
	-6, 4, -14, -7, -14, -6, -6, 386, 248, 26,
	107, 425, -7, -11, 4, 427, 426, 4,
}
var sqlDef = []int{

	0, -2, 1, 37, 39, 37, 3, 0, 0, 4,
	38, 40, 44, 19, 8, 10, 11, 12, 0, 0,
	46, 45, 32, 27, 28, 5, 36, 0, 0, 43,
	0, 0, 0, 0, 16, 0, 0, 2, 33, 0,
	0, 0, 0, 0, 0, 7, 41, 42, 58, 59,
	60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
	70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
	80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
	100, 101, 102, 103, 104, 105, 106, 107, 108, 109,
	110, 111, 112, 113, 114, 115, 116, 117, 118, 119,
	120, 121, 122, 123, 124, 125, 126, 127, 128, 129,
	130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
	140, 141, 142, 143, 144, 145, 146, 147, 148, 149,
	150, 151, 152, 153, 154, 155, 156, 157, 158, 159,
	160, 161, 162, 163, 164, 165, 166, 167, 168, 169,
	170, 171, 172, 173, 174, 175, 176, 177, 178, 179,
	180, 181, 182, 183, 184, 185, 186, 187, 188, 189,
	190, 191, 192, 193, 194, 195, 196, 197, 198, 199,
	200, 201, 202, 203, 204, 205, 206, 207, 208, 209,
	210, 211, 212, 213, 214, 215, 216, 217, 218, 219,
	220, 221, 222, 223, 224, 225, 226, 227, 228, 229,
	230, 231, 232, 233, 234, 235, 236, 237, 238, 239,
	240, 241, 242, 243, 244, 245, 246, 247, 248, 249,
	250, 251, 252, 253, 254, 255, 256, 257, 258, 259,
	260, 261, 262, 263, 264, 265, 266, 267, 268, 269,
	270, 271, 272, 273, 274, 275, 276, 277, 278, 279,
	280, 281, 282, 283, 284, 285, 286, 287, 288, 289,
	290, 291, 292, 293, 294, 295, 296, 297, 298, 299,
	300, 301, 302, 303, 304, 305, 306, 307, 308, 309,
	310, 311, 312, 313, 314, 315, 316, 317, 318, 319,
	320, 321, 322, 323, 324, 325, 326, 327, 328, 329,
	330, 331, 332, 333, 334, 335, 336, 337, 338, 339,
	340, 341, 342, 343, 344, 345, 346, 347, 348, 349,
	350, 351, 352, 353, 354, 355, 356, 357, 358, 359,
	360, 361, 362, 363, 364, 365, 366, 367, 368, 369,
	370, 371, 372, 373, 374, 375, 376, 377, 378, 379,
	380, 381, 382, 383, 384, 385, 386, 387, 388, 389,
	390, 391, 392, 393, 394, 395, 396, 397, 398, 399,
	400, 401, 402, 403, 404, 405, 406, 407, 408, 409,
	410, 411, 412, 413, 414, 415, 416, 417, 418, 419,
	420, 421, 422, 423, 424, 425, 426, 427, 428, 429,
	430, 431, 432, 433, 434, 435, 436, 437, 438, 439,
	440, 441, 442, 443, 444, 445, 446, 447, 448, 449,
	450, 451, 452, 453, 454, 455, 456, 457, 458, 459,
	460, 461, 462, 463, 464, 465, 466, 467, 468, 469,
	470, 471, 472, 473, 474, 475, 476, 477, 478, 13,
	14, 15, 9, 17, 18, 0, 0, 22, 0, 0,
	0, 6, 35, 29, 34, 23, 24, 0, 0, 30,
	31, 0, 26, 0, 20, 0, 25, 21,
}
var sqlTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	425, 426, 3, 3, 427, 3, 424,
}
var sqlTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
	92, 93, 94, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 120, 121,
	122, 123, 124, 125, 126, 127, 128, 129, 130, 131,
	132, 133, 134, 135, 136, 137, 138, 139, 140, 141,
	142, 143, 144, 145, 146, 147, 148, 149, 150, 151,
	152, 153, 154, 155, 156, 157, 158, 159, 160, 161,
	162, 163, 164, 165, 166, 167, 168, 169, 170, 171,
	172, 173, 174, 175, 176, 177, 178, 179, 180, 181,
	182, 183, 184, 185, 186, 187, 188, 189, 190, 191,
	192, 193, 194, 195, 196, 197, 198, 199, 200, 201,
	202, 203, 204, 205, 206, 207, 208, 209, 210, 211,
	212, 213, 214, 215, 216, 217, 218, 219, 220, 221,
	222, 223, 224, 225, 226, 227, 228, 229, 230, 231,
	232, 233, 234, 235, 236, 237, 238, 239, 240, 241,
	242, 243, 244, 245, 246, 247, 248, 249, 250, 251,
	252, 253, 254, 255, 256, 257,
}
var sqlTok3 = []int{
	57600, 258, 57601, 259, 57602, 260, 57603, 261, 57604, 262,
	57605, 263, 57606, 264, 57607, 265, 57608, 266, 57609, 267,
	57610, 268, 57611, 269, 57612, 270, 57613, 271, 57614, 272,
	57615, 273, 57616, 274, 57617, 275, 57618, 276, 57619, 277,
	57620, 278, 57621, 279, 57622, 280, 57623, 281, 57624, 282,
	57625, 283, 57626, 284, 57627, 285, 57628, 286, 57629, 287,
	57630, 288, 57631, 289, 57632, 290, 57633, 291, 57634, 292,
	57635, 293, 57636, 294, 57637, 295, 57638, 296, 57639, 297,
	57640, 298, 57641, 299, 57642, 300, 57643, 301, 57644, 302,
	57645, 303, 57646, 304, 57647, 305, 57648, 306, 57649, 307,
	57650, 308, 57651, 309, 57652, 310, 57653, 311, 57654, 312,
	57655, 313, 57656, 314, 57657, 315, 57658, 316, 57659, 317,
	57660, 318, 57661, 319, 57662, 320, 57663, 321, 57664, 322,
	57665, 323, 57666, 324, 57667, 325, 57668, 326, 57669, 327,
	57670, 328, 57671, 329, 57672, 330, 57673, 331, 57674, 332,
	57675, 333, 57676, 334, 57677, 335, 57678, 336, 57679, 337,
	57680, 338, 57681, 339, 57682, 340, 57683, 341, 57684, 342,
	57685, 343, 57686, 344, 57687, 345, 57688, 346, 57689, 347,
	57690, 348, 57691, 349, 57692, 350, 57693, 351, 57694, 352,
	57695, 353, 57696, 354, 57697, 355, 57698, 356, 57699, 357,
	57700, 358, 57701, 359, 57702, 360, 57703, 361, 57704, 362,
	57705, 363, 57706, 364, 57707, 365, 57708, 366, 57709, 367,
	57710, 368, 57711, 369, 57712, 370, 57713, 371, 57714, 372,
	57715, 373, 57716, 374, 57717, 375, 57718, 376, 57719, 377,
	57720, 378, 57721, 379, 57722, 380, 57723, 381, 57724, 382,
	57725, 383, 57726, 384, 57727, 385, 57728, 386, 57729, 387,
	57730, 388, 57731, 389, 57732, 390, 57733, 391, 57734, 392,
	57735, 393, 57736, 394, 57737, 395, 57738, 396, 57739, 397,
	57740, 398, 57741, 399, 57742, 400, 57743, 401, 57744, 402,
	57745, 403, 57746, 404, 57747, 405, 57748, 406, 57749, 407,
	57750, 408, 57751, 409, 57752, 410, 57753, 411, 57754, 412,
	57755, 413, 57756, 414, 57757, 415, 57758, 416, 57759, 417,
	57760, 418, 57761, 419, 57762, 420, 57763, 421, 57764, 422,
	57765, 423, 0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var sqlDebug = 0

type sqlLexer interface {
	Lex(lval *sqlSymType) int
	Error(s string)
}

const sqlFlag = -1000

func sqlTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(sqlToknames) {
		if sqlToknames[c-4] != "" {
			return sqlToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func sqlStatname(s int) string {
	if s >= 0 && s < len(sqlStatenames) {
		if sqlStatenames[s] != "" {
			return sqlStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func sqllex1(lex sqlLexer, lval *sqlSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = sqlTok1[0]
		goto out
	}
	if char < len(sqlTok1) {
		c = sqlTok1[char]
		goto out
	}
	if char >= sqlPrivate {
		if char < sqlPrivate+len(sqlTok2) {
			c = sqlTok2[char-sqlPrivate]
			goto out
		}
	}
	for i := 0; i < len(sqlTok3); i += 2 {
		c = sqlTok3[i+0]
		if c == char {
			c = sqlTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = sqlTok2[1] /* unknown char */
	}
	if sqlDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", sqlTokname(c), uint(char))
	}
	return c
}

func sqlParse(sqllex sqlLexer) int {
	var sqln int
	var sqllval sqlSymType
	var sqlVAL sqlSymType
	sqlS := make([]sqlSymType, sqlMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	sqlstate := 0
	sqlchar := -1
	sqlp := -1
	goto sqlstack

ret0:
	return 0

ret1:
	return 1

sqlstack:
	/* put a state and value onto the stack */
	if sqlDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", sqlTokname(sqlchar), sqlStatname(sqlstate))
	}

	sqlp++
	if sqlp >= len(sqlS) {
		nyys := make([]sqlSymType, len(sqlS)*2)
		copy(nyys, sqlS)
		sqlS = nyys
	}
	sqlS[sqlp] = sqlVAL
	sqlS[sqlp].yys = sqlstate

sqlnewstate:
	sqln = sqlPact[sqlstate]
	if sqln <= sqlFlag {
		goto sqldefault /* simple state */
	}
	if sqlchar < 0 {
		sqlchar = sqllex1(sqllex, &sqllval)
	}
	sqln += sqlchar
	if sqln < 0 || sqln >= sqlLast {
		goto sqldefault
	}
	sqln = sqlAct[sqln]
	if sqlChk[sqln] == sqlchar { /* valid shift */
		sqlchar = -1
		sqlVAL = sqllval
		sqlstate = sqln
		if Errflag > 0 {
			Errflag--
		}
		goto sqlstack
	}

sqldefault:
	/* default state action */
	sqln = sqlDef[sqlstate]
	if sqln == -2 {
		if sqlchar < 0 {
			sqlchar = sqllex1(sqllex, &sqllval)
		}

		/* look through exception table */
		xi := 0
		for {
			if sqlExca[xi+0] == -1 && sqlExca[xi+1] == sqlstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			sqln = sqlExca[xi+0]
			if sqln < 0 || sqln == sqlchar {
				break
			}
		}
		sqln = sqlExca[xi+1]
		if sqln < 0 {
			goto ret0
		}
	}
	if sqln == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			sqllex.Error("syntax error")
			Nerrs++
			if sqlDebug >= 1 {
				__yyfmt__.Printf("%s", sqlStatname(sqlstate))
				__yyfmt__.Printf(" saw %s\n", sqlTokname(sqlchar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for sqlp >= 0 {
				sqln = sqlPact[sqlS[sqlp].yys] + sqlErrCode
				if sqln >= 0 && sqln < sqlLast {
					sqlstate = sqlAct[sqln] /* simulate a shift of "error" */
					if sqlChk[sqlstate] == sqlErrCode {
						goto sqlstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if sqlDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", sqlS[sqlp].yys)
				}
				sqlp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if sqlDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", sqlTokname(sqlchar))
			}
			if sqlchar == sqlEofCode {
				goto ret1
			}
			sqlchar = -1
			goto sqlnewstate /* try again in the same state */
		}
	}

	/* reduction by production sqln */
	if sqlDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", sqln, sqlStatname(sqlstate))
	}

	sqlnt := sqln
	sqlpt := sqlp
	_ = sqlpt // guard against "declared and not used"

	sqlp -= sqlR2[sqln]
	sqlVAL = sqlS[sqlp+1]

	/* consult goto table to find next state */
	sqln = sqlR1[sqln]
	sqlg := sqlPgo[sqln]
	sqlj := sqlg + sqlS[sqlp].yys + 1

	if sqlj >= sqlLast {
		sqlstate = sqlAct[sqlg]
	} else {
		sqlstate = sqlAct[sqlj]
		if sqlChk[sqlstate] != -sqln {
			sqlstate = sqlAct[sqlg]
		}
	}
	// dummy call; replaced with literal code
	switch sqlnt {

	case 1:
		//line sql.y:149
		{
			sqlVAL.sqlSelect = sqlS[sqlpt-0].sqlSelect
		}
	case 2:
		//line sql.y:155
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.TargetList = sqlS[sqlpt-3].fields
			sqlVAL.sqlSelect.FromClause = sqlS[sqlpt-2].fromClause
			sqlVAL.sqlSelect.WhereClause = sqlS[sqlpt-1].whereClause
			sqlVAL.sqlSelect.OrderClause = sqlS[sqlpt-0].orderClause
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 3:
		//line sql.y:164
		{
			sqlVAL.sqlSelect = &SelectStmt{}
			sqlVAL.sqlSelect.TargetList = sqlS[sqlpt-1].fields
			sqlVAL.sqlSelect.WhereClause = sqlS[sqlpt-0].whereClause
			sqllex.(*sqlLex).stmt = sqlVAL.sqlSelect
		}
	case 4:
		//line sql.y:173
		{
			sqlVAL.fields = sqlS[sqlpt-0].fields
		}
	case 5:
		//line sql.y:179
		{
			sqlVAL.expr = sqlS[sqlpt-0].expr
		}
	case 6:
		//line sql.y:183
		{
			sqlVAL.expr = AliasedExpr{Expr: sqlS[sqlpt-2].expr, Alias: sqlS[sqlpt-0].src}
		}
	case 7:
		//line sql.y:187
		{
			sqlVAL.expr = AliasedExpr{Expr: sqlS[sqlpt-1].expr, Alias: sqlS[sqlpt-0].src}
		}
	case 8:
		//line sql.y:193
		{
			sqlVAL.expr = ColumnRef{Column: sqlS[sqlpt-0].src}
		}
	case 9:
		//line sql.y:197
		{
			sqlVAL.expr = ColumnRef{Table: sqlS[sqlpt-2].src, Column: sqlS[sqlpt-0].src}
		}
	case 10:
		//line sql.y:201
		{
			sqlVAL.expr = StringLiteral(sqlS[sqlpt-0].src)
		}
	case 11:
		//line sql.y:205
		{
			sqlVAL.expr = IntegerLiteral(sqlS[sqlpt-0].src)
		}
	case 12:
		//line sql.y:209
		{
			sqlVAL.expr = ColumnRef{Column: "true"}
		}
	case 13:
		//line sql.y:213
		{
			sqlVAL.expr = BinaryExpr{Left: sqlS[sqlpt-2].expr, Operator: sqlS[sqlpt-1].src, Right: sqlS[sqlpt-0].expr}
		}
	case 14:
		//line sql.y:217
		{
			sqlVAL.expr = BinaryExpr{Left: sqlS[sqlpt-2].expr, Operator: "and", Right: sqlS[sqlpt-0].expr}
		}
	case 15:
		//line sql.y:221
		{
			sqlVAL.expr = BinaryExpr{Left: sqlS[sqlpt-2].expr, Operator: "or", Right: sqlS[sqlpt-0].expr}
		}
	case 16:
		//line sql.y:225
		{
			sqlVAL.expr = NotExpr{Expr: sqlS[sqlpt-0].expr}
		}
	case 17:
		//line sql.y:229
		{
			sqlVAL.expr = ParenExpr{Expr: sqlS[sqlpt-1].expr}
		}
	case 18:
		//line sql.y:233
		{
			sqlVAL.expr = ParenExpr{Expr: sqlS[sqlpt-1].sqlSelect}
		}
	case 19:
		sqlVAL.expr = sqlS[sqlpt-0].expr
	case 20:
		//line sql.y:241
		{
			sqlVAL.identifiers = []string{sqlS[sqlpt-0].src}
		}
	case 21:
		//line sql.y:245
		{
			sqlVAL.identifiers = append(sqlS[sqlpt-2].identifiers, sqlS[sqlpt-0].src)
		}
	case 22:
		//line sql.y:251
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-2].expr, Join: ",", Right: sqlS[sqlpt-0].expr}
		}
	case 23:
		//line sql.y:255
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-3].expr, Join: "cross join", Right: sqlS[sqlpt-0].expr}
		}
	case 24:
		//line sql.y:259
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-3].expr, Join: "natural join", Right: sqlS[sqlpt-0].expr}
		}
	case 25:
		//line sql.y:263
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-6].expr, Join: "join", Right: sqlS[sqlpt-4].expr, Using: sqlS[sqlpt-1].identifiers}
		}
	case 26:
		//line sql.y:267
		{
			sqlVAL.expr = JoinExpr{Left: sqlS[sqlpt-4].expr, Join: "join", Right: sqlS[sqlpt-2].expr, On: sqlS[sqlpt-0].expr}
		}
	case 27:
		//line sql.y:273
		{
			sqlVAL.fromClause = &FromClause{Expr: sqlS[sqlpt-0].expr}
		}
	case 28:
		//line sql.y:277
		{
			sqlVAL.fromClause = &FromClause{Expr: sqlS[sqlpt-0].expr}
		}
	case 29:
		//line sql.y:283
		{
			sqlVAL.orderExpr = OrderExpr{Expr: sqlS[sqlpt-0].expr}
		}
	case 30:
		//line sql.y:287
		{
			sqlVAL.orderExpr = OrderExpr{Expr: sqlS[sqlpt-1].expr, Order: "asc"}
		}
	case 31:
		//line sql.y:291
		{
			sqlVAL.orderExpr = OrderExpr{Expr: sqlS[sqlpt-1].expr, Order: "desc"}
		}
	case 32:
		//line sql.y:297
		{
			sqlVAL.orderClause = nil
		}
	case 33:
		sqlVAL.orderClause = sqlS[sqlpt-0].orderClause
	case 34:
		//line sql.y:304
		{
			sqlVAL.orderClause = &OrderClause{Exprs: []OrderExpr{sqlS[sqlpt-0].orderExpr}}
		}
	case 35:
		//line sql.y:308
		{
			sqlS[sqlpt-2].orderClause.Exprs = append(sqlS[sqlpt-2].orderClause.Exprs, sqlS[sqlpt-0].orderExpr)
			sqlVAL.orderClause = sqlS[sqlpt-2].orderClause
		}
	case 36:
		//line sql.y:316
		{
			sqlVAL.whereClause = &WhereClause{Expr: sqlS[sqlpt-0].expr}
		}
	case 37:
		//line sql.y:317
		{
			sqlVAL.whereClause = nil
		}
	case 38:
		//line sql.y:327
		{
			sqlVAL.fields = sqlS[sqlpt-0].fields
		}
	case 39:
		//line sql.y:328
		{
			sqlVAL.fields = nil
		}
	case 40:
		//line sql.y:331
		{
			sqlVAL.fields = []Expr{sqlS[sqlpt-0].expr}
		}
	case 41:
		//line sql.y:333
		{
			sqlVAL.fields = append(sqlS[sqlpt-2].fields, sqlS[sqlpt-0].expr)
		}
	case 42:
		//line sql.y:339
		{
			sqlVAL.expr = AliasedExpr{Expr: sqlS[sqlpt-2].expr, Alias: sqlS[sqlpt-0].src}
		}
	case 43:
		//line sql.y:343
		{
			sqlVAL.expr = AliasedExpr{Expr: sqlS[sqlpt-1].expr, Alias: sqlS[sqlpt-0].src}
		}
	case 44:
		sqlVAL.expr = sqlS[sqlpt-0].expr
	case 45:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 46:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 47:
		//line sql.y:355
		{
			sqlVAL.src = sqlS[sqlpt-0].src
		}
	case 58:
		//line sql.y:398
		{
			sqlVAL.src = sqlS[sqlpt-0].src
		}
	case 59:
		//line sql.y:399
		{
			sqlVAL.src = sqlS[sqlpt-0].src
		}
	case 60:
		//line sql.y:400
		{
			sqlVAL.src = sqlS[sqlpt-0].src
		}
	case 61:
		//line sql.y:401
		{
			sqlVAL.src = sqlS[sqlpt-0].src
		}
	case 62:
		//line sql.y:402
		{
			sqlVAL.src = sqlS[sqlpt-0].src
		}
	case 63:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 64:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 65:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 66:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 67:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 68:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 69:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 70:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 71:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 72:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 73:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 74:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 75:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 76:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 77:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 78:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 79:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 80:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 81:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 82:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 83:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 84:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 85:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 86:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 87:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 88:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 89:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 90:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 91:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 92:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 93:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 94:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 95:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 96:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 97:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 98:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 99:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 100:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 101:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 102:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 103:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 104:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 105:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 106:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 107:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 108:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 109:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 110:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 111:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 112:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 113:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 114:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 115:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 116:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 117:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 118:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 119:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 120:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 121:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 122:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 123:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 124:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 125:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 126:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 127:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 128:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 129:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 130:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 131:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 132:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 133:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 134:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 135:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 136:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 137:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 138:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 139:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 140:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 141:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 142:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 143:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 144:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 145:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 146:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 147:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 148:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 149:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 150:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 151:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 152:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 153:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 154:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 155:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 156:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 157:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 158:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 159:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 160:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 161:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 162:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 163:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 164:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 165:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 166:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 167:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 168:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 169:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 170:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 171:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 172:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 173:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 174:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 175:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 176:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 177:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 178:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 179:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 180:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 181:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 182:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 183:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 184:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 185:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 186:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 187:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 188:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 189:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 190:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 191:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 192:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 193:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 194:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 195:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 196:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 197:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 198:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 199:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 200:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 201:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 202:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 203:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 204:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 205:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 206:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 207:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 208:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 209:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 210:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 211:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 212:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 213:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 214:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 215:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 216:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 217:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 218:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 219:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 220:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 221:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 222:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 223:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 224:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 225:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 226:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 227:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 228:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 229:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 230:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 231:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 232:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 233:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 234:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 235:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 236:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 237:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 238:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 239:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 240:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 241:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 242:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 243:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 244:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 245:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 246:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 247:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 248:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 249:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 250:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 251:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 252:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 253:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 254:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 255:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 256:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 257:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 258:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 259:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 260:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 261:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 262:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 263:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 264:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 265:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 266:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 267:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 268:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 269:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 270:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 271:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 272:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 273:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 274:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 275:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 276:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 277:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 278:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 279:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 280:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 281:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 282:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 283:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 284:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 285:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 286:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 287:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 288:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 289:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 290:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 291:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 292:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 293:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 294:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 295:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 296:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 297:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 298:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 299:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 300:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 301:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 302:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 303:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 304:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 305:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 306:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 307:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 308:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 309:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 310:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 311:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 312:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 313:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 314:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 315:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 316:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 317:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 318:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 319:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 320:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 321:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 322:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 323:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 324:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 325:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 326:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 327:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 328:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 329:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 330:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 331:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 332:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 333:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 334:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 335:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 336:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 337:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 338:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 339:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 340:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 341:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 342:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 343:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 344:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 345:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 346:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 347:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 348:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 349:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 350:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 351:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 352:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 353:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 354:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 355:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 356:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 357:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 358:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 359:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 360:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 361:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 362:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 363:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 364:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 365:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 366:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 367:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 368:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 369:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 370:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 371:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 372:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 373:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 374:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 375:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 376:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 377:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 378:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 379:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 380:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 381:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 382:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 383:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 384:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 385:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 386:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 387:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 388:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 389:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 390:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 391:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 392:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 393:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 394:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 395:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 396:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 397:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 398:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 399:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 400:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 401:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 402:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 403:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 404:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 405:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 406:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 407:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 408:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 409:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 410:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 411:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 412:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 413:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 414:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 415:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 416:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 417:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 418:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 419:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 420:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 421:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 422:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 423:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 424:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 425:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 426:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 427:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 428:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 429:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 430:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 431:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 432:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 433:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 434:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 435:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 436:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 437:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 438:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 439:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 440:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 441:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 442:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 443:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 444:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 445:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 446:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 447:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 448:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 449:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 450:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 451:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 452:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 453:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 454:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 455:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 456:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 457:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 458:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 459:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 460:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 461:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 462:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 463:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 464:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 465:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 466:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 467:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 468:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 469:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 470:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 471:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 472:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 473:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 474:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 475:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 476:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 477:
		sqlVAL.src = sqlS[sqlpt-0].src
	case 478:
		sqlVAL.src = sqlS[sqlpt-0].src
	}
	goto sqlstack /* stack new state and value */
}
