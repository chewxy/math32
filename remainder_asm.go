//go:build !noasm && !tinygo && (amd64 || arm || 386 || wasm)
// +build !noasm
// +build !tinygo
// +build amd64 arm 386 wasm

package math32

const haveArchRemainder = true

func archRemainder(x, y float32) float32
