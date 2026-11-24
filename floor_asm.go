//go:build !noasm && !tinygo && (arm64 || ppc64le || riscv64 || s390x || (loong64 && go1.24))
// +build !noasm
// +build !tinygo
// +build arm64 ppc64le riscv64 s390x loong64,go1.24

package math32

const haveArchFloor = true

func archFloor(x float32) float32

const haveArchCeil = true

func archCeil(x float32) float32

const haveArchTrunc = true

func archTrunc(x float32) float32
