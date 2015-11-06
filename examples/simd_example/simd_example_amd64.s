// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·addi32x4(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PADDL        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·subi32x4(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PSUBL        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·muli32x4(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        x+0(FP), X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X15
        PMULULQ      X14, X15
        PSRLO        $4, X14
        PSRLO        $4, X13
        MOVO         X13, X12
        PMULULQ      X14, X12
        PSHUFD       $8, X15, X11
        PSHUFD       $8, X12, X10
        PUNPCKLLQ    X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·shli32x4(SB),$24-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
block0:
        MOVB         shift+16(FP), R15
        MOVQ         R15, X15
        MOVOU        x+0(FP), X14
        PSLLL        X15, X14
        MOVOU        X14, ret0+24(FP)
        RET

TEXT ·shri32x4(SB),$24-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
block0:
        MOVB         shift+16(FP), R15
        MOVQ         R15, X15
        MOVOU        x+0(FP), X14
        PSRAL        X15, X14
        MOVOU        X14, ret0+24(FP)
        RET

TEXT ·addf32x4(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVUPS       y+16(FP), X15
        MOVUPS       x+0(FP), X14
        ADDPS        X15, X14
        MOVUPS       X14, ret0+32(FP)
        RET

TEXT ·subf32x4(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVUPS       y+16(FP), X15
        MOVUPS       x+0(FP), X14
        SUBPS        X15, X14
        MOVUPS       X14, ret0+32(FP)
        RET

TEXT ·mulf32x4(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVUPS       y+16(FP), X15
        MOVUPS       x+0(FP), X14
        MULPS        X15, X14
        MOVUPS       X14, ret0+32(FP)
        RET

TEXT ·divf32x4(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVUPS       y+16(FP), X15
        MOVUPS       x+0(FP), X14
        DIVPS        X15, X14
        MOVUPS       X14, ret0+32(FP)
        RET

