// Copyright 2014 Xuanyi Chew. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// The original code is lifted from the Go standard library which is governed by
// a BSD-style licence which can be found here: https://golang.org/LICENSE
package math32

import "math"

func Sincos(x float32) (sin, cos float32) {
	sin64, cos64 := math.Sincos(float64(x))
	sin = float32(sin64)
	cos = float32(cos64)
	return
}
