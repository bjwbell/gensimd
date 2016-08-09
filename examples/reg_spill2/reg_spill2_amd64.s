// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·regspill2(SB),$1184-52
        MOVL         $0, ret0+48(FP)
        MOVQ         $0, t3-16(SP)
        MOVQ         $0, t3-8(SP)
        MOVQ         $0, t9-32(SP)
        MOVQ         $0, t9-24(SP)
        MOVQ         $0, t21-48(SP)
        MOVQ         $0, t21-40(SP)
        MOVQ         $0, t23-64(SP)
        MOVQ         $0, t23-56(SP)
        MOVQ         $0, t27-80(SP)
        MOVQ         $0, t27-72(SP)
        MOVQ         $0, t35-96(SP)
        MOVQ         $0, t35-88(SP)
        MOVQ         $0, t37-112(SP)
        MOVQ         $0, t37-104(SP)
        MOVQ         $0, t41-128(SP)
        MOVQ         $0, t41-120(SP)
        MOVQ         $0, t49-144(SP)
        MOVQ         $0, t49-136(SP)
block0:
        MOVQ         x+8(FP), R15
        MOVQ         R15, R14
        MOVQ         y+32(FP), R13
        MOVQ         R13, R12
        CMPQ         R14, R12
        SETNE        R11
        MOVB         R11, t2-161(SP)
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
        MOVUPS       X15, t5-185(SP)
        MOVQ         $0, R12
        IMUL3Q       $16, R12, R12
        MOVQ         x+0(FP), R13
        ADDQ         R12, R13
        MOVQ         R13, R12
        MOVUPS       (R12), X15
        MOVUPS       X15, t7-209(SP)
        MOVOU        t7-209(SP), X15
        MOVOU        t5-185(SP), X14
        PSUBL        X15, X14
        MOVO         X14, X13
        MOVQ         $1, R11
        IMUL3Q       $16, R11, R11
        MOVQ         y+24(FP), R12
        ADDQ         R11, R12
        MOVQ         R12, R11
        MOVUPS       (R11), X12
        MOVUPS       X12, t11-249(SP)
        MOVQ         $0, R10
        IMUL3Q       $16, R10, R10
        MOVQ         y+24(FP), R11
        ADDQ         R10, R11
        MOVQ         R11, R10
        MOVUPS       (R10), X12
        MOVUPS       X12, t13-273(SP)
        MOVOU        t13-273(SP), X12
        MOVOU        t11-249(SP), X11
        PSUBL        X12, X11
        MOVO         X11, X10
        MOVO         X13, X9
        MOVO         X13, X8
        MOVO         X8, X7
        PMULULQ      X9, X7
        MOVOU        X9, t15-305(SP)
        PSRLO        $4, X9
        MOVOU        X8, t16-321(SP)
        PSRLO        $4, X8
        MOVO         X8, X6
        PMULULQ      X9, X6
        PSHUFD       $8, X7, X5
        PSHUFD       $8, X6, X4
        PUNPCKLLQ    X4, X5
        MOVO         X10, X9
        MOVO         X10, X8
        MOVO         X8, X7
        PMULULQ      X9, X7
        MOVOU        X9, t18-353(SP)
        PSRLO        $4, X9
        MOVOU        X8, t19-369(SP)
        PSRLO        $4, X8
        MOVO         X8, X6
        PMULULQ      X9, X6
        PSHUFD       $8, X7, X4
        PSHUFD       $8, X6, X3
        PUNPCKLLQ    X3, X4
        MOVOU        X5, t17-337(SP)
        PADDL        X4, X5
        MOVO         X5, X9
        MOVO         X13, X8
        MOVO         X10, X7
        MOVOU        X8, t24-417(SP)
        PSUBL        X7, X8
        MOVO         X8, X6
        MOVOU        t17-337(SP), X3
        PSUBL        X4, X3
        MOVO         X3, X2
        MOVO         X6, X1
        MOVO         X6, X0
        MOVO         X0, X3
        PMULULQ      X1, X3
        MOVOU        X1, t29-481(SP)
        PSRLO        $4, X1
        MOVOU        X0, t30-497(SP)
        PSRLO        $4, X0
        MOVO         X0, X4
        PMULULQ      X1, X4
        PSHUFD       $8, X3, X5
        MOVOU        X6, t23-64(SP)
        PSHUFD       $8, X4, X6
        PUNPCKLLQ    X6, X5
        MOVO         X2, X6
        MOVO         X2, X4
        MOVO         X4, X3
        PMULULQ      X6, X3
        MOVOU        X6, t32-529(SP)
        PSRLO        $4, X6
        MOVOU        X4, t33-545(SP)
        PSRLO        $4, X4
        MOVO         X4, X1
        PMULULQ      X6, X1
        PSHUFD       $8, X3, X0
        MOVOU        X2, t27-80(SP)
        PSHUFD       $8, X1, X2
        PUNPCKLLQ    X2, X0
        MOVOU        X5, t31-513(SP)
        PADDL        X0, X5
        MOVO         X5, X6
        MOVOU        t23-64(SP), X4
        MOVO         X4, X3
        MOVOU        t27-80(SP), X2
        MOVO         X2, X1
        MOVOU        X3, t38-593(SP)
        PSUBL        X1, X3
        MOVOU        X3, t37-112(SP)
        MOVOU        t31-513(SP), X1
        PSUBL        X0, X1
        MOVOU        X1, t41-128(SP)
        MOVOU        t37-112(SP), X0
        MOVOU        X0, t43-657(SP)
        MOVOU        X0, t44-673(SP)
        MOVOU        t43-657(SP), X1
        MOVOU        t44-673(SP), X2
        MOVO         X2, X0
        PMULULQ      X1, X0
        PSRLO        $4, X1
        PSRLO        $4, X2
        MOVO         X2, X3
        PMULULQ      X1, X3
        PSHUFD       $8, X0, X4
        PSHUFD       $8, X3, X5
        PUNPCKLLQ    X5, X4
        MOVOU        t41-128(SP), X5
        MOVO         X5, X3
        MOVO         X5, X2
        MOVO         X2, X1
        PMULULQ      X3, X1
        MOVOU        X3, t46-705(SP)
        PSRLO        $4, X3
        MOVOU        X2, t47-721(SP)
        PSRLO        $4, X2
        MOVO         X2, X0
        PMULULQ      X3, X0
        MOVOU        X4, t45-689(SP)
        PSHUFD       $8, X1, X4
        PSHUFD       $8, X0, X5
        PUNPCKLLQ    X5, X4
        MOVOU        t45-689(SP), X5
        PADDL        X4, X5
        MOVO         X5, X3
        MOVOU        X13, t3-16(SP)
        MOVQ         $0, R9
        IMUL3Q       $4, R9, R9
        LEAQ         t3-16(SP), R10
        ADDQ         R9, R10
        MOVL         (R10), R9
        MOVL         R9, t52-765(SP)
        MOVQ         $1, R8
        IMUL3Q       $4, R8, R8
        LEAQ         t3-16(SP), R9
        ADDQ         R8, R9
        MOVL         (R9), R8
        MOVL         R8, t54-777(SP)
        MOVLQZX      t52-765(SP), R9
        MOVLQZX      t54-777(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, BX
        IMUL3Q       $4, BX, BX
        LEAQ         t3-16(SP), BP
        ADDQ         BX, BP
        MOVL         (BP), BX
        MOVL         BX, t57-793(SP)
        MOVL         R8, t55-781(SP)
        MOVLQZX      t55-781(SP), R9
        MOVLQZX      t57-793(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, DI
        IMUL3Q       $4, DI, DI
        LEAQ         t3-16(SP), BX
        ADDQ         DI, BX
        MOVL         (BX), DI
        MOVL         DI, t60-809(SP)
        MOVL         R8, t58-797(SP)
        MOVLQZX      t58-797(SP), R9
        MOVLQZX      t60-809(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t23-64(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t62-821(SP)
        MOVQ         t62-821(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t63-825(SP)
        MOVL         R8, t61-813(SP)
        MOVLQZX      t61-813(SP), R9
        MOVLQZX      t63-825(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t23-64(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t65-837(SP)
        MOVQ         t65-837(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t66-841(SP)
        MOVL         R8, t64-829(SP)
        MOVLQZX      t64-829(SP), R9
        MOVLQZX      t66-841(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t23-64(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t68-853(SP)
        MOVQ         t68-853(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t69-857(SP)
        MOVL         R8, t67-845(SP)
        MOVLQZX      t67-845(SP), R9
        MOVLQZX      t69-857(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t23-64(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t71-869(SP)
        MOVQ         t71-869(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t72-873(SP)
        MOVL         R8, t70-861(SP)
        MOVLQZX      t70-861(SP), R9
        MOVLQZX      t72-873(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t37-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t74-885(SP)
        MOVQ         t74-885(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t75-889(SP)
        MOVL         R8, t73-877(SP)
        MOVLQZX      t73-877(SP), R9
        MOVLQZX      t75-889(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t37-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t77-901(SP)
        MOVQ         t77-901(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t78-905(SP)
        MOVL         R8, t76-893(SP)
        MOVLQZX      t76-893(SP), R9
        MOVLQZX      t78-905(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t37-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t80-917(SP)
        MOVQ         t80-917(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t81-921(SP)
        MOVL         R8, t79-909(SP)
        MOVLQZX      t79-909(SP), R9
        MOVLQZX      t81-921(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t37-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t83-933(SP)
        MOVQ         t83-933(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t84-937(SP)
        MOVL         R8, t82-925(SP)
        MOVLQZX      t82-925(SP), R9
        MOVLQZX      t84-937(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X10, t9-32(SP)
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t9-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t86-949(SP)
        MOVQ         t86-949(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t87-953(SP)
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t9-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t88-961(SP)
        MOVQ         t88-961(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t89-965(SP)
        MOVL         R8, t85-941(SP)
        MOVLQZX      t87-953(SP), R9
        MOVLQZX      t89-965(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t9-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t91-977(SP)
        MOVQ         t91-977(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t92-981(SP)
        MOVL         R8, t90-969(SP)
        MOVLQZX      t90-969(SP), R9
        MOVLQZX      t92-981(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t9-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t94-993(SP)
        MOVQ         t94-993(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t95-997(SP)
        MOVL         R8, t93-985(SP)
        MOVLQZX      t93-985(SP), R9
        MOVLQZX      t95-997(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t27-80(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t97-1009(SP)
        MOVQ         t97-1009(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t98-1013(SP)
        MOVL         R8, t96-1001(SP)
        MOVLQZX      t96-1001(SP), R9
        MOVLQZX      t98-1013(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t27-80(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t100-1025(SP)
        MOVQ         t100-1025(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t101-1029(SP)
        MOVL         R8, t99-1017(SP)
        MOVLQZX      t99-1017(SP), R9
        MOVLQZX      t101-1029(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t27-80(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t103-1041(SP)
        MOVQ         t103-1041(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t104-1045(SP)
        MOVL         R8, t102-1033(SP)
        MOVLQZX      t102-1033(SP), R9
        MOVLQZX      t104-1045(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t27-80(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t106-1057(SP)
        MOVQ         t106-1057(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t107-1061(SP)
        MOVL         R8, t105-1049(SP)
        MOVLQZX      t105-1049(SP), R9
        MOVLQZX      t107-1061(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t41-128(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t109-1073(SP)
        MOVQ         t109-1073(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t110-1077(SP)
        MOVL         R8, t108-1065(SP)
        MOVLQZX      t108-1065(SP), R9
        MOVLQZX      t110-1077(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t41-128(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t112-1089(SP)
        MOVQ         t112-1089(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t113-1093(SP)
        MOVL         R8, t111-1081(SP)
        MOVLQZX      t111-1081(SP), R9
        MOVLQZX      t113-1093(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t41-128(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t115-1105(SP)
        MOVQ         t115-1105(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t116-1109(SP)
        MOVL         R8, t114-1097(SP)
        MOVLQZX      t114-1097(SP), R9
        MOVLQZX      t116-1109(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t41-128(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t118-1121(SP)
        MOVQ         t118-1121(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t119-1125(SP)
        MOVL         R8, t117-1113(SP)
        MOVLQZX      t117-1113(SP), R9
        MOVLQZX      t119-1125(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVL         R8, t120-1129(SP)
        MOVLQZX      t85-941(SP), R9
        MOVLQZX      t120-1129(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X9, t21-48(SP)
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t21-48(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t122-1141(SP)
        MOVQ         t122-1141(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t123-1145(SP)
        MOVL         R8, t121-1133(SP)
        MOVLQZX      t121-1133(SP), R9
        MOVLQZX      t123-1145(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X6, t35-96(SP)
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t35-96(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t125-1157(SP)
        MOVQ         t125-1157(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t126-1161(SP)
        MOVL         R8, t124-1149(SP)
        MOVLQZX      t124-1149(SP), R9
        MOVLQZX      t126-1161(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X3, t49-144(SP)
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t49-144(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t128-1173(SP)
        MOVQ         t128-1173(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t129-1177(SP)
        MOVL         R8, t127-1165(SP)
        MOVLQZX      t127-1165(SP), R9
        MOVLQZX      t129-1177(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVL         R8, ret0+48(FP)
        RET

