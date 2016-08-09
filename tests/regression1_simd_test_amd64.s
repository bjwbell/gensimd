// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·regression1Simds(SB),$272-52
        MOVL         $0, ret0+48(FP)
        MOVQ         $0, t15-16(SP)
        MOVQ         $0, t15-8(SP)
        MOVQ         $0, t17-32(SP)
        MOVQ         $0, t17-24(SP)
block0:
        MOVQ         x+8(FP), R15
        MOVQ         R15, R14
        MOVQ         y+32(FP), R13
        MOVQ         R13, R12
        CMPQ         R14, R12
        SETNE        R11
        MOVB         R11, t2-49(SP)
        CMPB         R11, $0
        JEQ          block2
        JMP          block1
block1:
        MOVL         $-1, R15
        MOVL         R15, ret0+48(FP)
        RET
block2:
        MOVQ         $1, R13
        IMUL3Q       $16, R13, R13
        MOVQ         x+0(FP), R14
        ADDQ         R13, R14
        MOVQ         R14, R13
        MOVUPS       (R13), X15
        MOVUPS       X15, t4-73(SP)
        MOVQ         $0, R12
        IMUL3Q       $16, R12, R12
        MOVQ         x+0(FP), R13
        ADDQ         R12, R13
        MOVQ         R13, R12
        MOVUPS       (R12), X15
        MOVUPS       X15, t6-97(SP)
        MOVOU        t6-97(SP), X15
        MOVOU        t4-73(SP), X14
        PSUBL        X15, X14
        MOVQ         $1, R11
        IMUL3Q       $16, R11, R11
        MOVQ         y+24(FP), R12
        ADDQ         R11, R12
        MOVQ         R12, R11
        MOVUPS       (R11), X13
        MOVUPS       X13, t9-137(SP)
        MOVQ         $0, R10
        IMUL3Q       $16, R10, R10
        MOVQ         y+24(FP), R11
        ADDQ         R10, R11
        MOVQ         R11, R10
        MOVUPS       (R10), X13
        MOVUPS       X13, t11-161(SP)
        MOVOU        t11-161(SP), X13
        MOVOU        t9-137(SP), X12
        PSUBL        X13, X12
        MOVO         X14, X11
        PMULULQ      X14, X11
        MOVOU        X14, t7-113(SP)
        PSRLO        $4, X14
        MOVO         X14, X10
        PMULULQ      X14, X10
        PSHUFD       $8, X11, X9
        PSHUFD       $8, X10, X8
        PUNPCKLLQ    X8, X9
        MOVO         X12, X14
        PMULULQ      X12, X14
        MOVOU        X12, t12-177(SP)
        PSRLO        $4, X12
        MOVO         X12, X11
        PMULULQ      X12, X11
        PSHUFD       $8, X14, X10
        PSHUFD       $8, X11, X8
        PUNPCKLLQ    X8, X10
        MOVOU        X9, t13-193(SP)
        PADDL        X10, X9
        MOVO         X9, X14
        MOVOU        t13-193(SP), X12
        PSUBL        X10, X12
        MOVO         X12, X11
        MOVOU        X14, t15-16(SP)
        MOVQ         $0, R9
        IMUL3Q       $4, R9, R9
        LEAQ         t15-16(SP), R10
        ADDQ         R9, R10
        MOVL         (R10), R9
        MOVL         R9, t20-253(SP)
        MOVOU        X11, t17-32(SP)
        MOVQ         $2, R8
        IMUL3Q       $4, R8, R8
        LEAQ         t17-32(SP), R9
        ADDQ         R8, R9
        MOVL         (R9), R8
        MOVL         R8, t22-265(SP)
        MOVLQZX      t20-253(SP), R9
        MOVLQZX      t22-265(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVL         R8, ret0+48(FP)
        RET

