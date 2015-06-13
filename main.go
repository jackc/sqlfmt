package main

import (
	"bytes"
	"fmt"
)

func main() {
	var ss SelectStmt
	ss.Fields = append(ss.Fields, "foo")
	ss.Fields = append(ss.Fields, "bar")
	ss.FromTable = "baz"

	fmt.Println(ss)
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

	fmt.Fprintln(&buf, "from", s.FromTable)

	return buf.String()
}
