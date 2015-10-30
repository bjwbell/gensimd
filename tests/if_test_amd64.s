// +build amd64

TEXT ·ift0s(SB),NOSPLIT,$8-9
        MOVB         $0, ret0+8(FP)
        MOVB         $0, t0-1(SP)
        MOVB         $0, t1-2(SP)
block0:
        MOVB         x+0(FP), R14
        MOVB         $2, R13
        CMPB         R14, R13
        SETCS        R15
        MOVB         R15, t0-1(SP)
        MOVB         t0-1(SP), R15
        CMPB         R15, $0
        JEQ          block2
        JMP          block1
block1:
        MOVB         x+0(FP), R15
        MOVB         R15, ret0+8(FP)
        RET
block2:
        MOVB         x+0(FP), R14
        MOVB         x+0(FP), R13
        MOVB         R14, R15
        MOVB         R15, AX
        MULB         R13
        MOVB         AX, R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·ift1s(SB),NOSPLIT,$8-10
        MOVW         $0, ret0+8(FP)
        MOVB         $0, t0-1(SP)
        MOVW         $0, t1-3(SP)
block0:
        MOVW         x+0(FP), R14
        MOVW         $128, R13
        CMPW         R14, R13
        SETHI        R15
        MOVB         R15, t0-1(SP)
        MOVB         t0-1(SP), R15
        CMPB         R15, $0
        JEQ          block2
        JMP          block1
block1:
        MOVW         x+0(FP), R15
        XORQ         $-1, R15
        MOVW         R15, ret0+8(FP)
        RET
block2:
        MOVW         x+0(FP), R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·ift2s(SB),NOSPLIT,$16-12
        MOVL         $0, ret0+8(FP)
        MOVL         $0, t2-9(SP)
        MOVB         $0, t0-1(SP)
        MOVL         $0, t1-5(SP)
block0:
        MOVL         x+0(FP), R14
        MOVL         $1024, R13
        CMPL         R14, R13
        SETCS        R15
        MOVB         R15, t0-1(SP)
        MOVB         t0-1(SP), R15
        CMPB         R15, $0
        JEQ          block2
        JMP          block1
block1:
        MOVL         x+0(FP), R14
        MOVL         $509, R13
        MOVL         R13, R15
        ANDL         R14, R15
        MOVL         R15, ret0+8(FP)
        RET
block2:
        MOVL         x+0(FP), R13
        MOVL         $511, R12
        MOVL         R12, R14
        ANDL         R13, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·ift3s(SB),NOSPLIT,$32-16
        MOVQ         $0, ret0+8(FP)
        MOVQ         $0, t2-17(SP)
        MOVQ         $0, t3-25(SP)
        MOVQ         $0, t0-8(SP)
        MOVB         $0, t1-9(SP)
block0:
        MOVQ         x+0(FP), R14
        MOVQ         x+0(FP), R13
        MOVQ         R14, R15
        MOVQ         R15, AX
        MULQ         R13
        MOVQ         AX, R15
        MOVQ         $2046, R13
        CMPQ         R15, R13
        SETCS        R14
        MOVB         R14, t1-9(SP)
        MOVQ         R15, t0-8(SP)
        MOVB         t1-9(SP), R15
        CMPB         R15, $0
        JEQ          block2
        JMP          block1
block1:
        MOVQ         x+0(FP), R15
        MOVQ         R15, ret0+8(FP)
        RET
block2:
        MOVQ         $2, R14
        MOVQ         x+0(FP), R13
        MOVQ         R14, R15
        MOVQ         R15, AX
        MULQ         R13
        MOVQ         AX, R15
        MOVQ         x+0(FP), R13
        MOVQ         R15, R14
        SUBQ         R13, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·ift4s(SB),NOSPLIT,$8-9
        MOVB         $0, ret0+8(FP)
        MOVB         $0, t0-1(SP)
        MOVB         $0, t1-2(SP)
block0:
        MOVB         x+0(FP), R14
        MOVB         $0, R13
        CMPB         R14, R13
        SETLT        R15
        MOVB         R15, t0-1(SP)
        MOVB         t0-1(SP), R15
        CMPB         R15, $0
        JEQ          block2
        JMP          block1
block1:
        MOVB         x+0(FP), R13
        XORQ         R14, R14
        MOVB         R14, R15
        SUBB         R13, R15
        MOVB         R15, ret0+8(FP)
        RET
block2:
        MOVB         x+0(FP), R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·ift5s(SB),NOSPLIT,$8-10
        MOVW         $0, ret0+8(FP)
        MOVB         $0, t0-1(SP)
        MOVW         $0, t1-3(SP)
        MOVW         $0, t2-5(SP)
block0:
        MOVW         x+0(FP), R14
        MOVW         $-255, R13
        CMPW         R14, R13
        SETLT        R15
        MOVB         R15, t0-1(SP)
        MOVB         t0-1(SP), R15
        CMPB         R15, $0
        JEQ          block2
        JMP          block1
block1:
        MOVW         $-255, R14
        MOVW         x+0(FP), R13
        MOVW         R14, R15
        MOVW         R15, AX
        IMULW        R13
        MOVW         AX, R15
        MOVW         R15, ret0+8(FP)
        RET
block2:
        MOVW         $255, R13
        MOVW         x+0(FP), R12
        MOVW         R13, R14
        MOVW         R14, AX
        IMULW        R12
        MOVW         AX, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·ift6s(SB),NOSPLIT,$8-12
        MOVL         $0, ret0+8(FP)
        MOVB         $0, t0-1(SP)
block0:
        MOVL         x+0(FP), R14
        MOVL         $1, R13
        CMPL         R14, R13
        SETEQ        R15
        MOVB         R15, t0-1(SP)
        MOVB         t0-1(SP), R15
        CMPB         R15, $0
        JEQ          block2
        JMP          block1
block1:
        MOVL         $0, R15
        MOVL         R15, ret0+8(FP)
        RET
block2:
        MOVL         $1, R15
        MOVL         R15, ret0+8(FP)
        RET

TEXT ·ift7s(SB),NOSPLIT,$24-16
        MOVQ         $0, ret0+8(FP)
        MOVB         $0, t0-1(SP)
        MOVQ         $0, t1-9(SP)
        MOVQ         $0, t2-17(SP)
block0:
        MOVQ         x+0(FP), R14
        MOVQ         $-1, R13
        CMPQ         R14, R13
        SETLT        R15
        MOVB         R15, t0-1(SP)
        MOVB         t0-1(SP), R15
        CMPB         R15, $0
        JEQ          block2
        JMP          block1
block1:
        MOVQ         x+0(FP), R13
        XORQ         R14, R14
        MOVQ         R14, R15
        SUBQ         R13, R15
        MOVQ         R15, ret0+8(FP)
        RET
block2:
        MOVQ         $10, R13
        MOVQ         x+0(FP), R12
        MOVQ         R13, R14
        SUBQ         R12, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·ift8s(SB),NOSPLIT,$16-12
        MOVL         $0, ret0+8(FP)
        MOVB         $0, t0-1(SP)
        MOVL         $0, t1-5(SP)
        MOVL         $0, t2-9(SP)
block0:
        MOVL         x+0(FP), X15
        //           $1065353216 = 000000003f800000 = 1(float32)
        MOVQ         $1065353216, R14
        MOVQ         R14, X14
        UCOMISS      X15, X14
        SETHI        R15
        MOVB         R15, t0-1(SP)
        MOVB         t0-1(SP), R15
        CMPB         R15, $0
        JEQ          block2
        JMP          block1
block1:
        MOVL         x+0(FP), X13
        XORPD        X14, X14
        MOVO         X14, X15
        SUBSS        X13, X15
        MOVL         X15, ret0+8(FP)
        RET
block2:
        //           $1092616192 = 0000000041200000 = 10(float32)
        MOVQ         $1092616192, R15
        MOVQ         R15, X13
        MOVL         x+0(FP), X12
        MOVO         X13, X14
        SUBSS        X12, X14
        MOVL         X14, ret0+8(FP)
        RET

TEXT ·ift9s(SB),NOSPLIT,$32-16
        MOVQ         $0, ret0+8(FP)
        MOVB         $0, t1-9(SP)
        MOVQ         $0, t2-17(SP)
        MOVQ         $0, t3-25(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVQ         x+0(FP), X14
        MOVQ         x+0(FP), X13
        MOVO         X14, X15
        MULSD        X13, X15
        //           $4656713218608070656 = 409ff80000000000 = 2046(float64)
        MOVQ         $4656713218608070656, R14
        MOVQ         R14, X14
        UCOMISD      X15, X14
        SETHI        R15
        MOVB         R15, t1-9(SP)
        MOVQ         X15, t0-8(SP)
        MOVB         t1-9(SP), R15
        CMPB         R15, $0
        JEQ          block2
        JMP          block1
block1:
        MOVQ         x+0(FP), X15
        MOVQ         X15, ret0+8(FP)
        RET
block2:
        //           $4613937818241073152 = 4008000000000000 = 3(float64)
        MOVQ         $4613937818241073152, R15
        MOVQ         R15, X14
        MOVQ         x+0(FP), X13
        MOVO         X14, X15
        MULSD        X13, X15
        MOVQ         x+0(FP), X13
        MOVO         X15, X14
        SUBSD        X13, X14
        MOVQ         X14, ret0+8(FP)
        RET

