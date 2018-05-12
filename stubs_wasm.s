// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT ·Exp(SB),NOSPLIT,$0
	JMP ·exp(SB)

TEXT ·Log(SB),NOSPLIT,$0
	JMP ·log(SB)

TEXT ·Remainder(SB),NOSPLIT,$0
	JMP ·remainder(SB)

TEXT ·Sqrt(SB),NOSPLIT,$0
	JMP ·sqrt(SB)
