// +build amd64

TEXT ·addpd(SB),NOSPLIT,$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVUPD       x+0(FP), X15
        MOVUPD       y+16(FP), X14
        MOVUPD       X15, t2-48(SP)
        MOVUPD       X14, t3-64(SP)
        MOVUPD       t2-48(SP), X13
        MOVUPD       t3-64(SP), X12
        ADDPD        X13, X12
        MOVUPD       X12, ret0+32(FP)
        RET

