// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·addi32x4(SB),$88-48
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PADDL        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
        RET

TEXT ·subi32x4(SB),$88-48
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
        MOVOU        x+0(FP), X15
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PSUBL        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
        RET

TEXT ·muli32x4(SB),$88-48
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
        MOVOU        x+0(FP), X15
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t2-48(SP), X10
        MOVOU        t3-64(SP), X9
        MOVO         X9, X11
        PMULULQ      X10, X11
        PSRLO        $4, X10
        PSRLO        $4, X9
        MOVO         X9, X8
        PMULULQ      X10, X8
        PSHUFD       $8, X11, X7
        PSHUFD       $8, X8, X6
        PUNPCKLLQ    X6, X7
        MOVOU        X7, ret0+32(FP)
        RET

TEXT ·shli32x4(SB),$56-40
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
        MOVOU        X15, t0-16(SP)
        MOVOU        t0-16(SP), X14
        MOVOU        X14, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X13
        MOVOU        t1-32(SP), X12
        PSLLL        X13, X12
        MOVOU        X12, t2-48(SP)
        MOVOU        t2-48(SP), X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shri32x4(SB),$56-40
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
        MOVOU        X15, t0-16(SP)
        MOVOU        t0-16(SP), X14
        MOVOU        X14, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X13
        MOVOU        t1-32(SP), X12
        PSRAL        X13, X12
        MOVOU        X12, t2-48(SP)
        MOVOU        t2-48(SP), X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·addf32x4(SB),$88-48
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
        MOVUPS       X15, t0-16(SP)
        MOVUPS       y+16(FP), X14
        MOVUPS       X14, t1-32(SP)
        MOVUPS       t0-16(SP), X13
        MOVUPS       X13, t2-48(SP)
        MOVUPS       t1-32(SP), X12
        MOVUPS       X12, t3-64(SP)
        MOVUPS       t3-64(SP), X11
        MOVUPS       t2-48(SP), X10
        ADDPS        X11, X10
        MOVUPS       X10, t4-80(SP)
        MOVUPS       t4-80(SP), X9
        MOVUPS       X9, ret0+32(FP)
        RET

TEXT ·subf32x4(SB),$88-48
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
        MOVUPS       x+0(FP), X15
        MOVUPS       X15, t0-16(SP)
        MOVUPS       y+16(FP), X14
        MOVUPS       X14, t1-32(SP)
        MOVUPS       t0-16(SP), X13
        MOVUPS       X13, t2-48(SP)
        MOVUPS       t1-32(SP), X12
        MOVUPS       X12, t3-64(SP)
        MOVUPS       t3-64(SP), X11
        MOVUPS       t2-48(SP), X10
        SUBPS        X11, X10
        MOVUPS       X10, t4-80(SP)
        MOVUPS       t4-80(SP), X9
        MOVUPS       X9, ret0+32(FP)
        RET

TEXT ·mulf32x4(SB),$88-48
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
        MOVUPS       X15, t0-16(SP)
        MOVUPS       y+16(FP), X14
        MOVUPS       X14, t1-32(SP)
        MOVUPS       t0-16(SP), X13
        MOVUPS       X13, t2-48(SP)
        MOVUPS       t1-32(SP), X12
        MOVUPS       X12, t3-64(SP)
        MOVUPS       t3-64(SP), X11
        MOVUPS       t2-48(SP), X10
        MULPS        X11, X10
        MOVUPS       X10, t4-80(SP)
        MOVUPS       t4-80(SP), X9
        MOVUPS       X9, ret0+32(FP)
        RET

TEXT ·divf32x4(SB),$88-48
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
        MOVUPS       X15, t0-16(SP)
        MOVUPS       y+16(FP), X14
        MOVUPS       X14, t1-32(SP)
        MOVUPS       t0-16(SP), X13
        MOVUPS       X13, t2-48(SP)
        MOVUPS       t1-32(SP), X12
        MOVUPS       X12, t3-64(SP)
        MOVUPS       t3-64(SP), X11
        MOVUPS       t2-48(SP), X10
        DIVPS        X11, X10
        MOVUPS       X10, t4-80(SP)
        MOVUPS       t4-80(SP), X9
        MOVUPS       X9, ret0+32(FP)
        RET

