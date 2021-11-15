// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build noasm || !(arm64 || amd64 || 386 || arm || ppc64le || risc64 || wasm)
// +build noasm !arm64,!amd64,!386,!arm,!ppc64le,!risc64,!wasm

package math32

const haveArchExp2 = false

func archExp2(x float32) float32 {
	panic("not implemented")
}
