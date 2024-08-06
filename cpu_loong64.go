// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

const CacheLineSize = 64

// cpucfg is implemented in cpu_loong64.s.
func cpucfg(index uint32) (result uint32)

const (
	cpucfg_lam_word = 0x2
	cpucfg_lam_bit  = 1 << 22

	cpucfg_ual_word = 0x1
	cpucfg_ual_bit  = 1 << 20

	cpucfg_fpu_word = 0x2
	cpucfg_fpu_bit  = 1 << 0

	cpucfg_lsx_word = 0x2
	cpucfg_lsx_bit  = 1 << 6

	cpucfg_lasx_word = 0x2
	cpucfg_lasx_bit  = 1 << 7

	cpucfg_crc32_word = 0x1
	cpucfg_crc32_bit  = 1 << 25

	cpucfg_complex_word = 0x2
	cpucfg_complex_bit  = 1 << 8

	cpucfg_crypto_word = 0x2
	cpucfg_crypto_bit  = 1 << 9

	cpucfg_lvz_word = 0x2
	cpucfg_lvz_bit  = 1 << 10

	cpucfg_lbt_x86_word = 0x2
	cpucfg_lbt_x86_bit  = 1 << 18

	cpucfg_lbt_arm_word = 0x2
	cpucfg_lbt_arm_bit  = 1 << 19

	cpucfg_lbt_mips_word = 0x2
	cpucfg_lbt_mips_bit  = 1 << 20

	cpucfg_ptw_word = 0x2
	cpucfg_ptw_bit  = 1 << 24

	cpucfg_prid_word = 0x0
)

func doinit() {
	options = []option{
		{"lam", &Loong64.HasLAM},
		{"ual", &Loong64.HasUAL},
		{"fpu", &Loong64.HasFPU},
		{"lsx", &Loong64.HasLSX},
		{"lasx", &Loong64.HasLASX},
		{"crc32", &Loong64.HasCRC32},
		{"complex", &Loong64.HasCOMPLEX},
		{"crypto", &Loong64.HasCRYPTO},
		{"lvz", &Loong64.HasLVZ},
		{"lbt_x86", &Loong64.HasLBT_X86},
		{"lbt_arm", &Loong64.HasLBT_ARM},
		{"lbt_mips", &Loong64.HasLBT_MIPS},
		{"ptw", &Loong64.HasPTW},
	}
}

func isSet(word, bit) bool {
	return cpucfg(word)&bit != 0
}

func getPRID() uint32 {
	return cpucfg(cpucfg_prid_word)
}
