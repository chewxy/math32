// +build !amd64

#include "textflag.h"

TEXT ·Remainder(SB),NOSPLIT,$0
	JMP ·remainder(SB)
