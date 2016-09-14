// Copyright 2014 Xuanyi Chew. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// The original code is lifted from the Go standard library which is governed by
// a BSD-style licence which can be found here: https://golang.org/LICENSE
package math32

import "math"

func Jn(n int, x float32) float32 {
	return float32(math.Jn(n, float64(x)))
}

func Yn(n int, x float32) float32 {
	return float32(math.Yn(n, float64(x)))
}
