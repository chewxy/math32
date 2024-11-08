// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !noasm && !tinygo && (amd64 || 386 || arm || ppc64le || riscv64 || wasm)
// +build !noasm
// +build !tinygo
// +build amd64 386 arm ppc64le riscv64 wasm

package math32

const haveArchExp2 = true

func archExp2(x float32) float32
