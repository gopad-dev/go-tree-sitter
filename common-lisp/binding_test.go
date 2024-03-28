package common_lisp_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/common-lisp"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(";; Addiere 2 und 2\n(+ 2 2)"), common_lisp.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(source (comment) (list_lit value: (sym_lit) value: (num_lit) value: (num_lit)))",
		n.String(),
	)
}
