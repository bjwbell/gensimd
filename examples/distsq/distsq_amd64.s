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
        MOVQ         $0, t37-424(SP)
        MOVQ         $0, t37-416(SP)
        MOVQ         $0, t25-280(SP)
        MOVQ         $0, t25-272(SP)
        MOVL         $0, t51-494(SP)
        MOVL         $0, t12-135(SP)
        MOVQ         $0, t28-312(SP)
        MOVQ         $0, t28-304(SP)
        MOVQ         $0, t13-152(SP)
        MOVQ         $0, t31-344(SP)
        MOVQ         $0, t31-336(SP)
        MOVQ         $0, t23-248(SP)
        MOVQ         $0, t27-296(SP)
        MOVQ         $0, t27-288(SP)
        MOVQ         $0, t44-461(SP)
        MOVQ         $0, t47-478(SP)
        MOVL         $0, t55-499(SP)
        MOVQ         $0, t19-216(SP)
        MOVQ         $0, t19-208(SP)
        MOVQ         $0, t35-392(SP)
        MOVQ         $0, t35-384(SP)
        MOVL         $0, t45-465(SP)
        MOVQ         $0, t53-507(SP)
        MOVQ         $0, t56-519(SP)
        MOVL         $0, t9-121(SP)
        MOVQ         $0, t8-143(SP)
        MOVQ         $0, t18-200(SP)
        MOVQ         $0, t18-192(SP)
        MOVQ         $0, t21-224(SP)
        MOVQ         $0, t29-328(SP)
        MOVQ         $0, t29-320(SP)
        MOVQ         $0, t41-449(SP)
        MOVQ         $0, t1-96(SP)
        MOVQ         $0, t59-532(SP)
        MOVL         $0, t43-441(SP)
        MOVQ         $0, t50-490(SP)
        MOVQ         $0, t22-240(SP)
        MOVQ         $0, t22-232(SP)
        MOVQ         $0, t33-376(SP)
        MOVQ         $0, t33-368(SP)
        MOVQ         $0, t36-408(SP)
        MOVQ         $0, t36-400(SP)
        MOVQ         $0, t10-129(SP)
        MOVB         $0, t7-131(SP)
        MOVB         $0, t40-437(SP)
        MOVL         $0, t48-482(SP)
        MOVQ         $0, t0-88(SP)
        MOVB         $0, t46-466(SP)
        MOVL         $0, t57-523(SP)
        MOVL         $0, t60-536(SP)
        MOVQ         $0, t3-105(SP)
        MOVQ         $0, t5-117(SP)
        MOVQ         $0, t15-160(SP)
        MOVQ         $0, t24-264(SP)
        MOVQ         $0, t24-256(SP)
        MOVQ         $0, t32-360(SP)
        MOVQ         $0, t32-352(SP)
        MOVB         $0, t52-495(SP)
        MOVB         $0, t2-97(SP)
        MOVL         $0, t4-109(SP)
        MOVB         $0, t11-144(SP)
        MOVQ         $0, t16-176(SP)
        MOVQ         $0, t16-168(SP)
        MOVQ         $0, t17-184(SP)
        MOVQ         $0, t38-432(SP)
        MOVL         $0, t39-436(SP)
        MOVL         $0, t42-453(SP)
        MOVL         $0, t49-470(SP)
        MOVL         $0, t54-511(SP)
        MOVB         $0, t58-524(SP)
        MOVB         $0, t6-130(SP)
block0:
        MOVQ         x+8(FP), R15
        MOVQ         y+32(FP), R14
        CMPQ         R15, R14
        SETNE        R13
        MOVB         R13, t2-97(SP)
        MOVQ         R14, t1-96(SP)
        MOVQ         R15, t0-88(SP)
        MOVB         t2-97(SP), R15
        CMPB         R15, $0
        JEQ          block2
        JMP          block1
block1:
        MOVL         $-1, R15
        MOVL         R15, ret0+48(FP)
        RET
block2:
        MOVQ         x+8(FP), R15
        MOVL         $2147483647, R14
        MOVQ         $0, R13
        MOVQ         R13, t5-117(SP)
        MOVL         R14, t4-109(SP)
        MOVQ         R15, t3-105(SP)
        JMP block5
block3:
        MOVL         t4-109(SP), R15
        MOVQ         $0, R14
        MOVQ         R14, t10-129(SP)
        MOVL         R15, t9-121(SP)
        JMP block8
block4:
        MOVL         t4-109(SP), R15
        MOVL         R15, ret0+48(FP)
        RET
block5:
        MOVQ         t5-117(SP), R14
        MOVQ         t3-105(SP), R13
        CMPQ         R14, R13
        SETLT        R15
        MOVB         R15, t6-130(SP)
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
        MOVL         t9-121(SP), R14
        MOVQ         R15, t5-117(SP)
        MOVL         R14, t4-109(SP)
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
        MOVL         t12-135(SP), R14
        MOVQ         R15, t10-129(SP)
        MOVL         R14, t9-121(SP)
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
        MOVOU        X14, t14-16(SP)
        MOVQ         t10-129(SP), R12
        IMUL3Q       $16, R12, R12
        MOVQ         y+24(FP), R13
        ADDQ         R12, R13
        MOVQ         R13, R12
        MOVUPS       (R12), X15
        MOVUPS       X15, t22-240(SP)
        MOVQ         t5-117(SP), R11
        IMUL3Q       $16, R11, R11
        MOVQ         y+24(FP), R12
        ADDQ         R11, R12
        MOVQ         R12, R11
        MOVUPS       (R11), X15
        MOVUPS       X15, t24-264(SP)
        MOVOU        t24-264(SP), X15
        MOVOU        t22-240(SP), X13
        PSUBL        X15, X13
        MOVOU        X13, t20-32(SP)
        MOVOU        t14-16(SP), X15
        MOVOU        t14-16(SP), X12
        MOVO         X12, X11
        PMULULQ      X15, X11
        MOVOU        X15, t27-296(SP)
        PSRLO        $4, X15
        MOVOU        X12, t28-312(SP)
        PSRLO        $4, X12
        MOVO         X12, X10
        PMULULQ      X15, X10
        PSHUFD       $8, X11, X9
        PSHUFD       $8, X10, X8
        PUNPCKLLQ    X8, X9
        MOVOU        X9, t26-48(SP)
        MOVOU        t20-32(SP), X15
        MOVOU        t20-32(SP), X12
        MOVO         X12, X11
        PMULULQ      X15, X11
        MOVOU        X15, t31-344(SP)
        PSRLO        $4, X15
        MOVOU        X12, t32-360(SP)
        PSRLO        $4, X12
        MOVO         X12, X10
        PMULULQ      X15, X10
        PSHUFD       $8, X11, X8
        PSHUFD       $8, X10, X7
        PUNPCKLLQ    X7, X8
        MOVOU        X8, t30-64(SP)
        MOVOU        t26-48(SP), X15
        MOVOU        t30-64(SP), X12
        MOVOU        X15, t35-392(SP)
        PADDL        X12, X15
        MOVOU        X15, t34-80(SP)
        MOVQ         $0, R10
        IMUL3Q       $4, R10, R10
        LEAQ         t34-80(SP), R11
        ADDQ         R10, R11
        MOVQ         R11, R9
        MOVL         (R9), R10
        MOVL         R10, t39-436(SP)
        MOVL         t39-436(SP), R9
        MOVL         t9-121(SP), R8
        CMPL         R9, R8
        SETLT        R10
        MOVL         t9-121(SP), R9
        MOVL         R9, t43-441(SP)
        MOVB         R10, t40-437(SP)
        MOVOU        X8, t33-376(SP)
        MOVOU        X9, t29-328(SP)
        MOVOU        X12, t36-408(SP)
        MOVOU        X13, t25-280(SP)
        MOVOU        X14, t19-216(SP)
        MOVOU        X15, t37-424(SP)
        MOVB         t40-437(SP), R15
        CMPB         R15, $0
        JEQ          block12
        JMP          block11
block11:
        MOVQ         $0, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVQ         R15, R13
        MOVL         (R13), R14
        MOVL         R14, t42-453(SP)
        MOVL         t42-453(SP), R14
        MOVL         R14, t43-441(SP)
        JMP block12
block12:
        MOVQ         $1, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVQ         R15, R13
        MOVL         (R13), R14
        MOVL         R14, t45-465(SP)
        MOVL         t45-465(SP), R13
        MOVL         t43-441(SP), R12
        CMPL         R13, R12
        SETLT        R14
        MOVL         t43-441(SP), R13
        MOVL         R13, t49-470(SP)
        MOVB         R14, t46-466(SP)
        MOVB         t46-466(SP), R15
        CMPB         R15, $0
        JEQ          block14
        JMP          block13
block13:
        MOVQ         $1, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVQ         R15, R13
        MOVL         (R13), R14
        MOVL         R14, t48-482(SP)
        MOVL         t48-482(SP), R14
        MOVL         R14, t49-470(SP)
        JMP block14
block14:
        MOVQ         $2, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVQ         R15, R13
        MOVL         (R13), R14
        MOVL         R14, t51-494(SP)
        MOVL         t51-494(SP), R13
        MOVL         t49-470(SP), R12
        CMPL         R13, R12
        SETLT        R14
        MOVL         t49-470(SP), R13
        MOVL         R13, t55-499(SP)
        MOVB         R14, t52-495(SP)
        MOVB         t52-495(SP), R15
        CMPB         R15, $0
        JEQ          block16
        JMP          block15
block15:
        MOVQ         $2, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVQ         R15, R13
        MOVL         (R13), R14
        MOVL         R14, t54-511(SP)
        MOVL         t54-511(SP), R14
        MOVL         R14, t55-499(SP)
        JMP block16
block16:
        MOVQ         $3, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVQ         R15, R13
        MOVL         (R13), R14
        MOVL         R14, t57-523(SP)
        MOVL         t57-523(SP), R13
        MOVL         t55-499(SP), R12
        CMPL         R13, R12
        SETLT        R14
        MOVL         t55-499(SP), R13
        MOVL         R13, t12-135(SP)
        MOVB         R14, t58-524(SP)
        MOVB         t58-524(SP), R15
        CMPB         R15, $0
        JEQ          block9
        JMP          block17
block17:
        MOVQ         $3, R14
        IMUL3Q       $4, R14, R14
        LEAQ         t34-80(SP), R15
        ADDQ         R14, R15
        MOVQ         R15, R13
        MOVL         (R13), R14
        MOVL         R14, t60-536(SP)
        MOVL         t60-536(SP), R14
        MOVL         R14, t12-135(SP)
        JMP block9

