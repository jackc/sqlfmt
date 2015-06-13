//go:generate -command yacc go tool yacc
//go:generate yacc -o sql.go -p "sql" sql.y
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}
	lexer := NewSqlLexer(string(input))
	sqlParse(lexer)
	fmt.Print(lexer.stmt)
}

type SelectStmt struct {
	Fields    []string
	FromTable string
}

func (s SelectStmt) String() string {
	var buf bytes.Buffer

	fmt.Fprintln(&buf, "select")

	for i, f := range s.Fields {
		fmt.Fprintf(&buf, "  %s", f)
		if i < len(s.Fields)-1 {
			fmt.Fprint(&buf, ",")
		}
		fmt.Fprint(&buf, "\n")
	}

	if s.FromTable != "" {
		fmt.Fprintln(&buf, "from", s.FromTable)
	}

	return buf.String()
}
