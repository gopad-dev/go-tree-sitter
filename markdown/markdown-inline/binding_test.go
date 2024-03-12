package markdowninline_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	sitter "github.com/smacker/go-tree-sitter"
	markdowninline "github.com/smacker/go-tree-sitter/markdown/markdown-inline"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("`hi`lo`"), markdowninline.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(inline (code_span (code_span_delimiter) (code_span_delimiter)))",
		n.String(),
	)
}
