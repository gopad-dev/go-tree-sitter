package gitignore_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/gitignore"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(".idea/"), gitignore.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(document (ERROR (pattern (pattern_char) (pattern_char) (pattern_char) (pattern_char) (pattern_char) directory_flag: (directory_separator))))",
		n.String(),
	)
}
