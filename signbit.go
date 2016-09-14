// Copyright 2014 Xuanyi Chew. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// The original code is lifted from the Go standard library which is governed by
// a BSD-style licence which can be found here: https://golang.org/LICENSE
package math32

// Signbit returns true if x is negative or negative zero.
func Signbit(x float32) bool {
	return Float32bits(x)&(1<<31) != 0
}
