package main

import (
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

type stateFn func(*sqlLex) stateFn

type token struct {
	typ int
	src string
}

type sqlLex struct {
	src    string
	start  int
	pos    int
	width  int
	state  stateFn
	tokens []token
	stmt   *SelectStmt
}

func (x *sqlLex) Lex(yylval *sqlSymType) int {
	token := x.tokens[0]
	x.tokens = x.tokens[1:]

	yylval.src = token.src
	return token.typ
}

// The parser calls this method on a parse error.
func (x *sqlLex) Error(s string) {
	log.Printf("parse error: %s at character %d", s, x.start)
}

func NewSqlLexer(src string) *sqlLex {
	x := &sqlLex{src: src,
		tokens: make([]token, 0),
		state:  blankState,
	}

	for x.state != nil {
		x.state = x.state(x)
	}

	x.append(token{typ: eof})

	return x
}

func (l *sqlLex) append(t token) {
	l.tokens = append(l.tokens, t)

	if len(l.tokens) == 1 {
		return
	}

	prevToken := &l.tokens[len(l.tokens)-2]

	switch t.typ {
	case BETWEEN, IN_P, LIKE, ILIKE, SIMILAR:
		if prevToken.typ == NOT {
			prevToken.typ = NOT_LA
		}
	case FIRST_P, LAST_P:
		if prevToken.typ == NULLS_P {
			prevToken.typ = NULLS_LA
		}
	case TIME, ORDINALITY:
		if prevToken.typ == WITH {
			prevToken.typ = WITH_LA
		}
	}
}

func (l *sqlLex) next() (r rune) {
	if l.pos >= len(l.src) {
		l.width = 0 // because backing up from having read eof should read eof again
		return 0
	}

	r, l.width = utf8.DecodeRuneInString(l.src[l.pos:])
	l.pos += l.width

	return r
}

func (l *sqlLex) unnext() {
	l.pos -= l.width
}

func (l *sqlLex) ignore() {
	l.start = l.pos
}

func (l *sqlLex) acceptRunFunc(f func(rune) bool) {
	for f(l.next()) {
	}
	l.unnext()
}

func blankState(l *sqlLex) stateFn {
	switch r := l.next(); {
	case r == 0:
		return nil
	case r == ',' || r == '.' || r == '(' || r == ')':
		return lexSimple
	case r == '\'':
		return lexStringLiteral
	case r == '"':
		return lexQuotedIdentifier
	case isOperator(r):
		return lexOperator
	case unicode.IsDigit(r):
		return lexNumber
	case isWhitespace(r):
		l.skipWhitespace()
		return blankState
	case isAlphanumeric(r):
		return lexAlphanumeric
	}
	return nil
}

func lexNumber(l *sqlLex) stateFn {
	l.acceptRunFunc(unicode.IsDigit)
	t := token{src: l.src[l.start:l.pos], typ: ICONST}
	l.append(t)
	l.start = l.pos
	return blankState
}

func lexAlphanumeric(l *sqlLex) stateFn {
	l.acceptRunFunc(isAlphanumeric)

	t := token{src: l.src[l.start:l.pos]}

	if typ, ok := keywords[strings.ToLower(t.src)]; ok {
		t.typ = typ
	} else {
		t.typ = IDENT
	}

	l.append(t)
	l.start = l.pos
	return blankState
}

func lexStringLiteral(l *sqlLex) stateFn {
	for {
		var r rune
		r = l.next()
		if r == 0 {
			return nil // error for EOF inside of string literal
		}

		if r == '\'' {
			r = l.next()
			if r != '\'' {
				l.unnext()
				t := token{src: l.src[l.start:l.pos]}
				t.typ = SCONST
				l.append(t)
				l.start = l.pos
				return blankState
			}
		}
	}
}

func lexQuotedIdentifier(l *sqlLex) stateFn {
	for {
		var r rune
		r = l.next()
		if r == 0 {
			return nil // error for EOF inside of string literal
		}

		if r == '"' {
			r = l.next()
			if r != '"' {
				l.unnext()
				t := token{src: l.src[l.start:l.pos]}
				t.typ = IDENT
				l.append(t)
				l.start = l.pos
				return blankState
			}
		}
	}
}

func lexOperator(l *sqlLex) stateFn {
	l.acceptRunFunc(isOperator)

	t := token{src: l.src[l.start:l.pos]}
	if len(t.src) == 1 {
		t.typ = int(t.src[0])
	} else {
		t.typ = OP
	}

	l.append(t)
	l.start = l.pos
	return blankState
}

func lexSimple(l *sqlLex) stateFn {
	l.append(token{int(l.src[l.start]), l.src[l.start:l.pos]})
	l.start = l.pos
	return blankState
}

func (l *sqlLex) skipWhitespace() {
	var r rune
	for r = l.next(); isWhitespace(r); r = l.next() {
	}

	if r != 0 {
		l.unnext()
	}

	l.ignore()
}

func isWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func isAlphanumeric(r rune) bool {
	return r == '_' || unicode.In(r, unicode.Letter, unicode.Digit)
}

func isOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/' || r == '=' || r == '<' || r == '>' || r == '!'
}
