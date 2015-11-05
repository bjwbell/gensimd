// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·regspill3(SB),$1520-52
        MOVL         $0, ret0+48(FP)
        MOVQ         $0, t9-16(SP)
        MOVQ         $0, t9-8(SP)
        MOVQ         $0, t15-32(SP)
        MOVQ         $0, t15-24(SP)
        MOVQ         $0, t21-48(SP)
        MOVQ         $0, t21-40(SP)
        MOVQ         $0, t25-64(SP)
        MOVQ         $0, t25-56(SP)
        MOVQ         $0, t29-80(SP)
        MOVQ         $0, t29-72(SP)
        MOVQ         $0, t33-96(SP)
        MOVQ         $0, t33-88(SP)
        MOVQ         $0, t37-112(SP)
        MOVQ         $0, t37-104(SP)
        MOVQ         $0, t41-128(SP)
        MOVQ         $0, t41-120(SP)
        MOVQ         $0, t45-144(SP)
        MOVQ         $0, t45-136(SP)
        MOVQ         $0, t49-160(SP)
        MOVQ         $0, t49-152(SP)
        MOVQ         $0, t53-176(SP)
        MOVQ         $0, t53-168(SP)
        MOVQ         $0, t57-192(SP)
        MOVQ         $0, t57-184(SP)
        MOVQ         $0, t61-208(SP)
        MOVQ         $0, t61-200(SP)
        MOVQ         $0, t65-224(SP)
        MOVQ         $0, t65-216(SP)
        MOVQ         $0, t69-240(SP)
        MOVQ         $0, t69-232(SP)
block0:
        MOVQ         x+8(FP), R15
        MOVQ         R15, R14
        MOVQ         y+32(FP), R13
        MOVQ         R13, R12
        CMPQ         R14, R12
        SETNE        R11
        MOVB         R11, t2-257(SP)
        CMPB         R11, $0
        JEQ          block2
        JMP          block1
block1:
        MOVL         $-1, R15
        MOVL         R15, ret0+48(FP)
        RET
block2:
        MOVL         $2147483647, R14
        MOVL         R14, t3-261(SP)
        MOVQ         $0, R13
        MOVQ         R13, t4-269(SP)
        MOVL         $0, R12
        MOVL         R12, t5-273(SP)
        MOVL         R12, t6-277(SP)
        JMP block5
block3:
        MOVL         t3-261(SP), R15
        MOVL         R15, t156-281(SP)
        MOVQ         $0, R14
        MOVQ         R14, t157-289(SP)
        MOVL         t5-273(SP), R13
        MOVL         R13, t158-293(SP)
        MOVL         t6-277(SP), R12
        MOVL         R12, t159-297(SP)
        JMP block8
block4:
        MOVL         t3-261(SP), R15
        MOVL         R15, ret0+48(FP)
        RET
block5:
        MOVQ         x+8(FP), R15
        MOVQ         R15, R14
        MOVQ         t4-269(SP), R12
        CMPQ         R12, R14
        SETLT        R13
        MOVB         R13, t8-306(SP)
        CMPB         R13, $0
        JEQ          block4
        JMP          block3
block6:
        MOVQ         t157-289(SP), R14
        IMUL3Q       $16, R14, R14
        MOVQ         x+0(FP), R15
        ADDQ         R14, R15
        MOVQ         R15, R14
        MOVUPS       (R14), X15
        MOVUPS       X15, t11-330(SP)
        MOVQ         t4-269(SP), R13
        IMUL3Q       $16, R13, R13
        MOVQ         x+0(FP), R14
        ADDQ         R13, R14
        MOVQ         R14, R13
        MOVUPS       (R13), X15
        MOVUPS       X15, t13-354(SP)
        MOVOU        t13-354(SP), X15
        MOVOU        t11-330(SP), X14
        PSUBL        X15, X14
        MOVO         X14, X13
        MOVQ         t157-289(SP), R12
        IMUL3Q       $16, R12, R12
        MOVQ         y+24(FP), R13
        ADDQ         R12, R13
        MOVQ         R13, R12
        MOVUPS       (R12), X12
        MOVUPS       X12, t17-394(SP)
        MOVQ         t4-269(SP), R11
        IMUL3Q       $16, R11, R11
        MOVQ         y+24(FP), R12
        ADDQ         R11, R12
        MOVQ         R12, R11
        MOVUPS       (R11), X12
        MOVUPS       X12, t19-418(SP)
        MOVOU        t19-418(SP), X12
        MOVOU        t17-394(SP), X11
        PSUBL        X12, X11
        MOVO         X11, X10
        MOVO         X13, X9
        MOVO         X13, X8
        MOVO         X8, X7
        PMULULQ      X9, X7
        MOVOU        X9, t22-450(SP)
        PSRLO        $4, X9
        MOVOU        X8, t23-466(SP)
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
        MOVOU        X8, t26-498(SP)
        PSRLO        $4, X8
        MOVOU        X7, t27-514(SP)
        PSRLO        $4, X7
        MOVO         X7, X4
        PMULULQ      X8, X4
        PSHUFD       $8, X6, X3
        PSHUFD       $8, X4, X2
        PUNPCKLLQ    X2, X3
        MOVO         X3, X8
        MOVO         X9, X7
        MOVO         X8, X6
        PADDL        X6, X7
        MOVO         X7, X4
        MOVO         X13, X2
        MOVO         X10, X1
        PSUBL        X1, X2
        MOVO         X2, X0
        MOVOU        X9, t38-642(SP)
        MOVOU        X8, t39-658(SP)
        MOVOU        t39-658(SP), X1
        MOVOU        t38-642(SP), X2
        PSUBL        X1, X2
        MOVOU        X2, t37-112(SP)
        MOVOU        X0, t42-690(SP)
        MOVOU        X0, t43-706(SP)
        MOVOU        X0, t33-96(SP)
        MOVOU        t42-690(SP), X1
        MOVOU        t43-706(SP), X2
        MOVO         X2, X0
        PMULULQ      X1, X0
        PSRLO        $4, X1
        PSRLO        $4, X2
        MOVO         X2, X3
        PMULULQ      X1, X3
        PSHUFD       $8, X0, X5
        PSHUFD       $8, X3, X6
        PUNPCKLLQ    X6, X5
        MOVO         X5, X6
        MOVOU        t37-112(SP), X3
        MOVO         X3, X2
        MOVO         X3, X1
        MOVO         X1, X0
        PMULULQ      X2, X0
        MOVOU        X2, t46-738(SP)
        PSRLO        $4, X2
        MOVOU        X1, t47-754(SP)
        PSRLO        $4, X1
        MOVO         X1, X3
        PMULULQ      X2, X3
        PSHUFD       $8, X0, X5
        PSHUFD       $8, X3, X7
        PUNPCKLLQ    X7, X5
        MOVO         X5, X7
        MOVO         X6, X3
        MOVO         X7, X2
        PADDL        X2, X3
        MOVO         X3, X1
        MOVOU        t33-96(SP), X0
        MOVOU        X0, t54-834(SP)
        MOVOU        t37-112(SP), X0
        MOVOU        X0, t55-850(SP)
        MOVOU        t55-850(SP), X0
        MOVOU        t54-834(SP), X2
        PSUBL        X0, X2
        MOVOU        X2, t53-176(SP)
        MOVOU        X6, t58-882(SP)
        MOVOU        X7, t59-898(SP)
        MOVOU        t59-898(SP), X0
        MOVOU        t58-882(SP), X2
        PSUBL        X0, X2
        MOVOU        X2, t57-192(SP)
        MOVOU        t53-176(SP), X0
        MOVOU        X0, t62-930(SP)
        MOVOU        X0, t63-946(SP)
        MOVOU        t62-930(SP), X2
        MOVOU        t63-946(SP), X3
        MOVO         X3, X0
        PMULULQ      X2, X0
        PSRLO        $4, X2
        PSRLO        $4, X3
        MOVO         X3, X5
        PMULULQ      X2, X5
        MOVOU        X6, t41-128(SP)
        PSHUFD       $8, X0, X6
        MOVOU        X7, t45-144(SP)
        PSHUFD       $8, X5, X7
        PUNPCKLLQ    X7, X6
        MOVO         X6, X7
        MOVOU        t57-192(SP), X5
        MOVO         X5, X3
        MOVO         X5, X2
        MOVO         X2, X0
        PMULULQ      X3, X0
        MOVOU        X3, t66-978(SP)
        PSRLO        $4, X3
        MOVOU        X2, t67-994(SP)
        PSRLO        $4, X2
        MOVO         X2, X5
        PMULULQ      X3, X5
        PSHUFD       $8, X0, X6
        MOVOU        X8, t25-64(SP)
        PSHUFD       $8, X5, X8
        PUNPCKLLQ    X8, X6
        MOVO         X6, X8
        MOVO         X7, X5
        MOVO         X8, X3
        PADDL        X3, X5
        MOVO         X5, X2
        MOVOU        X13, t9-16(SP)
        MOVQ         $0, R10
        IMUL3Q       $4, R10, R10
        LEAQ         t9-16(SP), R11
        ADDQ         R10, R11
        MOVL         (R11), R10
        MOVL         R10, t74-1070(SP)
        MOVQ         $1, R9
        IMUL3Q       $4, R9, R9
        LEAQ         t9-16(SP), R10
        ADDQ         R9, R10
        MOVL         (R10), R9
        MOVL         R9, t76-1082(SP)
        MOVL         t74-1070(SP), R8
        MOVL         t76-1082(SP), R10
        MOVL         R8, R9
        ADDL         R10, R9
        MOVQ         $2, BX
        IMUL3Q       $4, BX, BX
        LEAQ         t9-16(SP), BP
        ADDQ         BX, BP
        MOVL         (BP), BX
        MOVL         BX, t79-1098(SP)
        MOVL         t79-1098(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, DI
        IMUL3Q       $4, DI, DI
        LEAQ         t9-16(SP), BX
        ADDQ         DI, BX
        MOVL         (BX), DI
        MOVL         DI, t82-1114(SP)
        MOVL         R8, t80-1102(SP)
        MOVL         t80-1102(SP), R9
        MOVL         t82-1114(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t33-96(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t84-1126(SP)
        MOVQ         t84-1126(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t85-1130(SP)
        MOVL         R8, t83-1118(SP)
        MOVL         t83-1118(SP), R9
        MOVL         t85-1130(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t33-96(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t87-1142(SP)
        MOVQ         t87-1142(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t88-1146(SP)
        MOVL         R8, t86-1134(SP)
        MOVL         t86-1134(SP), R9
        MOVL         t88-1146(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t33-96(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t90-1158(SP)
        MOVQ         t90-1158(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t91-1162(SP)
        MOVL         R8, t89-1150(SP)
        MOVL         t89-1150(SP), R9
        MOVL         t91-1162(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t33-96(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t93-1174(SP)
        MOVQ         t93-1174(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t94-1178(SP)
        MOVL         R8, t92-1166(SP)
        MOVL         t92-1166(SP), R9
        MOVL         t94-1178(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t53-176(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t96-1190(SP)
        MOVQ         t96-1190(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t97-1194(SP)
        MOVL         R8, t95-1182(SP)
        MOVL         t95-1182(SP), R9
        MOVL         t97-1194(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t53-176(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t99-1206(SP)
        MOVQ         t99-1206(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t100-1210(SP)
        MOVL         R8, t98-1198(SP)
        MOVL         t98-1198(SP), R9
        MOVL         t100-1210(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t53-176(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t102-1222(SP)
        MOVQ         t102-1222(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t103-1226(SP)
        MOVL         R8, t101-1214(SP)
        MOVL         t101-1214(SP), R9
        MOVL         t103-1226(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t53-176(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t105-1238(SP)
        MOVQ         t105-1238(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t106-1242(SP)
        MOVL         R8, t104-1230(SP)
        MOVL         t104-1230(SP), R9
        MOVL         t106-1242(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X10, t15-32(SP)
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t15-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t108-1254(SP)
        MOVQ         t108-1254(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t109-1258(SP)
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t15-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t110-1266(SP)
        MOVQ         t110-1266(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t111-1270(SP)
        MOVL         R8, t107-1246(SP)
        MOVL         t109-1258(SP), R9
        MOVL         t111-1270(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t15-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t113-1282(SP)
        MOVQ         t113-1282(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t114-1286(SP)
        MOVL         R8, t112-1274(SP)
        MOVL         t112-1274(SP), R9
        MOVL         t114-1286(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t15-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t116-1298(SP)
        MOVQ         t116-1298(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t117-1302(SP)
        MOVL         R8, t115-1290(SP)
        MOVL         t115-1290(SP), R9
        MOVL         t117-1302(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t37-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t119-1314(SP)
        MOVQ         t119-1314(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t120-1318(SP)
        MOVL         R8, t118-1306(SP)
        MOVL         t118-1306(SP), R9
        MOVL         t120-1318(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t37-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t122-1330(SP)
        MOVQ         t122-1330(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t123-1334(SP)
        MOVL         R8, t121-1322(SP)
        MOVL         t121-1322(SP), R9
        MOVL         t123-1334(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t37-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t125-1346(SP)
        MOVQ         t125-1346(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t126-1350(SP)
        MOVL         R8, t124-1338(SP)
        MOVL         t124-1338(SP), R9
        MOVL         t126-1350(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t37-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t128-1362(SP)
        MOVQ         t128-1362(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t129-1366(SP)
        MOVL         R8, t127-1354(SP)
        MOVL         t127-1354(SP), R9
        MOVL         t129-1366(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t57-192(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t131-1378(SP)
        MOVQ         t131-1378(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t132-1382(SP)
        MOVL         R8, t130-1370(SP)
        MOVL         t130-1370(SP), R9
        MOVL         t132-1382(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t57-192(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t134-1394(SP)
        MOVQ         t134-1394(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t135-1398(SP)
        MOVL         R8, t133-1386(SP)
        MOVL         t133-1386(SP), R9
        MOVL         t135-1398(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t57-192(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t137-1410(SP)
        MOVQ         t137-1410(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t138-1414(SP)
        MOVL         R8, t136-1402(SP)
        MOVL         t136-1402(SP), R9
        MOVL         t138-1414(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t57-192(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t140-1426(SP)
        MOVQ         t140-1426(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t141-1430(SP)
        MOVL         R8, t139-1418(SP)
        MOVL         t139-1418(SP), R9
        MOVL         t141-1430(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVL         R8, t142-1434(SP)
        MOVL         t107-1246(SP), R9
        MOVL         t142-1434(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X4, t29-80(SP)
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t29-80(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t144-1446(SP)
        MOVQ         t144-1446(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t145-1450(SP)
        MOVL         R8, t143-1438(SP)
        MOVL         t143-1438(SP), R9
        MOVL         t145-1450(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X1, t49-160(SP)
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t49-160(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t147-1462(SP)
        MOVQ         t147-1462(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t148-1466(SP)
        MOVL         R8, t146-1454(SP)
        MOVL         t146-1454(SP), R9
        MOVL         t148-1466(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X2, t69-240(SP)
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t69-240(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t150-1478(SP)
        MOVQ         t150-1478(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t151-1482(SP)
        MOVL         R8, t149-1470(SP)
        MOVL         t149-1470(SP), R9
        MOVL         t151-1482(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVL         R8, t152-1486(SP)
        MOVL         t156-281(SP), R9
        MOVL         t152-1486(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         t157-289(SP), SI
        MOVQ         $1, BX
        MOVQ         SI, DI
        ADDQ         BX, DI
        MOVL         R8, t156-281(SP)
        MOVQ         DI, t157-289(SP)
        MOVL         t107-1246(SP), R9
        MOVL         R9, t158-293(SP)
        MOVL         R8, t153-1490(SP)
        MOVL         t142-1434(SP), R8
        MOVL         R8, t159-297(SP)
        MOVQ         DI, t154-1498(SP)
        MOVOU        X7, t61-208(SP)
        MOVOU        X8, t65-224(SP)
        MOVOU        X9, t21-48(SP)
        JMP block8
block7:
        MOVQ         t4-269(SP), R14
        MOVQ         $1, R13
        MOVQ         R14, R15
        ADDQ         R13, R15
        MOVL         t156-281(SP), R12
        MOVL         R12, t3-261(SP)
        MOVQ         R15, t4-269(SP)
        MOVL         t158-293(SP), R14
        MOVL         R14, t5-273(SP)
        MOVL         t159-297(SP), R11
        MOVL         R11, t6-277(SP)
        MOVQ         R15, t155-1506(SP)
        JMP block5
block8:
        MOVQ         x+8(FP), R15
        MOVQ         R15, R14
        MOVQ         t157-289(SP), R12
        CMPQ         R12, R14
        SETLT        R13
        MOVB         R13, t161-1515(SP)
        CMPB         R13, $0
        JEQ          block7
        JMP          block6

