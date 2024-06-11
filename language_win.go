//go:build windows

package sitter

import (
	"golang.org/x/sys/windows"
)

func loadLib(path string) (uintptr, error) {
	handle, err := windows.LoadLibrary(path)
	return uintptr(handle), err
}
