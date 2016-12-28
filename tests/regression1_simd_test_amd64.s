// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·regression1Simds(SB),$464-52
        MOVL         $0, ret0+48(FP)
        MOVQ         $0, t3-16(SP)
        MOVQ         $0, t3-8(SP)
        MOVQ         $0, t9-32(SP)
        MOVQ         $0, t9-24(SP)
        MOVQ         $0, t15-48(SP)
        MOVQ         $0, t15-40(SP)
        MOVQ         $0, t19-64(SP)
        MOVQ         $0, t19-56(SP)
        MOVQ         $0, t23-80(SP)
        MOVQ         $0, t23-72(SP)
        MOVQ         $0, t27-96(SP)
        MOVQ         $0, t27-88(SP)
block0:
        MOVQ         x+8(FP), R15
        MOVQ         R15, R14
        MOVQ         y+32(FP), R13
        MOVQ         R13, R12
        CMPQ         R14, R12
        SETNE        R11
        MOVB         R11, t2-113(SP)
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
        MOVUPS       X15, t5-137(SP)
        MOVQ         $0, R12
        IMUL3Q       $16, R12, R12
        MOVQ         x+0(FP), R13
        ADDQ         R12, R13
        MOVQ         R13, R12
        MOVUPS       (R12), X15
        MOVUPS       X15, t7-161(SP)
        MOVOU        t7-161(SP), X15
        MOVOU        t5-137(SP), X14
        PSUBL        X15, X14
        MOVO         X14, X13
        MOVQ         $1, R11
        IMUL3Q       $16, R11, R11
        MOVQ         y+24(FP), R12
        ADDQ         R11, R12
        MOVQ         R12, R11
        MOVUPS       (R11), X12
        MOVUPS       X12, t11-201(SP)
        MOVQ         $0, R10
        IMUL3Q       $16, R10, R10
        MOVQ         y+24(FP), R11
        ADDQ         R10, R11
        MOVQ         R11, R10
        MOVUPS       (R10), X12
        MOVUPS       X12, t13-225(SP)
        MOVOU        t13-225(SP), X12
        MOVOU        t11-201(SP), X11
        PSUBL        X12, X11
        MOVO         X11, X10
        MOVO         X13, X9
        MOVO         X13, X8
        MOVO         X8, X7
        PMULULQ      X9, X7
        MOVOU        X9, t16-257(SP)
        PSRLO        $4, X9
        MOVOU        X8, t17-273(SP)
        PSRLO        $4, X8
        MOVO         X8, X6
        PMULULQ      X9, X6
        PSHUFD       $8, X7, X5
        PSHUFD       $8, X6, X4
        PUNPCKLLQ    X4, X5
        MOVO         X5, X9
        MOVO         X10, X8
        MOVO         X10, X7
        MOVO         X7, X6
        PMULULQ      X8, X6
        MOVOU        X8, t20-305(SP)
        PSRLO        $4, X8
        MOVOU        X7, t21-321(SP)
        PSRLO        $4, X7
        MOVO         X7, X4
        PMULULQ      X8, X4
        PSHUFD       $8, X6, X3
        PSHUFD       $8, X4, X2
        PUNPCKLLQ    X2, X3
        MOVO         X3, X8
        MOVO         X9, X7
        MOVO         X8, X6
        MOVOU        X7, t24-353(SP)
        PADDL        X6, X7
        MOVO         X7, X4
        MOVO         X9, X2
        MOVO         X8, X1
        MOVOU        X2, t28-401(SP)
        PSUBL        X1, X2
        MOVO         X2, X0
        MOVOU        X4, t23-80(SP)
        MOVQ         $0, R9
        IMUL3Q       $4, R9, R9
        LEAQ         t23-80(SP), R10
        ADDQ         R9, R10
        MOVL         (R10), R9
        MOVL         R9, t32-445(SP)
        MOVOU        X0, t27-96(SP)
        MOVQ         $2, R8
        IMUL3Q       $4, R8, R8
        LEAQ         t27-96(SP), R9
        ADDQ         R8, R9
        MOVL         (R9), R8
        MOVL         R8, t34-457(SP)
        MOVLQZX      t32-445(SP), R9
        MOVLQZX      t34-457(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVL         R8, ret0+48(FP)
        RET

