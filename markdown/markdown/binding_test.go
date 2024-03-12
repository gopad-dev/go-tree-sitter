package markdown_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/markdown/markdown"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`foo	baz	bim`), markdown.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(document (section (paragraph (inline)))",
		n.String(),
	)
}
