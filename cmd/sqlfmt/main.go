package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/jackc/sqlfmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

const Version = "0.1.0"

var options struct {
	write   bool
	version bool
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage:  %s [options] [path]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&options.write, "w", false, "write result to (source) file instead of stdout")
	flag.BoolVar(&options.version, "version", false, "print version and exit")
	flag.Parse()

	if options.version {
		fmt.Printf("sqlfmt v%v\n", Version)
		os.Exit(0)
	}

	var input []byte
	var err error

	if len(flag.Args()) == 0 {
		input, err = ioutil.ReadAll(os.Stdin)
	} else if len(flag.Args()) == 1 {
		input, err = ioutil.ReadFile(flag.Arg(0))
	} else {
		flag.Usage()
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	lexer := sqlfmt.NewSqlLexer(string(input))
	stmt, err := sqlfmt.Parse(lexer)
	if err != nil {
		os.Exit(1)
	}

	buf := &bytes.Buffer{}
	r := sqlfmt.NewTextRenderer(buf)
	stmt.RenderTo(r)

	if options.write {
		dir := filepath.Dir(flag.Arg(0))
		base := filepath.Base(flag.Arg(0))
		tmpPath := path.Join(dir, "."+base+".sqlfmt")
		err = ioutil.WriteFile(tmpPath, buf.Bytes(), os.ModePerm)
		if err != nil {
			os.Exit(1)
		}
		err = os.Rename(tmpPath, flag.Arg(0))
		if err != nil {
			os.Exit(1)
		}
	} else {
		buf.WriteTo(os.Stdout)
	}
}
