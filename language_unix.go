//go:build !windows

package sitter

import (
	"github.com/ebitengine/purego"
)

func loadLib(path string) (uintptr, error) {
	return purego.Dlopen(path, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}
