// (copied from Nigel Tao's shiny library, https://github.com/golang/exp/blob/master/shiny/driver/internal/swizzle/swizzle_amd64.s)
// func haveSSSE3() bool
TEXT ·SSSE3(SB),NOSPLIT,$0
        MOVQ	$1, AX
        CPUID
        SHRQ	$9, CX
        ANDQ	$1, CX
        MOVB	CX, ret+0(FP)
        RET

// (copied from the Go dist tool, https://github.com/golang/go/blob/master/src/cmd/dist/cpuid_amd64.s)
TEXT ·CpuId(SB),$0-12
        MOVL ax+8(FP), AX
        CPUID
        MOVQ info+0(FP), DI
        MOVL AX, 0(DI)
        MOVL BX, 4(DI)
        MOVL CX, 8(DI)
        MOVL DX, 12(DI)
        RET
