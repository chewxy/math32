// +build !amd64

#include "textflag.h"

// func Log(x float64) float64
TEXT ·Log(SB),NOSPLIT,$0
	JMP ·log(SB)
