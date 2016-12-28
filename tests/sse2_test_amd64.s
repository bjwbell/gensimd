// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·addpd(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVUPD       x+0(FP), X15
        MOVO         X15, X14
        MOVUPD       y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        ADDPD        X11, X10
        MOVUPD       X10, ret0+32(FP)
        RET

