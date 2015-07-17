package main_test

import (
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

func sqlfmt(sql string, args ...string) (string, error) {
	cmd := exec.Command("tmp/sqlfmt", args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", fmt.Errorf("cmd.StdinPipe failed: %v", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("cmd.StdoutPipe failed: %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", fmt.Errorf("cmd.StderrPipe failed: %v", err)
	}

	err = cmd.Start()
	if err != nil {
		return "", fmt.Errorf("cmd.Start failed: %v", err)
	}

	_, err = fmt.Fprint(stdin, sql)
	if err != nil {
		return "", fmt.Errorf("fmt.Fprint failed: %v", err)
	}

	err = stdin.Close()
	if err != nil {
		return "", fmt.Errorf("stdin.Close failed: %v", err)
	}

	output, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll(stdout) failed: %v", err)
	}

	errout, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll(stderr) failed: %v", err)
	}

	err = cmd.Wait()
	if err != nil {
		return "", fmt.Errorf("cmd.Wait failed: %v\n%s", err, string(errout))
	}

	return string(output), nil
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
		{
			inputFile:          "simple_select_with_selection_alias.sql",
			expectedOutputFile: "simple_select_with_selection_alias.fmt.sql",
		},
		{
			inputFile:          "simple_select_with_selection_alias_no_as.sql",
			expectedOutputFile: "simple_select_with_selection_alias.fmt.sql",
		},
		{
			inputFile:          "select_table_dot_column.sql",
			expectedOutputFile: "select_table_dot_column.fmt.sql",
		},
		{
			inputFile:          "simple_select_literal_integer.sql",
			expectedOutputFile: "simple_select_literal_integer.fmt.sql",
		},
		{
			inputFile:          "simple_select_literal_text.sql",
			expectedOutputFile: "simple_select_literal_text.fmt.sql",
		},
		{
			inputFile:          "arithmetic_expression.sql",
			expectedOutputFile: "arithmetic_expression.fmt.sql",
		},
		{
			inputFile:          "paren_expression.sql",
			expectedOutputFile: "paren_expression.fmt.sql",
		},
		{
			inputFile:          "subselect_expression.sql",
			expectedOutputFile: "subselect_expression.fmt.sql",
		},
		{
			inputFile:          "comparison_expression.sql",
			expectedOutputFile: "comparison_expression.fmt.sql",
		},
		{
			inputFile:          "select_from_comma_join.sql",
			expectedOutputFile: "select_from_comma_join.fmt.sql",
		},
		{
			inputFile:          "select_from_cross_join.sql",
			expectedOutputFile: "select_from_cross_join.fmt.sql",
		},
		{
			inputFile:          "select_from_natural_join.sql",
			expectedOutputFile: "select_from_natural_join.fmt.sql",
		},
		{
			inputFile:          "select_from_join_using.sql",
			expectedOutputFile: "select_from_join_using.fmt.sql",
		},
		{
			inputFile:          "select_from_join_using_multiple.sql",
			expectedOutputFile: "select_from_join_using_multiple.fmt.sql",
		},
		{
			inputFile:          "select_from_join_on.sql",
			expectedOutputFile: "select_from_join_on.fmt.sql",
		},
		{
			inputFile:          "quoted_identifier.sql",
			expectedOutputFile: "quoted_identifier.fmt.sql",
		},
		{
			inputFile:          "boolean_binary_op.sql",
			expectedOutputFile: "boolean_binary_op.fmt.sql",
		},
		{
			inputFile:          "boolean_not.sql",
			expectedOutputFile: "boolean_not.fmt.sql",
		},
		{
			inputFile:          "select_where.sql",
			expectedOutputFile: "select_where.fmt.sql",
		},
		{
			inputFile:          "order.sql",
			expectedOutputFile: "order.fmt.sql",
		},
		{
			inputFile:          "order_column_num.sql",
			expectedOutputFile: "order_column_num.fmt.sql",
		},
		{
			inputFile:          "order_desc.sql",
			expectedOutputFile: "order_desc.fmt.sql",
		},
		{
			inputFile:          "order_multiple.sql",
			expectedOutputFile: "order_multiple.fmt.sql",
		},
		{
			inputFile:          "select_wrapped_by_parens.sql",
			expectedOutputFile: "select_wrapped_by_parens.fmt.sql",
		},
		{
			inputFile:          "order_nulls.sql",
			expectedOutputFile: "order_nulls.fmt.sql",
		},
		{
			inputFile:          "order_using.sql",
			expectedOutputFile: "order_using.fmt.sql",
		},
		{
			inputFile:          "distinct.sql",
			expectedOutputFile: "distinct.fmt.sql",
		},
		{
			inputFile:          "distinct_on.sql",
			expectedOutputFile: "distinct_on.fmt.sql",
		},
		{
			inputFile:          "group_by.sql",
			expectedOutputFile: "group_by.fmt.sql",
		},
		{
			inputFile:          "having.sql",
			expectedOutputFile: "having.fmt.sql",
		},
		{
			inputFile:          "limit.sql",
			expectedOutputFile: "limit.fmt.sql",
		},
		{
			inputFile:          "limit_fetch.sql",
			expectedOutputFile: "limit.fmt.sql",
		},
		{
			inputFile:          "offset.sql",
			expectedOutputFile: "offset.fmt.sql",
		},
		{
			inputFile:          "limit_offset.sql",
			expectedOutputFile: "limit_offset.fmt.sql",
		},
		{
			inputFile:          "offset_limit.sql",
			expectedOutputFile: "limit_offset.fmt.sql",
		},
		{
			inputFile:          "offset_fetch.sql",
			expectedOutputFile: "limit_offset.fmt.sql",
		},
		{
			inputFile:          "select_for_update.sql",
			expectedOutputFile: "select_for_update.fmt.sql",
		},
		{
			inputFile:          "select_for_no_key_update.sql",
			expectedOutputFile: "select_for_no_key_update.fmt.sql",
		},
		{
			inputFile:          "select_for_share.sql",
			expectedOutputFile: "select_for_share.fmt.sql",
		},
		{
			inputFile:          "select_for_key_share.sql",
			expectedOutputFile: "select_for_key_share.fmt.sql",
		},
		{
			inputFile:          "select_for_update_of.sql",
			expectedOutputFile: "select_for_update_of.fmt.sql",
		},
		{
			inputFile:          "select_for_update_nowait.sql",
			expectedOutputFile: "select_for_update_nowait.fmt.sql",
		},
		{
			inputFile:          "select_star.sql",
			expectedOutputFile: "select_star.fmt.sql",
		},
		{
			inputFile:          "select_table_dot_star.sql",
			expectedOutputFile: "select_table_dot_star.fmt.sql",
		},
		{
			inputFile:          "case_full.sql",
			expectedOutputFile: "case_full.fmt.sql",
		},
		{
			inputFile:          "case_implicit.sql",
			expectedOutputFile: "case_implicit.fmt.sql",
		},
		{
			inputFile:          "typecast.sql",
			expectedOutputFile: "typecast.fmt.sql",
		},
		{
			inputFile:          "unary.sql",
			expectedOutputFile: "unary.fmt.sql",
		},
		{
			inputFile:          "values.sql",
			expectedOutputFile: "values.fmt.sql",
		},
		{
			inputFile:          "values_default.sql",
			expectedOutputFile: "values_default.fmt.sql",
		},
		{
			inputFile:          "collate.sql",
			expectedOutputFile: "collate.fmt.sql",
		},
		{
			inputFile:          "at_time_zone.sql",
			expectedOutputFile: "at_time_zone.fmt.sql",
		},
		{
			inputFile:          "like.sql",
			expectedOutputFile: "like.fmt.sql",
		},
		{
			inputFile:          "is_null.sql",
			expectedOutputFile: "is_null.fmt.sql",
		},
	}

	for i, tt := range tests {
		input, err := ioutil.ReadFile(path.Join("testdata", tt.inputFile))
		if err != nil {
			t.Errorf("%d. %v", i, err)
			continue
		}

		expected, err := ioutil.ReadFile(path.Join("testdata", tt.expectedOutputFile))
		if err != nil {
			t.Errorf("%d. %v", i, err)
			continue
		}

		output, err := sqlfmt(string(input))
		if err != nil {
			t.Errorf("%d. sqlfmt failed with %s: %v", i, tt.inputFile, err)
			continue
		}

		if output != string(expected) {
			actualFileName := path.Join("tmp", fmt.Sprintf("%d.sql", i))
			err = ioutil.WriteFile(actualFileName, []byte(output), os.ModePerm)
			if err != nil {
				t.Fatal(err)
			}

			t.Errorf("%d. Given %s, did not receive %s. Unexpected output written to %s", i, tt.inputFile, tt.expectedOutputFile, actualFileName)
		}
	}
}
