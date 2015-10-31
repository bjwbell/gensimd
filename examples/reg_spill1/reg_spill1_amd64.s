// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·regspill1(SB),$40-12
        MOVL         $0, ret0+8(FP)
        MOVL         $0, t0-4(SP)
        MOVL         $0, t2-12(SP)
        MOVL         $0, t4-20(SP)
        MOVL         $0, t7-32(SP)
        MOVL         $0, t8-36(SP)
        MOVL         $0, t1-8(SP)
        MOVL         $0, t3-16(SP)
        MOVL         $0, t5-24(SP)
        MOVL         $0, t6-28(SP)
block0:
        MOVL         x+0(FP), R14
        MOVL         x+0(FP), R13
        MOVL         R14, R15
        ADDL         R13, R15
        MOVL         y+4(FP), R13
        MOVL         y+4(FP), R12
        MOVL         R13, R14
        ADDL         R12, R14
        MOVL         R15, R13
        MOVL         R13, AX
        IMULL        R15
        MOVL         AX, R13
        MOVL         R14, R12
        MOVL         R12, AX
        IMULL        R14
        MOVL         AX, R12
        MOVL         R13, R11
        ADDL         R12, R11
        MOVL         R15, R10
        SUBL         R14, R10
        MOVL         $2, R8
        MOVL         R8, R9
        MOVL         R9, AX
        IMULL        R10
        MOVL         AX, R9
        MOVL         R9, R8
        ADDL         R14, R8
        MOVL         R8, t7-32(SP)
        MOVL         R9, t6-28(SP)
        MOVL         t7-32(SP), R9
        MOVL         R11, R8
        ADDL         R9, R8
        MOVL         R8, ret0+8(FP)
        RET

