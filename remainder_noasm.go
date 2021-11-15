//go:build noasm || (!amd64 && !s390x)
// +build noasm !amd64,!s390x

package math32

const haveArchRemainder = false

func archRemainder(x, y float32) float32 {
	panic("not implemented")
}
