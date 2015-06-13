package main

import (
	"testing"
)

// func TestSelect(t *testing.T) {
// 	tests := []struct {
// 		src string
// 		out string
// 	}{
// 		{"", ""},
// 		{"select 1", "select 1"},
// 		{"select a, b, c", `select
//   a,
//   b,
//   c`},
// 	}

// 	for i, tt := range tests {
// 		stmt, err := Parse(tt.src)
// 		NewFormatter
// 		lex := lexer{src: tt.src}
// 		lex.run()
// 		if !reflect.DeepEqual(tt.expected, lex.tokens) {
// 			t.Errorf("%d. Expected `%s`, got `%s`", tt.i, tt.in, lex.out)
// 		}
// 	}
// }
