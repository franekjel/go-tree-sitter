package ocaml_test

import (
	"testing"

	sitter "github.com/franekjel/go-tree-sitter"
	"github.com/franekjel/go-tree-sitter/ocaml"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n := sitter.Parse([]byte(`print_endline "Hello World!"`), ocaml.GetLanguage())
	assert.Equal(
		"(compilation_unit (application_expression (value_path (value_name)) (string)))",
		n.String(),
	)
}
