// +build amd64

package simd

func Available() bool {
	return SSE2()
}

func CpuId(info *[4]uint32, ax uint32)

// SSE2 returns true if the the CPU supports SSE2 instructions
func SSE2() bool {
	var info [4]uint32
	CpuId(&info, 1)
	return info[3]&(1<<26) != 0 // SSE2
}

// SSSE3 returns true if the the CPU supports SSSE3 instructions
func SSSE3() bool
