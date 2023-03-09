package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sqlToGoType(t *testing.T) {
	testCases := []struct {
		in  string
		exp string
	}{
		{"text", "string"},
		{"int", "int"},
		{"float", "float64"},
		{"varchar", "string"},
	}
	for _, tc := range testCases {
		out := sqlToGoType(tc.in)
		assert.Equalf(
			t,
			tc.exp,
			out,
			"sqlToGoType(%q) = %q, expected %q", tc.in, out, tc.exp,
		)
	}
}

func Test_goToSqlType(t *testing.T) {
	testCases := []struct {
		in  string
		exp string
	}{
		{"string", "VARCHAR(255)"},
		{"uint8", "INTEGER"},
		{"int64", "BIGINT"},
		{"float64", "FLOAT"},
		{"time.Time", "TIMESTAMP"},
	}
	for _, tc := range testCases {
		out := goToSqlType(tc.in)
		assert.Equalf(
			t,
			tc.exp,
			out,
			"goToSqlType(%q) = %q, expected %q", tc.in, out, tc.exp,
		)
	}
	// goToSqlType(goType string)
	// switch goType {
	// case "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32":
	// return "INTEGER"
	// case "int64", "uint64":
	// return "BIGINT"
	// case "float32", "float64":
	// return "FLOAT"
	// case "string":
	// return "VARCHAR(255)"
	// case "time.Time":
	// return "TIMESTAMP"
	// default:
	// return ""
	// }
}

func Test_pluralize(t *testing.T) {
	testCases := []struct {
		in  string
		exp string
	}{
		{"world", "worlds"},
		{"Pass", "Passes"},
		{"key", "keys"},
		{"ruby", "rubies"},
		{"flash", "flashes"},
		{"REACH", "REACHES"},
	}
	for _, tc := range testCases {
		out := pluralize(tc.in)
		assert.Equalf(
			t,
			tc.exp,
			out,
			"pluralize(%q) = %q, expected %q", tc.in, out, tc.exp,
		)
	}
}

func Test_singularize(t *testing.T) {
	testCases := []struct {
		in  string
		exp string
	}{
		{"worlds", "world"},
		{"Passes", "Pass"},
		{"keys", "key"},
		{"rubies", "ruby"},
		{"flashes", "flash"},
		{"REACHES", "REACH"},
	}

	for _, tc := range testCases {
		out := singularize(tc.in)
		assert.Equalf(
			t,
			tc.exp,
			out,
			"singularize(%q) = %q, expected %q", tc.in, out, tc.exp,
		)
	}
}

func Test_snakeToCamel(t *testing.T) {
	testCases := []struct {
		in  string
		exp string
	}{
		{"hello_world", "HelloWorld"},
		{"my_var", "MyVar"},
		{"user_id", "UserID"},
		{"api_key", "APIKey"},
		{"snake_case_example", "SnakeCaseExample"},
		{"CamelCaseExample", "CamelCaseExample"},
	}

	for _, tc := range testCases {
		out := snakeToCamel(tc.in)
		assert.Equalf(
			t,
			tc.exp,
			out,
			"camelToSnake(%q) = %q, expected %q", tc.in, out, tc.exp,
		)
	}
}

func Test_camelToSnake(t *testing.T) {
	testCases := []struct {
		in  string
		exp string
	}{
		{"HelloWorld", "hello_world"},
		{"MyVar", "my_var"},
		{"UserID", "user_id"},
		{"APIKey", "api_key"},
		{"CamelCaseExample", "camel_case_example"},
		{"snake_case_example", "snake_case_example"},
	}

	for _, tc := range testCases {
		out := camelToSnake(tc.in)
		assert.Equalf(
			t,
			tc.exp,
			out,
			"camelToSnake(%q) = %q, expected %q", tc.in, out, tc.exp,
		)
	}
}
