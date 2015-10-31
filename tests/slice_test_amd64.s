// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·slicet0s(SB),$24-32
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-16(SP)
block0:
        MOVQ         $0, R14
        IMUL3Q       $8, R14, R14
        MOVQ         x+0(FP), R15
        ADDQ         R14, R15
        MOVQ         R15, R13
        MOVQ         (R13), R14
        MOVQ         R14, t1-16(SP)
        MOVQ         t1-16(SP), R14
        MOVQ         R14, ret0+24(FP)
        RET

TEXT ·slicet1s(SB),$24-32
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, t1-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVQ         $1, R14
        IMUL3Q       $8, R14, R14
        MOVQ         x+0(FP), R15
        ADDQ         R14, R15
        MOVQ         R15, R13
        MOVQ         (R13), R14
        MOVQ         R14, t1-16(SP)
        MOVQ         t1-16(SP), R14
        MOVQ         R14, ret0+24(FP)
        RET

TEXT ·slicet2s(SB),$72-32
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, t6-56(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t2-24(SP)
        MOVQ         $0, t4-40(SP)
        MOVQ         $0, t5-48(SP)
        MOVQ         $0, t7-64(SP)
        MOVQ         $0, t1-16(SP)
        MOVQ         $0, t3-32(SP)
block0:
        MOVQ         $0, R14
        IMUL3Q       $8, R14, R14
        MOVQ         x+0(FP), R15
        ADDQ         R14, R15
        MOVQ         R15, R13
        MOVQ         (R13), R14
        MOVQ         R14, t1-16(SP)
        MOVQ         $1, R13
        IMUL3Q       $8, R13, R13
        MOVQ         x+0(FP), R14
        ADDQ         R13, R14
        MOVQ         R14, R12
        MOVQ         (R12), R13
        MOVQ         R13, t3-32(SP)
        MOVQ         t1-16(SP), R12
        MOVQ         t3-32(SP), R11
        MOVQ         R12, R13
        ADDQ         R11, R13
        MOVQ         $2, R11
        IMUL3Q       $8, R11, R11
        MOVQ         x+0(FP), R12
        ADDQ         R11, R12
        MOVQ         R12, R10
        MOVQ         (R10), R11
        MOVQ         R11, t6-56(SP)
        MOVQ         t6-56(SP), R10
        MOVQ         R13, R11
        ADDQ         R10, R11
        MOVQ         R11, ret0+24(FP)
        RET

