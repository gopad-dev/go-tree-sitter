package json_test

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/json"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	code := `{"test":"lol","test2":1,"test3":1.0,"test4":true,"test5":null,"test6":[1,2,3],"test7":{"test8":1}}`

	n, err := sitter.ParseCtx(context.Background(), []byte(code), json.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(document (object (pair key: (string (string_content)) value: (string (string_content))) (pair key: (string (string_content)) value: (number)) (pair key: (string (string_content)) value: (number)) (pair key: (string (string_content)) value: (true)) (pair key: (string (string_content)) value: (null)) (pair key: (string (string_content)) value: (array (number) (number) (number))) (pair key: (string (string_content)) value: (object (pair key: (string (string_content)) value: (number))))))",
		n.String(),
	)
}
