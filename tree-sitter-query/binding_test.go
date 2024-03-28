package tree_sitter_query_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	sitter "github.com/smacker/go-tree-sitter"
	tree_sitter_query "github.com/smacker/go-tree-sitter/tree-sitter-query"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`(node _ @wildcard)
;       ^ @punctuation.special
;        ^ @type`), tree_sitter_query.GetLanguage())

	assert.NoError(err)
	assert.Equal(
		"(program (named_node name: (identifier) (anonymous_node (capture name: (identifier)))) (comment) (comment))",
		n.String(),
	)
}
