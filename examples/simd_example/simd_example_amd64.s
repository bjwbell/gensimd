// +build amd64

TEXT ·addi32x4(SB),NOSPLIT,$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PADDL        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·subi32x4(SB),NOSPLIT,$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PSUBL        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·muli32x4(SB),NOSPLIT,$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t2-48(SP), X12
        MOVOU        t3-64(SP), X11
        MOVO         X11, X13
        PMULULQ      X12, X13
        PSRLO        $4, X12
        PSRLO        $4, X11
        MOVO         X11, X10
        PMULULQ      X12, X10
        PSHUFD       $8, X13, X9
        PSHUFD       $8, X10, X8
        PUNPCKLLQ    X8, X9
        MOVOU        X9, ret0+32(FP)
        RET

TEXT ·shli32x4(SB),NOSPLIT,$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X14
        MOVOU        t1-32(SP), X13
        PSLLL        X14, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shri32x4(SB),NOSPLIT,$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X14
        MOVOU        t1-32(SP), X13
        PSRAL        X14, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·addf32x4(SB),NOSPLIT,$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVUPS       y+16(FP), X14
        MOVUPS       X15, t2-48(SP)
        MOVUPS       X14, t3-64(SP)
        MOVUPS       t3-64(SP), X13
        MOVUPS       t2-48(SP), X12
        ADDPS        X13, X12
        MOVUPS       X12, ret0+32(FP)
        RET

TEXT ·subf32x4(SB),NOSPLIT,$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVUPS       y+16(FP), X14
        MOVUPS       X15, t2-48(SP)
        MOVUPS       X14, t3-64(SP)
        MOVUPS       t3-64(SP), X13
        MOVUPS       t2-48(SP), X12
        SUBPS        X13, X12
        MOVUPS       X12, ret0+32(FP)
        RET

TEXT ·mulf32x4(SB),NOSPLIT,$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVUPS       y+16(FP), X14
        MOVUPS       X15, t2-48(SP)
        MOVUPS       X14, t3-64(SP)
        MOVUPS       t3-64(SP), X13
        MOVUPS       t2-48(SP), X12
        MULPS        X13, X12
        MOVUPS       X12, ret0+32(FP)
        RET

TEXT ·divf32x4(SB),NOSPLIT,$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVUPS       y+16(FP), X14
        MOVUPS       X15, t2-48(SP)
        MOVUPS       X14, t3-64(SP)
        MOVUPS       t3-64(SP), X13
        MOVUPS       t2-48(SP), X12
        DIVPS        X13, X12
        MOVUPS       X12, ret0+32(FP)
        RET

