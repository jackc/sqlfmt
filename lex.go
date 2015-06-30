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
	tokens chan token
	stmt   *SelectStmt
}

func (x *sqlLex) Lex(yylval *sqlSymType) int {
	for {
		select {
		case token := <-x.tokens:
			yylval.src = token.src
			return token.typ
		default:
			x.state = x.state(x)
			if x.state == nil {
				return eof
			}
		}
	}
}

// The parser calls this method on a parse error.
func (x *sqlLex) Error(s string) {
	log.Printf("parse error: %s", s)
}

func NewSqlLexer(src string) *sqlLex {
	return &sqlLex{src: src,
		tokens: make(chan token, 2),
		state:  blankState,
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
	case r == ',':
		return lexComma
	case isWhitespace(r):
		l.skipWhitespace()
		return blankState
	case isAlphanumeric(r):
		return lexAlphanumeric
	}
	return nil
}

func lexAlphanumeric(l *sqlLex) stateFn {
	l.acceptRunFunc(isAlphanumeric)

	t := token{src: l.src[l.start:l.pos]}

	switch {
	case strings.EqualFold(t.src, "select"):
		t.typ = SELECT
	case strings.EqualFold(t.src, "as"):
		t.typ = AS
	case strings.EqualFold(t.src, "from"):
		t.typ = FROM
	default:
		t.typ = IDENTIFIER
	}

	l.tokens <- t
	l.start = l.pos
	return blankState
}

func lexComma(l *sqlLex) stateFn {
	l.tokens <- token{COMMA, l.src[l.start:l.pos]}
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
