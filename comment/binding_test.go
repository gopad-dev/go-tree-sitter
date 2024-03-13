package comment_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/comment"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("TODO: test"), comment.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(source (tag (name)))",
		n.String(),
	)
}
