// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·distsq(SB),$544-52
        MOVL         $0, ret0+48(FP)
        MOVQ         $0, t14-16(SP)
        MOVQ         $0, t14-8(SP)
        MOVQ         $0, t20-32(SP)
        MOVQ         $0, t20-24(SP)
        MOVQ         $0, t26-48(SP)
        MOVQ         $0, t26-40(SP)
        MOVQ         $0, t30-64(SP)
        MOVQ         $0, t30-56(SP)
        MOVQ         $0, t34-80(SP)
        MOVQ         $0, t34-72(SP)
block0:
        MOVQ         x+8(FP), R15
        MOVQ         R15, t0-88(SP)
        MOVQ         y+32(FP), R14
        MOVQ         R14, t1-96(SP)
        MOVQ         t0-88(SP), R12
        MOVQ         t1-96(SP), R11
        CMPQ         R12, R11
        SETNE        R13
        MOVB         R13, t2-97(SP)
        MOVB         t2-97(SP), R15
        CMPB         R15, $0
        JEQ          block2
        JMP          block1
block1:
        MOVL         $-1, R15
        MOVL         R15, ret0+48(FP)
        RET
block2:
        MOVQ         x+8(FP), R14
        MOVQ         R14, t3-105(SP)
        MOVL         $2147483647, R13
        MOVL         R13, t4-109(SP)
        MOVQ         $0, R12
        MOVQ         R12, t5-117(SP)
        JMP block5
block3:
        MOVL         t4-109(SP), R15
        MOVL         R15, t9-121(SP)
        MOVQ         $0, R14
        MOVQ         R14, t10-129(SP)
        JMP block8
block4:
        MOVL         t4-109(SP), R15
        MOVL         R15, ret0+48(FP)
        RET
block5:
        MOVQ         t5-117(SP), R13
        MOVQ         t3-105(SP), R12
        CMPQ         R13, R12
        SETLT        R14
        MOVB         R14, t6-130(SP)
        MOVB         t6-130(SP), R15
        CMPB         R15, $0
        JEQ          block4
        JMP          block3
block6:
        MOVQ         t5-117(SP), R14
        MOVQ         t10-129(SP), R13
        CMPQ         R14, R13
        SETEQ        R15
        MOVB         R15, t7-131(SP)
        MOVB         t7-131(SP), R15
        CMPB         R15, $0
        JEQ          block10
        MOVL         t9-121(SP), R15
        MOVL         R15, t12-135(SP)
        JMP          block9
block7:
        MOVQ         t5-117(SP), R14
        MOVQ         $1, R13
        MOVQ         R14, R15
        ADDQ         R13, R15
        MOVL         t9-121(SP), R12
        MOVL         R12, t4-109(SP)
        MOVQ         R15, R14
        MOVQ         R14, t5-117(SP)
        MOVQ         R15, t8-143(SP)
        JMP block5
block8:
        MOVQ         t10-129(SP), R14
        MOVQ         t3-105(SP), R13
        CMPQ         R14, R13
        SETLT        R15
        MOVB         R15, t11-144(SP)
        MOVB         t11-144(SP), R15
        CMPB         R15, $0
        JEQ          block7
        JMP          block6
block9:
        MOVQ         t10-129(SP), R14
        MOVQ         $1, R13
        MOVQ         R14, R15
        ADDQ         R13, R15
        MOVL         t12-135(SP), R12
        MOVL         R12, t9-121(SP)
        MOVQ         R15, R14
        MOVQ         R14, t10-129(SP)
        MOVQ         R15, t13-152(SP)
        JMP block8
block10:
        MOVQ         t10-129(SP), R14
        IMUL3Q       $16, R14, R14
        MOVQ         x+0(FP), R15
        ADDQ         R14, R15
        MOVQ         R15, R14
        MOVUPS       (R14), X15
        MOVUPS       X15, t16-176(SP)
        MOVQ         t5-117(SP), R13
        IMUL3Q       $16, R13, R13
        MOVQ         x+0(FP), R14
        ADDQ         R13, R14
        MOVQ         R14, R13
        MOVUPS       (R13), X15
        MOVUPS       X15, t18-200(SP)
        MOVOU        t18-200(SP), X15
        MOVOU        t16-176(SP), X14
        PSUBL        X15, X14
        MOVOU        X14, t19-216(SP)
        MOVOU        t19-216(SP), X13
        MOVOU        X13, t14-16(SP)
        MOVQ         t10-129(SP), R12
        IMUL3Q       $16, R12, R12
        MOVQ         y+24(FP), R13
        ADDQ         R12, R13
        MOVQ         R13, R12
        MOVUPS       (R12), X12
        MOVUPS       X12, t22-240(SP)
        MOVQ         t5-117(SP), R11
        IMUL3Q       $16, R11, R11
        MOVQ         y+24(FP), R12
        ADDQ         R11, R12
        MOVQ         R12, R11
        MOVUPS       (R11), X12
        MOVUPS       X12, t24-264(SP)
        MOVOU        t24-264(SP), X12
        MOVOU        t22-240(SP), X11
        PSUBL        X12, X11
        MOVOU        X11, t25-280(SP)
        MOVOU        t25-280(SP), X10
        MOVOU        X10, t20-32(SP)
        MOVOU        t14-16(SP), X9
        MOVOU        X9, t27-296(SP)
        MOVOU        X9, t28-312(SP)
        MOVOU        t27-296(SP), X7
        MOVOU        t28-312(SP), X6
        MOVO         X6, X8
        PMULULQ      X7, X8
        PSRLO        $4, X7
        PSRLO        $4, X6
        MOVO         X6, X5
        PMULULQ      X7, X5
        PSHUFD       $8, X8, X4
        PSHUFD       $8, X5, X3
        PUNPCKLLQ    X3, X4
        MOVOU        X4, t26-48(SP)
        MOVOU        t20-32(SP), X8
        MOVOU        X8, t31-344(SP)
        MOVOU        X8, t32-360(SP)
        MOVOU        t31-344(SP), X6
        MOVOU        t32-360(SP), X5
        MOVO         X5, X7
        PMULULQ      X6, X7
        PSRLO        $4, X6
        PSRLO        $4, X5
        MOVO         X5, X3
        PMULULQ      X6, X3
        PSHUFD       $8, X7, X2
        PSHUFD       $8, X3, X1
        PUNPCKLLQ    X1, X2
        MOVOU        X2, t30-64(SP)
        MOVOU        t26-48(SP), X7
        MOVOU        X7, t35-392(SP)
        MOVOU        t30-64(SP), X6
        MOVOU        X6, t36-408(SP)
        MOVOU        t36-408(SP), X5
        MOVOU        t35-392(SP), X3
        PADDL        X5, X3
        MOVOU        X3, t37-424(SP)
        MOVOU        t37-424(SP), X1
        MOVOU        X1, t34-80(SP)
        MOVQ         $0, R10
        IMUL3Q       $4, R10, R10
        LEAQ         t34-80(SP), R11
        ADDQ         R10, R11
        MOVL         (R11), R10
        MOVL         R10, t39-436(SP)
        MOVL         t39-436(SP), R9
        MOVL         t9-121(SP), R8
        CMPL         R9, R8
        SETLT        R10
        MOVL         R8, t43-441(SP)
        MOVB         R10, t40-437(SP)
        MOVQ         R11, t38-432(SP)
        MOVQ         R12, t23-248(SP)
        MOVQ         R13, t21-224(SP)
        MOVQ         R14, t17-184(SP)
        MOVQ         R15, t15-160(SP)
        MOVOU        X2, t33-376(SP)
        MOVOU        X3, t35-392(SP)
        MOVOU        X4, t29-328(SP)
        MOVOU        X11, t22-240(SP)
        MOVOU        X14, t16-176(SP)
        MOVB         t40-437(SP), R15
        CMPB         R15, $0
        JEQ          block12
        JMP          block11
block11:
        MOVQ         $0, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t42-453(SP)
        MOVL         t42-453(SP), R14
        MOVL         R14, t43-441(SP)
        MOVQ         R15, t41-449(SP)
        JMP block12
block12:
        MOVQ         $1, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t45-465(SP)
        MOVL         t45-465(SP), R13
        MOVL         t43-441(SP), R12
        CMPL         R13, R12
        SETLT        R14
        MOVL         R12, t49-470(SP)
        MOVB         R14, t46-466(SP)
        MOVQ         R15, t44-461(SP)
        MOVB         t46-466(SP), R15
        CMPB         R15, $0
        JEQ          block14
        JMP          block13
block13:
        MOVQ         $1, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t48-482(SP)
        MOVL         t48-482(SP), R14
        MOVL         R14, t49-470(SP)
        MOVQ         R15, t47-478(SP)
        JMP block14
block14:
        MOVQ         $2, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t51-494(SP)
        MOVL         t51-494(SP), R13
        MOVL         t49-470(SP), R12
        CMPL         R13, R12
        SETLT        R14
        MOVL         R12, t55-499(SP)
        MOVB         R14, t52-495(SP)
        MOVQ         R15, t50-490(SP)
        MOVB         t52-495(SP), R15
        CMPB         R15, $0
        JEQ          block16
        JMP          block15
block15:
        MOVQ         $2, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t54-511(SP)
        MOVL         t54-511(SP), R14
        MOVL         R14, t55-499(SP)
        MOVQ         R15, t53-507(SP)
        JMP block16
block16:
        MOVQ         $3, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t57-523(SP)
        MOVL         t57-523(SP), R13
        MOVL         t55-499(SP), R12
        CMPL         R13, R12
        SETLT        R14
        MOVL         R12, t12-135(SP)
        MOVB         R14, t58-524(SP)
        MOVQ         R15, t56-519(SP)
        MOVB         t58-524(SP), R15
        CMPB         R15, $0
        JEQ          block9
        JMP          block17
block17:
        MOVQ         $3, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVL         (R15), R14
        MOVL         R14, t60-536(SP)
        MOVL         t60-536(SP), R14
        MOVL         R14, t12-135(SP)
        MOVQ         R15, t59-532(SP)
        JMP block9

