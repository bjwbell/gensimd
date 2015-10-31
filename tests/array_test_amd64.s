// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·arrayt0s(SB),$32-16
        MOVQ         $0, ret0+8(FP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-16(SP)
        MOVQ         $0, t2-24(SP)
block0:
        MOVQ         x+0(FP), R15
        MOVQ         R15, t0-8(SP)
        MOVQ         $0, R14
        IMUL3Q       $8, R14, R14
        LEAQ         t0-8(SP), R15
        ADDQ         R14, R15
        MOVQ         R15, R13
        MOVQ         (R13), R14
        MOVQ         R14, t2-24(SP)
        MOVQ         t2-24(SP), R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·arrayt1s(SB),$40-24
        MOVQ         $0, ret0+16(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-32(SP)
block0:
        MOVQ         x+0(FP), R15
        MOVQ         x+8(FP), R14
        MOVQ         R15, t0-16(SP)
        MOVQ         R14, t0-8(SP)
        MOVQ         $1, R14
        IMUL3Q       $8, R14, R14
        LEAQ         t0-16(SP), R15
        ADDQ         R14, R15
        MOVQ         R15, R13
        MOVQ         (R13), R14
        MOVQ         R14, t2-32(SP)
        MOVQ         t2-32(SP), R14
        MOVQ         R14, ret0+16(FP)
        RET

TEXT ·arrayt2s(SB),$96-32
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, t0-24(SP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t5-64(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t4-56(SP)
        MOVQ         $0, t6-72(SP)
        MOVQ         $0, t7-80(SP)
        MOVQ         $0, t8-88(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-48(SP)
block0:
        MOVQ         x+0(FP), R15
        MOVQ         x+8(FP), R14
        MOVQ         x+16(FP), R13
        MOVQ         R15, t0-24(SP)
        MOVQ         R14, t0-16(SP)
        MOVQ         R13, t0-8(SP)
        MOVQ         $0, R14
        IMUL3Q       $8, R14, R14
        LEAQ         t0-24(SP), R15
        ADDQ         R14, R15
        MOVQ         R15, R13
        MOVQ         (R13), R14
        MOVQ         R14, t2-40(SP)
        MOVQ         $1, R13
        IMUL3Q       $8, R13, R13
        LEAQ         t0-24(SP), R14
        ADDQ         R13, R14
        MOVQ         R14, R12
        MOVQ         (R12), R13
        MOVQ         R13, t4-56(SP)
        MOVQ         t2-40(SP), R12
        MOVQ         t4-56(SP), R11
        MOVQ         R12, R13
        ADDQ         R11, R13
        MOVQ         $2, R11
        IMUL3Q       $8, R11, R11
        LEAQ         t0-24(SP), R12
        ADDQ         R11, R12
        MOVQ         R12, R10
        MOVQ         (R10), R11
        MOVQ         R11, t7-80(SP)
        MOVQ         t7-80(SP), R10
        MOVQ         R13, R11
        ADDQ         R10, R11
        MOVQ         R11, ret0+24(FP)
        RET

