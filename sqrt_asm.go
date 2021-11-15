// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !noasm && (386 || amd64 || arm64 || arm || ppc64le || s390x || riscv64 || wasm)
// +build !noasm
// +build 386 amd64 arm64 arm ppc64le s390x riscv64 wasm

package math32

const haveArchSqrt = true

func archSqrt(x float32) float32
