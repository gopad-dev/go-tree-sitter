package gomod_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	sitter "github.com/smacker/go-tree-sitter"
	gomod "github.com/smacker/go-tree-sitter/go-mod"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("module test"), gomod.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(source_file (module_directive (module_path)))",
		n.String(),
	)
}
