package sitter

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
)

func LoadLanguage(symbolName string, path string) (*Language, error) {
	lib, err := purego.Dlopen(path, purego.RTLD_NOW)
	if err != nil {
		return nil, fmt.Errorf("failed to open language library: %w", err)
	}

	var newTreeSitter func() uintptr
	purego.RegisterLibFunc(&newTreeSitter, lib, "tree_sitter_"+symbolName)

	return NewLanguage(unsafe.Pointer(newTreeSitter())), nil
}
