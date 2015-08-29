package main

import (
	"flag"
	"fmt"
	"github.com/jackc/sqlfmt"
	"io/ioutil"
	"os"
)

const Version = "0.1.0"

var options struct {
	version bool
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage:  %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&options.version, "version", false, "print version and exit")
	flag.Parse()

	if options.version {
		fmt.Printf("sqlfmt v%v\n", Version)
		os.Exit(0)
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}
	lexer := sqlfmt.NewSqlLexer(string(input))
	stmt, err := sqlfmt.Parse(lexer)
	if err != nil {
		os.Exit(1)
	}

	r := sqlfmt.NewTextRenderer(os.Stdout)
	stmt.RenderTo(r)
}
