// +build amd64

TEXT ·ptrt0s(SB),NOSPLIT,$16-12
        MOVL         $0, ret0+8(FP)
        MOVL         $0, t0-4(SP)
        MOVL         $0, t1-8(SP)
block0:
        MOVQ         x+0(FP), R15
        MOVSS        (R15), X15
        MOVSS        X15, t0-4(SP)
        //           $1073741824 = 0000000040000000 = 2(float32)
        MOVQ         $1073741824, R15
        MOVQ         R15, X14
        MOVL         t0-4(SP), X13
        MOVO         X14, X15
        MULSS        X13, X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·ptrt1s(SB),NOSPLIT,$40-16
        MOVQ         $0, ret0+8(FP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-16(SP)
        MOVQ         $0, t2-24(SP)
        MOVQ         $0, t3-32(SP)
block0:
        MOVQ         x+0(FP), R15
        MOVSD        (R15), X15
        MOVSD        X15, t0-8(SP)
        //           $4611686018427387904 = 4000000000000000 = 2(float64)
        MOVQ         $4611686018427387904, R15
        MOVQ         R15, X14
        MOVQ         t0-8(SP), X13
        MOVO         X14, X15
        MULSD        X13, X15
        MOVQ         x+0(FP), R15
        MOVSD        (R15), X14
        MOVSD        X14, t2-24(SP)
        MOVQ         t2-24(SP), X13
        MOVO         X15, X14
        ADDSD        X13, X14
        MOVQ         X14, ret0+8(FP)
        RET

TEXT ·addf32s(SB),NOSPLIT,$8-12
        MOVL         $0, ret0+8(FP)
        MOVL         $0, t0-4(SP)
block0:
        MOVL         x+0(FP), X14
        MOVL         y+4(FP), X13
        MOVO         X14, X15
        ADDSS        X13, X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·subf32s(SB),NOSPLIT,$8-12
        MOVL         $0, ret0+8(FP)
        MOVL         $0, t0-4(SP)
block0:
        MOVL         x+0(FP), X14
        MOVL         y+4(FP), X13
        MOVO         X14, X15
        SUBSS        X13, X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·negf32s(SB),NOSPLIT,$8-12
        MOVL         $0, ret0+8(FP)
        MOVL         $0, t0-4(SP)
block0:
        MOVL         x+0(FP), X13
        XORPD        X14, X14
        MOVO         X14, X15
        SUBSS        X13, X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·mulf32s(SB),NOSPLIT,$8-12
        MOVL         $0, ret0+8(FP)
        MOVL         $0, t0-4(SP)
block0:
        MOVL         x+0(FP), X14
        MOVL         y+4(FP), X13
        MOVO         X14, X15
        MULSS        X13, X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·divf32s(SB),NOSPLIT,$8-12
        MOVL         $0, ret0+8(FP)
        MOVL         $0, t0-4(SP)
block0:
        MOVL         x+0(FP), X14
        MOVL         y+4(FP), X13
        MOVO         X14, X15
        DIVSS        X13, X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·addf64s(SB),NOSPLIT,$16-24
        MOVQ         $0, ret0+16(FP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVQ         x+0(FP), X14
        MOVQ         y+8(FP), X13
        MOVO         X14, X15
        ADDSD        X13, X15
        MOVQ         X15, ret0+16(FP)
        RET

TEXT ·subf64s(SB),NOSPLIT,$16-24
        MOVQ         $0, ret0+16(FP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVQ         x+0(FP), X14
        MOVQ         y+8(FP), X13
        MOVO         X14, X15
        SUBSD        X13, X15
        MOVQ         X15, ret0+16(FP)
        RET

TEXT ·negf64s(SB),NOSPLIT,$16-16
        MOVQ         $0, ret0+8(FP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVQ         x+0(FP), X13
        XORPD        X14, X14
        MOVO         X14, X15
        SUBSD        X13, X15
        MOVQ         X15, ret0+8(FP)
        RET

TEXT ·mulf64s(SB),NOSPLIT,$16-24
        MOVQ         $0, ret0+16(FP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVQ         x+0(FP), X14
        MOVQ         y+8(FP), X13
        MOVO         X14, X15
        MULSD        X13, X15
        MOVQ         X15, ret0+16(FP)
        RET

TEXT ·divf64s(SB),NOSPLIT,$16-24
        MOVQ         $0, ret0+16(FP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVQ         x+0(FP), X14
        MOVQ         y+8(FP), X13
        MOVO         X14, X15
        DIVSD        X13, X15
        MOVQ         X15, ret0+16(FP)
        RET

