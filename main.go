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

type ColumnRef struct {
	Table  string
	Column string
}

type SelectExpr struct {
	Expr  ColumnRef
	Alias string
}

type SelectStmt struct {
	Fields    []SelectExpr
	FromTable string
}

func (s SelectStmt) String() string {
	var buf bytes.Buffer

	fmt.Fprintln(&buf, "select")

	for i, f := range s.Fields {
		if f.Expr.Table != "" {
			fmt.Fprintf(&buf, "  %s.%s", f.Expr.Table, f.Expr.Column)
		} else {
			fmt.Fprintf(&buf, "  %s", f.Expr.Column)
		}
		if f.Alias != "" {
			fmt.Fprintf(&buf, " as %s", f.Alias)
		}
		if i < len(s.Fields)-1 {
			fmt.Fprint(&buf, ",")
		}
		fmt.Fprint(&buf, "\n")
	}

	if s.FromTable != "" {
		fmt.Fprintln(&buf, "from")
		fmt.Fprintf(&buf, "  %s\n", s.FromTable)
	}

	return buf.String()
}
