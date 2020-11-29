package toml_test

import (
	"testing"

	sitter "github.com/franekjel/go-tree-sitter"
	"github.com/franekjel/go-tree-sitter/toml"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n := sitter.Parse([]byte(`key = "value"`), toml.GetLanguage())
	assert.Equal(
		"(document (pair (key) (string)))",
		n.String(),
	)
}
