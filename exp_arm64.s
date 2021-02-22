// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Exp returns e**x, the base-e exponential of x.
// This is an assembly implementation of the method used for function Exp in file exp.go.
//
// func Exp(x float64) float64
TEXT ·Exp(SB),$0-16
	B  ·exp(SB)

// Exp2 returns 2**x, the base-2 exponential of x.
// This is an assembly implementation of the method used for function Exp2 in file exp.go.
//
// func Exp2(x float64) float64
TEXT ·Exp2(SB),$0-16
	B  ·exp2(SB)
