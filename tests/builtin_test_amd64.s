// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·lent0s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVQ         x+0(FP), R15
        MOVQ         R15, R14
        MOVQ         $1, R13
        MOVQ         R13, ret0+8(FP)
        RET

TEXT ·lent1s(SB),$24-24
        MOVQ         $0, ret0+16(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVQ         x+0(FP), R15
        MOVQ         R15, R14
        MOVQ         x+8(FP), R13
        MOVQ         R13, R12
        MOVQ         $2, R11
        MOVQ         R11, ret0+16(FP)
        RET

TEXT ·lent2s(SB),$16-32
        MOVQ         $0, ret0+24(FP)
block0:
        MOVQ         x+8(FP), R15
        MOVQ         R15, R14
        MOVQ         R14, ret0+24(FP)
        RET

