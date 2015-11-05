// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·addi32x4(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        PADDL        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·subi32x4(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        PSUBL        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·muli32x4(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVO         X10, X9
        PMULULQ      X11, X9
        MOVOU        X11, t2-48(SP)
        PSRLO        $4, X11
        MOVOU        X10, t3-64(SP)
        PSRLO        $4, X10
        MOVO         X10, X8
        PMULULQ      X11, X8
        PSHUFD       $8, X9, X7
        PSHUFD       $8, X8, X6
        PUNPCKLLQ    X6, X7
        MOVOU        X7, ret0+32(FP)
        RET

TEXT ·shli32x4(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVO         X14, X13
        MOVB         shift+16(FP), R15
        MOVQ         R15, X12
        PSLLL        X12, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shri32x4(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVO         X14, X13
        MOVB         shift+16(FP), R15
        MOVQ         R15, X12
        PSRAL        X12, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·addf32x4(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVO         X15, X14
        MOVUPS       y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        ADDPS        X10, X11
        MOVUPS       X11, ret0+32(FP)
        RET

TEXT ·subf32x4(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVO         X15, X14
        MOVUPS       y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        SUBPS        X10, X11
        MOVUPS       X11, ret0+32(FP)
        RET

TEXT ·mulf32x4(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVO         X15, X14
        MOVUPS       y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MULPS        X10, X11
        MOVUPS       X11, ret0+32(FP)
        RET

TEXT ·divf32x4(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVO         X15, X14
        MOVUPS       y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        DIVPS        X10, X11
        MOVUPS       X11, ret0+32(FP)
        RET

