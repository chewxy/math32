//go:build noasm || tinygo || (!amd64 && !arm && !386 && !wasm)
// +build noasm tinygo !amd64,!arm,!386,!wasm

package math32

const haveArchRemainder = false

func archRemainder(x, y float32) float32 {
	panic("not implemented")
}
