//go:build !tinygo && !noasm && go1.24

#include "textflag.h"

// The Go assembler supports FTINT*, FFINT* and FCOPYSG* on loong64 since Go 1.24.
//
// LoongArch has no float-to-float round-to-integral instruction with a fixed
// rounding mode, so these round through an int32 with FTINTR{M,P,Z}WF and back
// with FFINTFW.
//
// Inputs with |x| >= 2^23, ±Inf and NaN are already integral and are returned
// unchanged. Every other input fits in an int32, so the conversion is exact.
// FCOPYSGF then copies the sign of x onto the result, so that results that
// round to zero keep the sign of x, e.g. Ceil(-0.5) = -0. No nonzero result
// can have a sign different from x, so this is always correct.

#define ROUNDFN(NAME, FTINT)	\
TEXT NAME(SB),NOSPLIT,$0-12;	\
	MOVF	x+0(FP), F0;	\
	MOVWU	x+0(FP), R11;	\
	SRLV	$23, R11, R11;	\
	AND	$0xFF, R11, R11;	/* R11 = biased exponent of x */	\
	MOVV	$149, R12;	/* 149 = 127 + 22 */	\
	BLTU	R12, R11, done;	/* |x| >= 2^23, ±Inf or NaN */	\
	FTINT	F0, F1;	\
	FFINTFW	F1, F1;	\
	FCOPYSGF	F0, F1, F0;	/* F0 = |F1| with the sign of x */	\
done:;	\
	MOVF	F0, ret+8(FP);	\
	RET

// func archFloor(x float32) float32
ROUNDFN(·archFloor, FTINTRMWF)

// func archCeil(x float32) float32
ROUNDFN(·archCeil, FTINTRPWF)

// func archTrunc(x float32) float32
ROUNDFN(·archTrunc, FTINTRZWF)
