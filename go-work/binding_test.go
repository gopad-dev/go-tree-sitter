package gowork_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	sitter "github.com/smacker/go-tree-sitter"
	gowork "github.com/smacker/go-tree-sitter/go-work"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("go 1.22\n\nuse .\n"), gowork.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(source_file (go_directive (go_version)) (use_directive (use_spec (file_path))))",
		n.String(),
	)
}
