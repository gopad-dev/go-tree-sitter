package gosum_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	sitter "github.com/smacker/go-tree-sitter"
	gosum "github.com/smacker/go-tree-sitter/go-sum"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d/go.mod h1:asat636LX7Bqt5lYEZ27JNDcqxfjdBQuJ/MM4CN/Lzo=\n"), gosum.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(checksum_database (checksum (module_path) (version (module_version) major: (module_version) (module_version) minor: (module_version) (module_version) patch: (module_version) pre_release: (number_with_decimal) build: (hex_number)) (checksum_value (hash_version) (hash))))",
		n.String(),
	)
}
