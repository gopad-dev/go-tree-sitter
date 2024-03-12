package gitattributes_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/gitattributes"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("/src/** linguist-generated=true"), gitattributes.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(document (attributes_line (pattern relative_flag: (directory_separator) (pattern_char) (pattern_char) (pattern_char) relative_flag: (directory_separator) (wildcard_chars_allow_slash)) (attribute_set_to (attribute) (value))) (MISSING \"_line_token2\"))",
		n.String(),
	)
}
