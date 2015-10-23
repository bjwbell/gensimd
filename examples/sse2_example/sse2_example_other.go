// +build !amd64,gc

package main

import (
	"fmt"

	"github.com/bjwbell/gensimd/simd"
	"github.com/bjwbell/gensimd/simd/sse2"
)

func addpd(x, y simd.M128d) simd.M128d { return sse2.AddPd(x, y) }

func main() {
	// sse2 is not supported on non-amd64 platforms
	fmt.Println("SSE2 not available")
}
