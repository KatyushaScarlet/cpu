// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build loong64

// func cpucfg() (uint32, uint32)
TEXT ·cpucfg(SB), NOSPLIT, $0-16
    MOVL R1, index+0(FP)
    CPUCFG	R2, R1
    MOVL R2, eax+8(FP)
    RET
