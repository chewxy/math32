//go:build !noasm && (amd64 || s390x)
// +build !noasm
// +build amd64 s390x

package math32

const haveArchRemainder = true

func archRemainder(x, y float32) float32
