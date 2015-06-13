package main_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	os.MkdirAll("tmp", os.ModeDir|os.ModePerm)
	err := exec.Command("go", "build", "-o", "tmp/sqlfmt").Run()
	if err != nil {
		fmt.Println("Failed to build sqlfmt binary:", err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func sqlfmt(t *testing.T, sql string, args ...string) string {
	cmd := exec.Command("tmp/sqlfmt", args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("cmd.StdinPipe failed: %v", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatalf("cmd.StdoutPipe failed: %v", err)
	}

	err = cmd.Start()
	if err != nil {
		t.Fatalf("cmd.Start failed: %v", err)
	}

	_, err = fmt.Fprint(stdin, sql)
	if err != nil {
		t.Fatalf("fmt.Fprint failed: %v", err)
	}

	err = stdin.Close()
	if err != nil {
		t.Fatalf("stdin.Close failed: %v", err)
	}

	output, err := ioutil.ReadAll(stdout)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(stdout) failed: %v", err)
	}

	err = cmd.Wait()
	if err != nil {
		t.Fatalf("cmd.Wait failed: %v", err)
	}

	return string(output)
}

func TestSqlFmt(t *testing.T) {
	output := sqlfmt(t, "select foo, bar")
	if output != `select
  foo,
  bar
` {
		t.Errorf("Unexpected output: %v", output)
	}
}
