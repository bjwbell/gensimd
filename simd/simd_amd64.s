// func haveSSSE3() bool
TEXT ·SSSE3(SB),NOSPLIT,$0
        MOVQ	$1, AX
        CPUID
        SHRQ	$9, CX
        ANDQ	$1, CX
        MOVB	CX, ret+0(FP)
        RET


TEXT ·cpuid(SB),$0-12
        MOVL ax+8(FP), AX
        CPUID
        MOVQ info+0(FP), DI
        MOVL AX, 0(DI)
        MOVL BX, 4(DI)
        MOVL CX, 8(DI)
        MOVL DX, 12(DI)
        RET
