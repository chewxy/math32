// +build !amd64

#include "textflag.h"

// func Exp(x float32) float32
TEXT ·Exp(SB),NOSPLIT,$0
	JMP ·exp(SB)
