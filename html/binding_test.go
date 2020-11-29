package html_test

import (
	"testing"

	sitter "github.com/franekjel/go-tree-sitter"
	"github.com/franekjel/go-tree-sitter/html"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n := sitter.Parse([]byte(`<HTML><BODY><P>Hello World</P></BODY></HTML>`), html.GetLanguage())
	assert.Equal(
		"(fragment (element (start_tag (tag_name)) (element (start_tag (tag_name)) (element (start_tag (tag_name)) (text) (end_tag (tag_name))) (end_tag (tag_name))) (end_tag (tag_name))))",
		n.String(),
	)
}
