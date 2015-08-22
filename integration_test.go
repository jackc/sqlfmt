package sqlfmt_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/jackc/sqlfmt"
)

func TestIntegration(t *testing.T) {
	tests := []struct {
		inputFile          string
		expectedOutputFile string
	}{
		{inputFile: "float_constant.sql"},
		{inputFile: "simple_select_without_from.sql"},
		{inputFile: "simple_select_with_from.sql"},
		{inputFile: "select_from_aliased.sql"},
		{inputFile: "simple_select_with_selection_alias.sql"},
		{
			inputFile:          "simple_select_with_selection_alias_no_as.sql",
			expectedOutputFile: "simple_select_with_selection_alias.fmt.sql",
		},
		{inputFile: "select_table_dot_column.sql"},
		{inputFile: "simple_select_literal_integer.sql"},
		{inputFile: "simple_select_literal_text.sql"},
		{inputFile: "arithmetic_expression.sql"},
		{inputFile: "paren_expression.sql"},
		{inputFile: "subselect_expression.sql"},
		{inputFile: "comparison_expression.sql"},
		{inputFile: "select_from_comma_join.sql"},
		{inputFile: "select_from_cross_join.sql"},
		{inputFile: "select_from_natural_join.sql"},
		{inputFile: "select_from_join_using.sql"},
		{inputFile: "select_from_join_using_multiple.sql"},
		{inputFile: "select_from_join_on.sql"},
		{inputFile: "quoted_identifier.sql"},
		{inputFile: "boolean_binary_op.sql"},
		{inputFile: "boolean_not.sql"},
		{inputFile: "select_where.sql"},
		{inputFile: "order.sql"},
		{inputFile: "order_column_num.sql"},
		{inputFile: "order_desc.sql"},
		{inputFile: "order_multiple.sql"},
		{inputFile: "select_wrapped_by_parens.sql"},
		{inputFile: "order_nulls.sql"},
		{inputFile: "order_using.sql"},
		{inputFile: "distinct.sql"},
		{inputFile: "distinct_on.sql"},
		{inputFile: "group_by.sql"},
		{inputFile: "having.sql"},
		{inputFile: "limit.sql"},
		{
			inputFile:          "limit_fetch.sql",
			expectedOutputFile: "limit.fmt.sql",
		},
		{inputFile: "offset.sql"},
		{inputFile: "limit_offset.sql"},
		{
			inputFile:          "offset_limit.sql",
			expectedOutputFile: "limit_offset.fmt.sql",
		},
		{
			inputFile:          "offset_fetch.sql",
			expectedOutputFile: "limit_offset.fmt.sql",
		},
		{inputFile: "select_for_update.sql"},
		{inputFile: "select_for_no_key_update.sql"},
		{inputFile: "select_for_share.sql"},
		{inputFile: "select_for_key_share.sql"},
		{inputFile: "select_for_update_of.sql"},
		{inputFile: "select_for_update_nowait.sql"},
		{inputFile: "select_star.sql"},
		{inputFile: "select_table_dot_star.sql"},
		{inputFile: "table.sql"},
		{inputFile: "table_star.sql"},
		{inputFile: "table_only.sql"},
		{
			inputFile:          "table_only_paren.sql",
			expectedOutputFile: "table_only.fmt.sql",
		},
		{inputFile: "case_full.sql"},
		{inputFile: "case_implicit.sql"},
		{inputFile: "typecast.sql"},
		{inputFile: "const_type_name.sql"},
		{inputFile: "unary.sql"},
		{inputFile: "values.sql"},
		{inputFile: "values_order.sql"},
		{inputFile: "values_default.sql"},
		{inputFile: "collate.sql"},
		{inputFile: "collation_for.sql"},
		{inputFile: "at_time_zone.sql"},
		{inputFile: "like.sql"},
		{inputFile: "is_null.sql"},
		{inputFile: "is_bool_op.sql"},
		{inputFile: "is_distinct_from.sql"},
		{inputFile: "is_document.sql"},
		{inputFile: "union.sql"},
		{inputFile: "intersect.sql"},
		{inputFile: "except.sql"},
		{inputFile: "function_call_without_args.sql"},
		{inputFile: "function_call_with_star_arg.sql"},
		{inputFile: "function_call_with_positional_args.sql"},
		{inputFile: "function_call_with_pg_named_args.sql"},
		{inputFile: "function_call_with_sql_named_args.sql"},
		{inputFile: "function_call_with_order.sql"}, // TODO - fix formatting when order by is inside function call
		{inputFile: "function_call_with_all.sql"},
		{inputFile: "function_call_with_distinct.sql"},
		{inputFile: "function_call_with_filter.sql"},
		{inputFile: "null.sql"},
		{inputFile: "window_function.sql"},
		{inputFile: "window_function_partition_by.sql"},
		{inputFile: "window_function_order_by.sql"}, // TODO - fix formatting when order by is inside function call
		{inputFile: "window_function_named.sql"},
		{inputFile: "window_function_named_multiple.sql"},
		{inputFile: "window_function_frame.sql"},
		{inputFile: "exists.sql"},
		{inputFile: "array_constructor.sql"},
		{inputFile: "array_index.sql"},
		{inputFile: "array_slice.sql"},
		{inputFile: "array_typecast.sql"},
		{inputFile: "array_subselect.sql"},
		{inputFile: "cast_as.sql"},
		{inputFile: "func_expr_expr_list.sql"},
		{inputFile: "func_expr_no_parens.sql"},
		{inputFile: "func_expr_one_arg.sql"},
		{inputFile: "is_of_type_list.sql"},
		{inputFile: "between.sql"},
		{inputFile: "custom_operators.sql"},
		{inputFile: "postfix_operator.sql"},
		{inputFile: "b_expr.sql"}, // b_expr is duplicated subset of a_expr -- test its clauses
		{inputFile: "semicolon.sql"},
		{inputFile: "row.sql"},
		{inputFile: "overlaps.sql"},
		{inputFile: "interval.sql"},
		{inputFile: "in.sql"}, // TODO - fix formatting when spacing / new line is improved
		{inputFile: "subquery_op.sql"},
		{inputFile: "extract.sql"},
		{inputFile: "overlay.sql"},
		{inputFile: "position.sql"},
		{inputFile: "substring.sql"},
		{inputFile: "treat.sql"},
		{inputFile: "trim.sql"},
		{inputFile: "nullif.sql"},
		{inputFile: "xmlelement.sql"},
		{inputFile: "xmlexists.sql"},
		{inputFile: "xmlforest.sql"},
		{inputFile: "xmlparse.sql"},
		{inputFile: "xmlpi.sql"},
	}

	for i, tt := range tests {
		input, err := ioutil.ReadFile(path.Join("testdata", tt.inputFile))
		if err != nil {
			t.Errorf("%d. %v", i, err)
			continue
		}

		if tt.expectedOutputFile == "" {
			tt.expectedOutputFile = tt.inputFile[:len(tt.inputFile)-3] + "fmt.sql"
		}

		expected, err := ioutil.ReadFile(path.Join("testdata", tt.expectedOutputFile))
		if err != nil {
			t.Errorf("%d. %v", i, err)
			continue
		}

		lexer := sqlfmt.NewSqlLexer(string(input))
		stmt, err := sqlfmt.Parse(lexer)
		if err != nil {
			t.Errorf("%d. Given %s, %v", i, tt.inputFile, err)
			continue
		}

		var outBuf bytes.Buffer
		r := sqlfmt.NewTextRenderer(&outBuf)
		stmt.RenderTo(r)

		if outBuf.String() != string(expected) {
			actualFileName := path.Join("tmp", fmt.Sprintf("%d.sql", i))
			err = ioutil.WriteFile(actualFileName, outBuf.Bytes(), os.ModePerm)
			if err != nil {
				t.Fatal(err)
			}

			t.Errorf("%d. Given %s, did not receive %s. Unexpected output written to %s", i, tt.inputFile, tt.expectedOutputFile, actualFileName)
		}
	}
}
