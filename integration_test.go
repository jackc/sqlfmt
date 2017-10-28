package sqlfmt_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackc/sqlfmt"
)

func TestIntegration(t *testing.T) {
	// TODO - put these notes in testdata input files when sqlfmt handles comments
	//	{inputFile: "b_expr.sql"}, // b_expr is duplicated subset of a_expr -- test its clauses
	//	{inputFile: "in.sql"}, // TODO - fix formatting when spacing / new line is improved

	fileInfos, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}

	for _, fi := range fileInfos {
		if fi.Name()[len(fi.Name())-10:] != ".input.sql" {
			continue
		}

		testName := fi.Name()[:len(fi.Name())-10]
		inputPath := filepath.Join("testdata", fi.Name())
		goldenPath := filepath.Join("testdata", testName+".golden.sql")

		input, err := ioutil.ReadFile(inputPath)
		if err != nil {
			t.Errorf("%s: %v", testName, err)
			continue
		}

		expected, err := ioutil.ReadFile(goldenPath)
		if err != nil {
			t.Errorf("%s: %v", testName, err)
			continue
		}

		lexer := sqlfmt.NewSqlLexer(string(input))
		stmt, err := sqlfmt.Parse(lexer)
		if err != nil {
			t.Errorf("%s: Given %s, %v", testName, inputPath, err)
			continue
		}

		var outBuf bytes.Buffer
		r := sqlfmt.NewTextRenderer(&outBuf)
		stmt.RenderTo(r)

		if outBuf.String() != string(expected) {
			actualFileName := filepath.Join("tmp", fmt.Sprintf("%s.sql", testName))
			err = ioutil.WriteFile(actualFileName, outBuf.Bytes(), os.ModePerm)
			if err != nil {
				t.Fatal(err)
			}

			t.Errorf("%s: Unexpected output written to %s", testName, actualFileName)
		}
	}
}
