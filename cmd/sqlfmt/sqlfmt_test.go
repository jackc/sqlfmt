package main_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
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

func sqlfmt(sql []byte, args ...string) ([]byte, error) {
	cmd := exec.Command("tmp/sqlfmt", args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("cmd.StdinPipe failed: %v", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("cmd.StdoutPipe failed: %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("cmd.StderrPipe failed: %v", err)
	}

	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("cmd.Start failed: %v", err)
	}

	_, err = stdin.Write(sql)
	if err != nil {
		return nil, fmt.Errorf("stdin.Write failed: %v", err)
	}

	err = stdin.Close()
	if err != nil {
		return nil, fmt.Errorf("stdin.Close failed: %v", err)
	}

	output, err := ioutil.ReadAll(stdout)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll(stdout) failed: %v", err)
	}

	errout, err := ioutil.ReadAll(stderr)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll(stderr) failed: %v", err)
	}

	err = cmd.Wait()
	if err != nil {
		return nil, fmt.Errorf("cmd.Wait failed: %v\n%s", err, string(errout))
	}

	return output, nil
}

func TestFileInput(t *testing.T) {
	inputFile := "simple_select_without_from.sql"
	expectedOutputFile := "simple_select_without_from.fmt.sql"

	expected, err := ioutil.ReadFile(path.Join("../../testdata", expectedOutputFile))
	if err != nil {
		t.Fatal(err)
	}

	filePath := path.Join("../../testdata", inputFile)
	output, err := sqlfmt(nil, filePath)
	if err != nil {
		t.Fatalf("sqlfmt failed with %s: %v", filePath, err)
	}

	if bytes.Compare(output, expected) != 0 {
		actualFileName := path.Join("tmp", "TestFileInput.sql")
		err = ioutil.WriteFile(actualFileName, output, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		t.Errorf("Given %s, did not receive %s. Unexpected output written to %s", filePath, expectedOutputFile, actualFileName)
	}
}

func TestSqlFmt(t *testing.T) {
	tests := []struct {
		inputFile          string
		expectedOutputFile string
	}{
		{
			inputFile:          "simple_select_without_from.sql",
			expectedOutputFile: "simple_select_without_from.fmt.sql",
		},
		{
			inputFile:          "simple_select_with_from.sql",
			expectedOutputFile: "simple_select_with_from.fmt.sql",
		},
		{
			inputFile:          "select_from_aliased.sql",
			expectedOutputFile: "select_from_aliased.fmt.sql",
		},
	}

	for i, tt := range tests {
		input, err := ioutil.ReadFile(path.Join("../../testdata", tt.inputFile))
		if err != nil {
			t.Errorf("%d. %v", i, err)
			continue
		}

		expected, err := ioutil.ReadFile(path.Join("../../testdata", tt.expectedOutputFile))
		if err != nil {
			t.Errorf("%d. %v", i, err)
			continue
		}

		output, err := sqlfmt(input)
		if err != nil {
			t.Errorf("%d. sqlfmt failed with %s: %v", i, tt.inputFile, err)
			continue
		}

		if bytes.Compare(output, expected) != 0 {
			actualFileName := path.Join("tmp", fmt.Sprintf("%d.sql", i))
			err = ioutil.WriteFile(actualFileName, output, os.ModePerm)
			if err != nil {
				t.Fatal(err)
			}

			t.Errorf("%d. Given %s, did not receive %s. Unexpected output written to %s", i, tt.inputFile, tt.expectedOutputFile, actualFileName)
		}
	}
}
