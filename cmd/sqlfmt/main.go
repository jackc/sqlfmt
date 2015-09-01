package main

import (
	"flag"
	"fmt"
	"github.com/jackc/sqlfmt"
	"io"
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

type job struct {
	name string
	r    io.ReadCloser
	w    io.WriteCloser
}

func (j *job) run() error {
	if j.r == nil {
		var err error
		j.r, err = os.Open(j.name)
		if err != nil {
			return err
		}
	}

	input, err := ioutil.ReadAll(j.r)
	if err != nil {
		return err
	}

	err = j.r.Close()
	if err != nil {
		return err
	}

	lexer := sqlfmt.NewSqlLexer(string(input))
	stmt, err := sqlfmt.Parse(lexer)
	if err != nil {
		return err
	}

	var inPlace bool
	var tmpPath string

	if j.w == nil {
		dir := filepath.Dir(j.name)
		base := filepath.Base(j.name)
		tmpPath = path.Join(dir, "."+base+".sqlfmt")
		j.w, err = os.Create(tmpPath)
		if err != nil {
			return err
		}
		inPlace = true
	}

	r := sqlfmt.NewTextRenderer(j.w)
	stmt.RenderTo(r)
	if r.Error() != nil {
		return err
	}

	if inPlace {
		err = j.w.Close()
		if err != nil {
			return err
		}
		err = os.Rename(tmpPath, j.name)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage:  %s [options] [path ...]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&options.write, "w", false, "write result to (source) file instead of stdout")
	flag.BoolVar(&options.version, "version", false, "print version and exit")
	flag.Parse()

	if options.version {
		fmt.Printf("sqlfmt v%v\n", Version)
		os.Exit(0)
	}

	var jobs []job

	if len(flag.Args()) > 0 {
		for _, fp := range flag.Args() {
			j := job{name: fp}
			if !options.write {
				j.w = os.Stdout
			}
			jobs = append(jobs, j)
		}
	} else {
		jobs = append(jobs, job{r: os.Stdin, w: os.Stdout})
	}

	var errors []error

	for _, j := range jobs {
		if err := j.run(); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		for _, e := range errors {
			fmt.Fprintln(os.Stderr, e)
		}
		os.Exit(1)
	}
}
