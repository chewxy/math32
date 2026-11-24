//go:build !noasm && !tinygo && (arm64 || riscv64)
// +build !noasm
// +build !tinygo
// +build arm64 riscv64

package math32

const haveArchFMA = true

func archFMA(x, y, z float32) float32
