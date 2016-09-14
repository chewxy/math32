// Copyright 2014 Xuanyi Chew. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// The original code is lifted from the Go standard library which is governed by
// a BSD-style licence which can be found here: https://golang.org/LICENSE
package math32

import "math"

func Lgamma(x float32) (lgamma float32, sign int) {
	lg, sign := math.Lgamma(float64(x))
	return float32(lg), sign
}
