//go:build !noasm && !tinygo && (arm64 || riscv64)
// +build !noasm
// +build !tinygo
// +build arm64 riscv64

package math32

const haveArchFloor = true

func archFloor(x float32) float32

const haveArchCeil = true

func archCeil(x float32) float32

const haveArchTrunc = true

func archTrunc(x float32) float32
