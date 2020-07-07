package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCamelCase(t *testing.T) {
	cases := [][]string{
		{"abc", "Abc"},
		{"abc_def", "AbcDef"},
		{"abc_def_g", "AbcDefG"},
		{"abc_def_gh", "AbcDefGh"},
		{"abc_def_g_hi", "AbcDefGHi"},
		{"abc_def_gh_ij", "AbcDefGhIj"},
	}
	for _, c := range cases {
		assert.Equal(t, c[1], ToCamelCase(c[0]))
	}
}

func TestToSnakeCase(t *testing.T) {
	cases := [][]string{
		{"abc", "Abc"},
		{"abc", "abc"},
		{"abc_def", "AbcDef"},
		{"abc_def_g", "AbcDefG"},
		{"abc_def_gh", "AbcDefGH"},
		{"abc_def_g_hi", "AbcDefGHi"},
		{"abc_def_gh_ij", "AbcDefGHIj"},
	}
	for _, c := range cases {
		assert.Equal(t, c[0], ToSnakeCase(c[1]))
	}
}

func TestUpperInitial(t *testing.T) {
	cases := [][]string{
		{"Abc", "abc"},
	}
	for _, c := range cases {
		assert.Equal(t, c[0], UpperInitial(c[1]))
	}
}
