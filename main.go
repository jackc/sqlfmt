package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, world")
	io.Copy(os.Stdout, os.Stdin)
}

type SelectStmt struct {
  Fields    []string
  FromTable string
}

type PrettyPrinter interface {
  io.Writer
  Indent(string)
  Unindent()
  Newline() error
}

func (s *SelectStmt)
