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
	inputFile := "simple_select_without_from.input.sql"
	expectedOutputFile := "simple_select_without_from.golden.sql"

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

func TestFileFormatInPlace(t *testing.T) {
	inputFile := "simple_select_without_from.input.sql"
	expectedOutputFile := "simple_select_without_from.golden.sql"
	expectedOutputPath := path.Join("../../testdata", expectedOutputFile)

	expected, err := ioutil.ReadFile(expectedOutputPath)
	if err != nil {
		t.Fatal(err)
	}

	sourcePath := path.Join("../../testdata", inputFile)
	source, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		t.Fatal(err)
	}

	tmpFilePath := path.Join("tmp", inputFile)
	err = ioutil.WriteFile(tmpFilePath, source, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	output, err := sqlfmt(nil, "-w", tmpFilePath)
	if err != nil {
		t.Fatalf("sqlfmt failed with %s: %v", tmpFilePath, err)
	}

	if len(output) != 0 {
		t.Fatal("Expected output to be empty, but it wasn't")
	}

	output, err = ioutil.ReadFile(tmpFilePath)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(output, expected) != 0 {
		t.Errorf("Given %s, did not receive %s. Unexpected output written to %s", sourcePath, expectedOutputPath, tmpFilePath)
	}
}

func TestMultipleFileFormatInPlace(t *testing.T) {
	tests := []struct {
		Name        string
		Source      []byte
		Expected    []byte
		TmpFilePath string
	}{
		{Name: "between"},
		{Name: "bitconst"},
	}

	var err error
	args := []string{"-w"}
	for i, _ := range tests {
		sourcePath := path.Join("../../testdata", tests[i].Name+".input.sql")
		tests[i].Source, err = ioutil.ReadFile(sourcePath)
		if err != nil {
			t.Fatal(err)
		}

		expectedOutputPath := path.Join("../../testdata", tests[i].Name+".golden.sql")
		tests[i].Expected, err = ioutil.ReadFile(expectedOutputPath)
		if err != nil {
			t.Fatal(err)
		}

		tests[i].TmpFilePath = path.Join("tmp", tests[i].Name+".sql")
		err = ioutil.WriteFile(tests[i].TmpFilePath, tests[i].Source, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		args = append(args, tests[i].TmpFilePath)
	}

	output, err := sqlfmt(nil, args...)
	if err != nil {
		t.Fatalf("sqlfmt failed with %s: %v", args, err)
	}

	if len(output) != 0 {
		t.Fatal("Expected output to be empty, but it wasn't")
	}

	for _, tt := range tests {
		output, err = ioutil.ReadFile(tt.TmpFilePath)
		if err != nil {
			t.Fatal(err)
		}

		if bytes.Compare(output, tt.Expected) != 0 {
			t.Errorf("%s: unexpected output written to %s", tt.Name, tt.TmpFilePath)
		}
	}
}

func TestSqlFmtAll(t *testing.T) {
	fileInfos, err := ioutil.ReadDir("../../testdata")
	if err != nil {
		t.Fatal(err)
	}

	for _, fi := range fileInfos {
		if fi.Name()[len(fi.Name())-10:] != ".input.sql" {
			continue
		}

		testName := fi.Name()[:len(fi.Name())-10]
		inputPath := path.Join("../../testdata", fi.Name())
		goldenPath := path.Join("../../testdata", testName+".golden.sql")

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

		output, err := sqlfmt(input)
		if err != nil {
			t.Errorf("%s: sqlfmt failed with: %v", testName, err)
			continue
		}

		if bytes.Compare(output, expected) != 0 {
			actualFileName := path.Join("tmp", fmt.Sprintf("%s.sql", testName))
			err = ioutil.WriteFile(actualFileName, output, os.ModePerm)
			if err != nil {
				t.Fatal(err)
			}

			t.Errorf("%s: Unexpected output written to %s", testName, actualFileName)
		}
	}
}
