// +build !amd64

#include "textflag.h"

// func Sqrt(x float32) float32
TEXT ·Sqrt(SB),NOSPLIT,$0
	JMP ·sqrt(SB)
