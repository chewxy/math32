//go:build noasm || tinygo || (!arm64 && !riscv64)
// +build noasm tinygo !arm64,!riscv64

package math32

const haveArchFMA = false

func archFMA(x, y, z float32) float32 {
	panic("not implemented")
}
