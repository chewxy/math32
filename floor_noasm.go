//go:build noasm || tinygo || (!arm64 && !ppc64le && !riscv64 && !s390x)
// +build noasm tinygo !arm64,!ppc64le,!riscv64,!s390x

package math32

const haveArchFloor = false

func archFloor(x float32) float32 {
	panic("not implemented")
}

const haveArchCeil = false

func archCeil(x float32) float32 {
	panic("not implemented")
}

const haveArchTrunc = false

func archTrunc(x float32) float32 {
	panic("not implemented")
}
