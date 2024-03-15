package scala_test

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/scala"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`package com.foo.bar`), scala.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		`(compilation_unit (package_clause name: (package_identifier (identifier) (identifier) (identifier))))`,
		n.String(),
	)
}
