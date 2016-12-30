// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·addpd(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVUPD       x+0(FP), X15
        MOVUPD       y+16(FP), X14
        ADDPD        X15, X14
        MOVO         X14, X13
        MOVUPD       X13, ret0+32(FP)
        RET

