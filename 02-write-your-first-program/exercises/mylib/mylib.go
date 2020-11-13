package mylib

import (
	"runtime"
)

// Version returns the Go version
func Version() string {
	return runtime.Version()
}
