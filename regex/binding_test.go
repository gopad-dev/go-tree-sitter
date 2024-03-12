package regex_test

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/regex"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	code := `[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`

	n, err := sitter.ParseCtx(context.Background(), []byte(code), regex.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(pattern (term (character_class (class_character) (class_range (class_character) (class_character)) (class_range (class_character) (class_character)) (class_range (class_character) (class_character)) (class_character) (class_character) (class_character) (class_character) (class_character) (identity_escape) (class_character) (class_character) (class_character)) (count_quantifier (decimal_digits) (decimal_digits)) (identity_escape) (character_class (class_character) (class_character) (class_character)) (count_quantifier (decimal_digits) (decimal_digits)) (boundary_assertion) (anonymous_capturing_group (pattern (term (character_class (class_character) (class_range (class_character) (class_character)) (class_range (class_character) (class_character)) (class_range (class_character) (class_character)) (class_character) (class_character) (class_character) (class_character) (identity_escape) (class_character) (class_character) (class_character) (class_character) (class_character) (class_character) (class_character) (class_character)) (zero_or_more))))))",
		n.String(),
	)
}
