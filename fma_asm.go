//go:build !noasm && !tinygo && (arm64 || ppc64le || riscv64 || s390x || (loong64 && go1.24))
// +build !noasm
// +build !tinygo
// +build arm64 ppc64le riscv64 s390x loong64,go1.24

package math32

const haveArchFMA = true

func archFMA(x, y, z float32) float32
