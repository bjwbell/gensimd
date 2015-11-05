// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·regspill1(SB),$40-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R14
        MOVL         R14, R15
        ADDL         R14, R15
        MOVL         y+4(FP), R12
        MOVL         R12, R13
        ADDL         R12, R13
        MOVL         R15, R11
        MOVL         R11, AX
        IMULL        R15
        MOVL         AX, R11
        MOVL         R13, R10
        MOVL         R10, AX
        IMULL        R13
        MOVL         AX, R10
        MOVL         R11, R9
        ADDL         R10, R9
        MOVL         R15, R8
        SUBL         R13, R8
        MOVL         R8, t5-24(SP)
        MOVL         R9, t4-20(SP)
        MOVL         $2, R9
        MOVL         t5-24(SP), R10
        MOVL         R9, R8
        MOVL         R8, AX
        IMULL        R10
        MOVL         AX, R8
        MOVL         R8, t6-28(SP)
        MOVL         t6-28(SP), R9
        MOVL         R9, R8
        ADDL         R13, R8
        MOVL         R8, t7-32(SP)
        MOVL         t4-20(SP), R9
        MOVL         t7-32(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVL         R8, ret0+8(FP)
        RET

