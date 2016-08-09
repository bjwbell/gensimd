// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·regspill3(SB),$1248-52
        MOVL         $0, ret0+48(FP)
        MOVQ         $0, t7-16(SP)
        MOVQ         $0, t7-8(SP)
        MOVQ         $0, t13-32(SP)
        MOVQ         $0, t13-24(SP)
        MOVQ         $0, t25-48(SP)
        MOVQ         $0, t25-40(SP)
        MOVQ         $0, t27-64(SP)
        MOVQ         $0, t27-56(SP)
        MOVQ         $0, t31-80(SP)
        MOVQ         $0, t31-72(SP)
        MOVQ         $0, t39-96(SP)
        MOVQ         $0, t39-88(SP)
        MOVQ         $0, t41-112(SP)
        MOVQ         $0, t41-104(SP)
        MOVQ         $0, t45-128(SP)
        MOVQ         $0, t45-120(SP)
        MOVQ         $0, t53-144(SP)
        MOVQ         $0, t53-136(SP)
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
        MOVL         $2147483647, R14
        MOVL         R14, t3-165(SP)
        MOVQ         $0, R13
        MOVQ         R13, t4-173(SP)
        JMP block5
block3:
        MOVLQZX      t3-165(SP), R15
        MOVL         R15, t138-177(SP)
        MOVQ         $0, R14
        MOVQ         R14, t139-185(SP)
        JMP block8
block4:
        MOVLQZX      t3-165(SP), R15
        MOVL         R15, ret0+48(FP)
        RET
block5:
        MOVQ         x+8(FP), R15
        MOVQ         R15, R14
        MOVQ         t4-173(SP), R12
        CMPQ         R12, R14
        SETLT        R13
        MOVB         R13, t6-194(SP)
        CMPB         R13, $0
        JEQ          block4
        JMP          block3
block6:
        MOVQ         t139-185(SP), R14
        IMUL3Q       $16, R14, R14
        MOVQ         x+0(FP), R15
        ADDQ         R14, R15
        MOVQ         R15, R14
        MOVUPS       (R14), X15
        MOVUPS       X15, t9-218(SP)
        MOVQ         t4-173(SP), R13
        IMUL3Q       $16, R13, R13
        MOVQ         x+0(FP), R14
        ADDQ         R13, R14
        MOVQ         R14, R13
        MOVUPS       (R13), X15
        MOVUPS       X15, t11-242(SP)
        MOVOU        t11-242(SP), X15
        MOVOU        t9-218(SP), X14
        PSUBL        X15, X14
        MOVO         X14, X13
        MOVQ         t139-185(SP), R12
        IMUL3Q       $16, R12, R12
        MOVQ         y+24(FP), R13
        ADDQ         R12, R13
        MOVQ         R13, R12
        MOVUPS       (R12), X12
        MOVUPS       X12, t15-282(SP)
        MOVQ         t4-173(SP), R11
        IMUL3Q       $16, R11, R11
        MOVQ         y+24(FP), R12
        ADDQ         R11, R12
        MOVQ         R12, R11
        MOVUPS       (R11), X12
        MOVUPS       X12, t17-306(SP)
        MOVOU        t17-306(SP), X12
        MOVOU        t15-282(SP), X11
        PSUBL        X12, X11
        MOVO         X11, X10
        MOVO         X13, X9
        MOVO         X13, X8
        MOVO         X8, X7
        PMULULQ      X9, X7
        MOVOU        X9, t19-338(SP)
        PSRLO        $4, X9
        MOVOU        X8, t20-354(SP)
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
        MOVOU        X9, t22-386(SP)
        PSRLO        $4, X9
        MOVOU        X8, t23-402(SP)
        PSRLO        $4, X8
        MOVO         X8, X6
        PMULULQ      X9, X6
        PSHUFD       $8, X7, X4
        PSHUFD       $8, X6, X3
        PUNPCKLLQ    X3, X4
        MOVOU        X5, t21-370(SP)
        PADDL        X4, X5
        MOVO         X5, X9
        MOVO         X13, X8
        MOVO         X10, X7
        MOVOU        X8, t28-450(SP)
        PSUBL        X7, X8
        MOVO         X8, X6
        MOVOU        t21-370(SP), X3
        PSUBL        X4, X3
        MOVO         X3, X2
        MOVO         X6, X1
        MOVO         X6, X0
        MOVO         X0, X3
        PMULULQ      X1, X3
        MOVOU        X1, t33-514(SP)
        PSRLO        $4, X1
        MOVOU        X0, t34-530(SP)
        PSRLO        $4, X0
        MOVO         X0, X4
        PMULULQ      X1, X4
        PSHUFD       $8, X3, X5
        MOVOU        X6, t27-64(SP)
        PSHUFD       $8, X4, X6
        PUNPCKLLQ    X6, X5
        MOVO         X2, X6
        MOVO         X2, X4
        MOVO         X4, X3
        PMULULQ      X6, X3
        MOVOU        X6, t36-562(SP)
        PSRLO        $4, X6
        MOVOU        X4, t37-578(SP)
        PSRLO        $4, X4
        MOVO         X4, X1
        PMULULQ      X6, X1
        PSHUFD       $8, X3, X0
        MOVOU        X2, t31-80(SP)
        PSHUFD       $8, X1, X2
        PUNPCKLLQ    X2, X0
        MOVOU        X5, t35-546(SP)
        PADDL        X0, X5
        MOVO         X5, X6
        MOVOU        t27-64(SP), X4
        MOVO         X4, X3
        MOVOU        t31-80(SP), X2
        MOVO         X2, X1
        MOVOU        X3, t42-626(SP)
        PSUBL        X1, X3
        MOVOU        X3, t41-112(SP)
        MOVOU        t35-546(SP), X1
        PSUBL        X0, X1
        MOVOU        X1, t45-128(SP)
        MOVOU        t41-112(SP), X0
        MOVOU        X0, t47-690(SP)
        MOVOU        X0, t48-706(SP)
        MOVOU        t47-690(SP), X1
        MOVOU        t48-706(SP), X2
        MOVO         X2, X0
        PMULULQ      X1, X0
        PSRLO        $4, X1
        PSRLO        $4, X2
        MOVO         X2, X3
        PMULULQ      X1, X3
        PSHUFD       $8, X0, X4
        PSHUFD       $8, X3, X5
        PUNPCKLLQ    X5, X4
        MOVOU        t45-128(SP), X5
        MOVO         X5, X3
        MOVO         X5, X2
        MOVO         X2, X1
        PMULULQ      X3, X1
        MOVOU        X3, t50-738(SP)
        PSRLO        $4, X3
        MOVOU        X2, t51-754(SP)
        PSRLO        $4, X2
        MOVO         X2, X0
        PMULULQ      X3, X0
        MOVOU        X4, t49-722(SP)
        PSHUFD       $8, X1, X4
        PSHUFD       $8, X0, X5
        PUNPCKLLQ    X5, X4
        MOVOU        t49-722(SP), X5
        PADDL        X4, X5
        MOVO         X5, X3
        MOVOU        X13, t7-16(SP)
        MOVQ         $0, R10
        IMUL3Q       $4, R10, R10
        LEAQ         t7-16(SP), R11
        ADDQ         R10, R11
        MOVL         (R11), R10
        MOVL         R10, t56-798(SP)
        MOVQ         $1, R9
        IMUL3Q       $4, R9, R9
        LEAQ         t7-16(SP), R10
        ADDQ         R9, R10
        MOVL         (R10), R9
        MOVL         R9, t58-810(SP)
        MOVLQZX      t56-798(SP), R8
        MOVLQZX      t58-810(SP), R10
        MOVL         R8, R9
        ADDL         R10, R9
        MOVQ         $2, BX
        IMUL3Q       $4, BX, BX
        LEAQ         t7-16(SP), BP
        ADDQ         BX, BP
        MOVL         (BP), BX
        MOVL         BX, t61-826(SP)
        MOVLQZX      t61-826(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, DI
        IMUL3Q       $4, DI, DI
        LEAQ         t7-16(SP), BX
        ADDQ         DI, BX
        MOVL         (BX), DI
        MOVL         DI, t64-842(SP)
        MOVL         R8, t62-830(SP)
        MOVLQZX      t62-830(SP), R9
        MOVLQZX      t64-842(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t27-64(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t66-854(SP)
        MOVQ         t66-854(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t67-858(SP)
        MOVL         R8, t65-846(SP)
        MOVLQZX      t65-846(SP), R9
        MOVLQZX      t67-858(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t27-64(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t69-870(SP)
        MOVQ         t69-870(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t70-874(SP)
        MOVL         R8, t68-862(SP)
        MOVLQZX      t68-862(SP), R9
        MOVLQZX      t70-874(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t27-64(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t72-886(SP)
        MOVQ         t72-886(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t73-890(SP)
        MOVL         R8, t71-878(SP)
        MOVLQZX      t71-878(SP), R9
        MOVLQZX      t73-890(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t27-64(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t75-902(SP)
        MOVQ         t75-902(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t76-906(SP)
        MOVL         R8, t74-894(SP)
        MOVLQZX      t74-894(SP), R9
        MOVLQZX      t76-906(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t41-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t78-918(SP)
        MOVQ         t78-918(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t79-922(SP)
        MOVL         R8, t77-910(SP)
        MOVLQZX      t77-910(SP), R9
        MOVLQZX      t79-922(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t41-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t81-934(SP)
        MOVQ         t81-934(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t82-938(SP)
        MOVL         R8, t80-926(SP)
        MOVLQZX      t80-926(SP), R9
        MOVLQZX      t82-938(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t41-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t84-950(SP)
        MOVQ         t84-950(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t85-954(SP)
        MOVL         R8, t83-942(SP)
        MOVLQZX      t83-942(SP), R9
        MOVLQZX      t85-954(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t41-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t87-966(SP)
        MOVQ         t87-966(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t88-970(SP)
        MOVL         R8, t86-958(SP)
        MOVLQZX      t86-958(SP), R9
        MOVLQZX      t88-970(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X10, t13-32(SP)
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t13-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t90-982(SP)
        MOVQ         t90-982(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t91-986(SP)
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t13-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t92-994(SP)
        MOVQ         t92-994(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t93-998(SP)
        MOVL         R8, t89-974(SP)
        MOVLQZX      t91-986(SP), R9
        MOVLQZX      t93-998(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t13-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t95-1010(SP)
        MOVQ         t95-1010(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t96-1014(SP)
        MOVL         R8, t94-1002(SP)
        MOVLQZX      t94-1002(SP), R9
        MOVLQZX      t96-1014(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t13-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t98-1026(SP)
        MOVQ         t98-1026(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t99-1030(SP)
        MOVL         R8, t97-1018(SP)
        MOVLQZX      t97-1018(SP), R9
        MOVLQZX      t99-1030(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t31-80(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t101-1042(SP)
        MOVQ         t101-1042(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t102-1046(SP)
        MOVL         R8, t100-1034(SP)
        MOVLQZX      t100-1034(SP), R9
        MOVLQZX      t102-1046(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t31-80(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t104-1058(SP)
        MOVQ         t104-1058(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t105-1062(SP)
        MOVL         R8, t103-1050(SP)
        MOVLQZX      t103-1050(SP), R9
        MOVLQZX      t105-1062(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t31-80(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t107-1074(SP)
        MOVQ         t107-1074(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t108-1078(SP)
        MOVL         R8, t106-1066(SP)
        MOVLQZX      t106-1066(SP), R9
        MOVLQZX      t108-1078(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t31-80(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t110-1090(SP)
        MOVQ         t110-1090(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t111-1094(SP)
        MOVL         R8, t109-1082(SP)
        MOVLQZX      t109-1082(SP), R9
        MOVLQZX      t111-1094(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t45-128(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t113-1106(SP)
        MOVQ         t113-1106(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t114-1110(SP)
        MOVL         R8, t112-1098(SP)
        MOVLQZX      t112-1098(SP), R9
        MOVLQZX      t114-1110(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t45-128(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t116-1122(SP)
        MOVQ         t116-1122(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t117-1126(SP)
        MOVL         R8, t115-1114(SP)
        MOVLQZX      t115-1114(SP), R9
        MOVLQZX      t117-1126(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t45-128(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t119-1138(SP)
        MOVQ         t119-1138(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t120-1142(SP)
        MOVL         R8, t118-1130(SP)
        MOVLQZX      t118-1130(SP), R9
        MOVLQZX      t120-1142(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t45-128(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t122-1154(SP)
        MOVQ         t122-1154(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t123-1158(SP)
        MOVL         R8, t121-1146(SP)
        MOVLQZX      t121-1146(SP), R9
        MOVLQZX      t123-1158(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVL         R8, t124-1162(SP)
        MOVLQZX      t89-974(SP), R9
        MOVLQZX      t124-1162(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X9, t25-48(SP)
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t25-48(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t126-1174(SP)
        MOVQ         t126-1174(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t127-1178(SP)
        MOVL         R8, t125-1166(SP)
        MOVLQZX      t125-1166(SP), R9
        MOVLQZX      t127-1178(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X6, t39-96(SP)
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t39-96(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t129-1190(SP)
        MOVQ         t129-1190(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t130-1194(SP)
        MOVL         R8, t128-1182(SP)
        MOVLQZX      t128-1182(SP), R9
        MOVLQZX      t130-1194(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X3, t53-144(SP)
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t53-144(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t132-1206(SP)
        MOVQ         t132-1206(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t133-1210(SP)
        MOVL         R8, t131-1198(SP)
        MOVLQZX      t131-1198(SP), R9
        MOVLQZX      t133-1210(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVL         R8, t134-1214(SP)
        MOVLQZX      t138-177(SP), R9
        MOVLQZX      t134-1214(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         t139-185(SP), SI
        MOVQ         $1, BX
        MOVQ         SI, DI
        ADDQ         BX, DI
        MOVL         R8, t138-177(SP)
        MOVQ         DI, t139-185(SP)
        MOVQ         DI, t136-1226(SP)
        MOVL         R8, t135-1218(SP)
        JMP block8
block7:
        MOVQ         t4-173(SP), R14
        MOVQ         $1, R13
        MOVQ         R14, R15
        ADDQ         R13, R15
        MOVLQZX      t138-177(SP), R12
        MOVL         R12, t3-165(SP)
        MOVQ         R15, t4-173(SP)
        MOVQ         R15, t137-1234(SP)
        JMP block5
block8:
        MOVQ         x+8(FP), R15
        MOVQ         R15, R14
        MOVQ         t139-185(SP), R12
        CMPQ         R12, R14
        SETLT        R13
        MOVB         R13, t141-1243(SP)
        CMPB         R13, $0
        JEQ          block7
        JMP          block6

