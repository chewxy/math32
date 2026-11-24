//go:build noasm || tinygo || (!arm64 && !ppc64le && !riscv64 && !s390x && !loong64) || (!arm64 && !ppc64le && !riscv64 && !s390x && !go1.24)
// +build noasm tinygo !arm64,!ppc64le,!riscv64,!s390x,!loong64 !arm64,!ppc64le,!riscv64,!s390x,!go1.24

package math32

const haveArchFMA = false

func archFMA(x, y, z float32) float32 {
	panic("not implemented")
}
