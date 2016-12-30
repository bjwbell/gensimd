// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·distsq(SB),$384-52
        MOVL         $0, ret0+48(FP)
        MOVQ         $0, t26-16(SP)
        MOVQ         $0, t26-8(SP)
block0:
        MOVQ         x+8(FP), R15
        MOVQ         R15, R14
        MOVQ         y+32(FP), R13
        MOVQ         R13, R12
        CMPQ         R14, R12
        SETNE        R11
        MOVB         R11, t2-33(SP)
        CMPB         R11, $0
        JEQ          block2
        JMP          block1
block1:
        MOVL         $-1, R15
        MOVL         R15, ret0+48(FP)
        RET
block2:
        MOVQ         x+8(FP), R14
        MOVQ         R14, R13
        MOVL         $2147483647, R12
        MOVL         R12, t4-45(SP)
        MOVQ         $0, R11
        MOVQ         R11, t5-53(SP)
        MOVQ         R13, t3-41(SP)
        JMP block5
block3:
        MOVLQZX      t4-45(SP), R15
        MOVL         R15, t9-57(SP)
        MOVQ         $0, R14
        MOVQ         R14, t10-65(SP)
        JMP block8
block4:
        MOVLQZX      t4-45(SP), R15
        MOVL         R15, ret0+48(FP)
        RET
block5:
        MOVQ         t5-53(SP), R14
        MOVQ         t3-41(SP), R13
        CMPQ         R14, R13
        SETLT        R15
        MOVB         R15, t6-66(SP)
        CMPB         R15, $0
        JEQ          block4
        JMP          block3
block6:
        MOVQ         t5-53(SP), R14
        MOVQ         t10-65(SP), R13
        CMPQ         R14, R13
        SETEQ        R15
        MOVB         R15, t7-67(SP)
        CMPB         R15, $0
        JEQ          block10
        MOVLQZX      t9-57(SP), R15
        MOVL         R15, t12-71(SP)
        JMP          block9
block7:
        MOVQ         t5-53(SP), R14
        MOVQ         $1, R13
        MOVQ         R14, R15
        ADDQ         R13, R15
        MOVLQZX      t9-57(SP), R12
        MOVL         R12, t4-45(SP)
        MOVQ         R15, t5-53(SP)
        MOVQ         R15, t8-79(SP)
        JMP block5
block8:
        MOVQ         t10-65(SP), R14
        MOVQ         t3-41(SP), R13
        CMPQ         R14, R13
        SETLT        R15
        MOVB         R15, t11-80(SP)
        CMPB         R15, $0
        JEQ          block7
        JMP          block6
block9:
        MOVQ         t10-65(SP), R14
        MOVQ         $1, R13
        MOVQ         R14, R15
        ADDQ         R13, R15
        MOVLQZX      t12-71(SP), R12
        MOVL         R12, t9-57(SP)
        MOVQ         R15, t10-65(SP)
        MOVQ         R15, t13-88(SP)
        JMP block8
block10:
        MOVQ         t10-65(SP), R14
        IMUL3Q       $16, R14, R14
        MOVQ         x+0(FP), R15
        ADDQ         R14, R15
        MOVQ         R15, R14
        MOVUPS       (R14), X15
        MOVUPS       X15, t15-112(SP)
        MOVQ         t5-53(SP), R13
        IMUL3Q       $16, R13, R13
        MOVQ         x+0(FP), R14
        ADDQ         R13, R14
        MOVQ         R14, R13
        MOVUPS       (R13), X15
        MOVUPS       X15, t17-136(SP)
        MOVOU        t17-136(SP), X15
        MOVOU        t15-112(SP), X14
        PSUBL        X15, X14
        MOVQ         t10-65(SP), R12
        IMUL3Q       $16, R12, R12
        MOVQ         y+24(FP), R13
        ADDQ         R12, R13
        MOVQ         R13, R12
        MOVUPS       (R12), X13
        MOVUPS       X13, t20-176(SP)
        MOVQ         t5-53(SP), R11
        IMUL3Q       $16, R11, R11
        MOVQ         y+24(FP), R12
        ADDQ         R11, R12
        MOVQ         R12, R11
        MOVUPS       (R11), X13
        MOVUPS       X13, t22-200(SP)
        MOVOU        t22-200(SP), X13
        MOVOU        t20-176(SP), X12
        PSUBL        X13, X12
        MOVO         X14, X11
        PMULULQ      X14, X11
        MOVOU        X14, t18-152(SP)
        PSRLO        $4, X14
        MOVO         X14, X10
        PMULULQ      X14, X10
        PSHUFD       $8, X11, X9
        PSHUFD       $8, X10, X8
        PUNPCKLLQ    X8, X9
        MOVO         X12, X14
        PMULULQ      X12, X14
        MOVOU        X12, t23-216(SP)
        PSRLO        $4, X12
        MOVO         X12, X11
        PMULULQ      X12, X11
        PSHUFD       $8, X14, X10
        PSHUFD       $8, X11, X8
        PUNPCKLLQ    X8, X10
        MOVOU        X9, t24-232(SP)
        PADDL        X10, X9
        MOVO         X9, X14
        MOVOU        X14, t26-16(SP)
        MOVQ         $0, R10
        IMUL3Q       $4, R10, R10
        LEAQ         t26-16(SP), R11
        ADDQ         R10, R11
        MOVL         (R11), R10
        MOVL         R10, t29-276(SP)
        MOVLQZX      t29-276(SP), R9
        MOVLQZX      t9-57(SP), R8
        CMPL         R9, R8
        SETLT        R10
        MOVL         R8, t33-281(SP)
        MOVB         R10, t30-277(SP)
        CMPB         R10, $0
        JEQ          block12
        JMP          block11
block11:
        MOVQ         $0, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t26-16(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t32-293(SP)
        MOVLQZX      t32-293(SP), R14
        MOVL         R14, t33-281(SP)
        JMP block12
block12:
        MOVQ         $1, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t26-16(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t35-305(SP)
        MOVLQZX      t35-305(SP), R13
        MOVLQZX      t33-281(SP), R12
        CMPL         R13, R12
        SETLT        R14
        MOVL         R12, t39-310(SP)
        MOVB         R14, t36-306(SP)
        CMPB         R14, $0
        JEQ          block14
        JMP          block13
block13:
        MOVQ         $1, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t26-16(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t38-322(SP)
        MOVLQZX      t38-322(SP), R14
        MOVL         R14, t39-310(SP)
        JMP block14
block14:
        MOVQ         $2, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t26-16(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t41-334(SP)
        MOVLQZX      t41-334(SP), R13
        MOVLQZX      t39-310(SP), R12
        CMPL         R13, R12
        SETLT        R14
        MOVL         R12, t45-339(SP)
        MOVB         R14, t42-335(SP)
        CMPB         R14, $0
        JEQ          block16
        JMP          block15
block15:
        MOVQ         $2, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t26-16(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t44-351(SP)
        MOVLQZX      t44-351(SP), R14
        MOVL         R14, t45-339(SP)
        JMP block16
block16:
        MOVQ         $3, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t26-16(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t47-363(SP)
        MOVLQZX      t47-363(SP), R13
        MOVLQZX      t45-339(SP), R12
        CMPL         R13, R12
        SETLT        R14
        MOVL         R12, t12-71(SP)
        MOVB         R14, t48-364(SP)
        CMPB         R14, $0
        JEQ          block9
        JMP          block17
block17:
        MOVQ         $3, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t26-16(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t50-376(SP)
        MOVLQZX      t50-376(SP), R14
        MOVL         R14, t12-71(SP)
        JMP block9

