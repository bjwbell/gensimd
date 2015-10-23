// +build !amd64

package simd

func Available() bool { return false }
func SSE2() bool      { panic("unreachable") }
func SSSE3() bool     { panic("unreachable") }
