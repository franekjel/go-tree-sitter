package golang_test

import (
	"testing"

	sitter "github.com/franekjel/go-tree-sitter"
	"github.com/franekjel/go-tree-sitter/golang"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n := sitter.Parse([]byte("package main"), golang.GetLanguage())
	assert.Equal(
		"(source_file (package_clause (package_identifier)))",
		n.String(),
	)
}
