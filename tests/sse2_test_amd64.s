// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·addpd(SB),$88-48
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
        MOVUPD       x+0(FP), X15
        MOVUPD       X15, t0-16(SP)
        MOVUPD       y+16(FP), X14
        MOVUPD       X14, t1-32(SP)
        MOVUPD       t0-16(SP), X13
        MOVUPD       X13, t2-48(SP)
        MOVUPD       t1-32(SP), X12
        MOVUPD       X12, t3-64(SP)
        MOVUPD       t2-48(SP), X11
        MOVUPD       t3-64(SP), X10
        ADDPD        X11, X10
        MOVUPD       X10, t4-80(SP)
        MOVUPD       t4-80(SP), X9
        MOVUPD       X9, ret0+32(FP)
        RET

