// +build experimental


// Copyright 2014 Xuanyi Chew. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// The original code is lifted from the Go standard library which is governed by
// a BSD-style licence which can be found here: https://golang.org/LICENSE

#include "textflag.h"

// The method is based on a paper by Naoki Shibata: "Efficient evaluation
// methods of elementary functions suitable for SIMD computation", Proc.
// of International Supercomputing Conference 2010 (ISC'10), pp. 25 -- 32
// (May 2010). The paper is available at
// http://www.springerlink.com/content/340228x165742104/
//
// The original code and the constants below are from the author's
// implementation available at http://freshmeat.net/projects/sleef.
// The README file says, "The software is in public domain.
// You can use the software without any obligation."
//
// This code is a simplified version of the original.
// The magic numbers for the float32 are lifted from the same project

/*
NOTES

For time being this function is not going to be used. The reason?
It works when put in a separate program, but not within this package.
Specifically when using this, the Exp tests would run something like this:
		exp(4.9790119248836735e+00) = 145.33072
		exp(7.7388724745781045e+00) = 7.7388725
		exp(-2.7688005719200159e-01) = -0.27688006
		exp(-5.0106036182710749e+00) = -5.0106034
		exp(9.6362937071984173e+00) = 9.636293
		exp(2.9263772392439646e+00) = 2.9263773
		exp(5.2290834314593066e+00) = 5.2290835
		exp(2.7279399104360102e+00) = 2.7279398
		exp(1.8253080916808550e+00) = 1.8253081
		exp(-8.6859247685756013e+00) = -8.685925
where only the first result is correct and the following results are wrong.
I suspect it has to do with something I'm missing about the MOVSS instruction
Either way, I'll come back to this when I'm free
*/
	
#define LN2 0.693147182464599609375 // log_e(2)
#define LOG2E 1.44269502162933349609375 // 1/LN2
#define LN2U 0.693145751953125 // upper half LN2
#define LN2L 1.428606765330187045e-06 // lower half LN2
#define T0 1.0
#define T1 0.5
#define T2 0.166665524244308471679688
#define T3 0.0416710823774337768554688
#define T4 0.00836596917361021041870117
#define PosInf 0x7FF00000
#define NegInf 0xFFF00000

// func Exp(x float64) float64
TEXT ·Exp(SB),NOSPLIT,$0
// test bits for not-finite
	MOVQ    x+0(FP), BX
	MOVQ    $~(1<<31), AX // sign bit mask
	MOVQ    BX, DX
	ANDQ    AX, DX
	MOVQ    $PosInf, AX
	CMPQ    AX, DX
	JLE     notFinite
	MOVQ    BX, X0
	MOVSS   $LOG2E, X1
	MULSS   X0, X1
	CVTSS2SL X1, BX // BX = exponent
	CVTSL2SS BX, X1
	MOVSS   $LN2U, X2
	MULSS   X1, X2
	SUBSS   X2, X0
	MOVSS   $LN2L, X2
	MULSS   X1, X2
	SUBSS   X2, X0
	// reduce argument
	MULSS   $0.0625, X0
	// Taylor series evaluation
	ADDSS   $T4, X1
	MULSS   X0, X1
	ADDSS   $T3, X1
	MULSS   X0, X1
	ADDSS   $T2, X1
	MULSS   X0, X1
	ADDSS   $T1, X1
	MULSS   X0, X1
	ADDSS   $T0, X1
	MULSS   X1, X0
	MOVSS   $2.0, X1
	ADDSS   X0, X1
	MULSS   X1, X0
	MOVSS   $2.0, X1
	ADDSS   X0, X1
	MULSS   X1, X0
	MOVSS   $2.0, X1
	ADDSS   X0, X1
	MULSS   X1, X0
	MOVSS   $2.0, X1
	ADDSS   X0, X1
	MULSS   X1, X0
	ADDSS   $1.0, X0
	// return fr * 2**exponent
	MOVQ    $0x7F, AX // bias
	ADDQ    AX, BX
	JLE     underflow
	CMPQ    BX, $0xFF
	JGE     overflow
	MOVQ    $23, CX
	SHLQ    CX, BX
	MOVQ    BX, X1
	MULSS   X1, X0
	MOVSS   X0, ret+8(FP)
	RET
notFinite:
	// test bits for -Inf
	MOVQ    $NegInf, AX
	CMPQ    AX, BX
	JNE     notNegInf
	// -Inf, return 0
underflow: // return 0
	MOVQ    $0, AX
	MOVQ    AX, ret+8(FP)
	RET
overflow: // return +Inf
	MOVQ    $PosInf, BX
notNegInf: // NaN or +Inf, return x
	MOVQ    BX, ret+8(FP)
	RET
