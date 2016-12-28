// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·regspill2(SB),$1440-52
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
        MOVQ         $0, t31-112(SP)
        MOVQ         $0, t31-104(SP)
        MOVQ         $0, t35-128(SP)
        MOVQ         $0, t35-120(SP)
        MOVQ         $0, t39-144(SP)
        MOVQ         $0, t39-136(SP)
        MOVQ         $0, t43-160(SP)
        MOVQ         $0, t43-152(SP)
        MOVQ         $0, t47-176(SP)
        MOVQ         $0, t47-168(SP)
        MOVQ         $0, t51-192(SP)
        MOVQ         $0, t51-184(SP)
        MOVQ         $0, t55-208(SP)
        MOVQ         $0, t55-200(SP)
        MOVQ         $0, t59-224(SP)
        MOVQ         $0, t59-216(SP)
        MOVQ         $0, t63-240(SP)
        MOVQ         $0, t63-232(SP)
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
        MOVQ         $1, R13
        IMUL3Q       $16, R13, R13
        MOVQ         x+0(FP), R14
        ADDQ         R13, R14
        MOVQ         R14, R13
        MOVUPS       (R13), X15
        MOVUPS       X15, t5-281(SP)
        MOVQ         $0, R12
        IMUL3Q       $16, R12, R12
        MOVQ         x+0(FP), R13
        ADDQ         R12, R13
        MOVQ         R13, R12
        MOVUPS       (R12), X15
        MOVUPS       X15, t7-305(SP)
        MOVOU        t7-305(SP), X15
        MOVOU        t5-281(SP), X14
        PSUBL        X15, X14
        MOVO         X14, X13
        MOVQ         $1, R11
        IMUL3Q       $16, R11, R11
        MOVQ         y+24(FP), R12
        ADDQ         R11, R12
        MOVQ         R12, R11
        MOVUPS       (R11), X12
        MOVUPS       X12, t11-345(SP)
        MOVQ         $0, R10
        IMUL3Q       $16, R10, R10
        MOVQ         y+24(FP), R11
        ADDQ         R10, R11
        MOVQ         R11, R10
        MOVUPS       (R10), X12
        MOVUPS       X12, t13-369(SP)
        MOVOU        t13-369(SP), X12
        MOVOU        t11-345(SP), X11
        PSUBL        X12, X11
        MOVO         X11, X10
        MOVO         X13, X9
        MOVO         X13, X8
        MOVO         X8, X7
        PMULULQ      X9, X7
        MOVOU        X9, t16-401(SP)
        PSRLO        $4, X9
        MOVOU        X8, t17-417(SP)
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
        MOVOU        X8, t20-449(SP)
        PSRLO        $4, X8
        MOVOU        X7, t21-465(SP)
        PSRLO        $4, X7
        MOVO         X7, X4
        PMULULQ      X8, X4
        PSHUFD       $8, X6, X3
        PSHUFD       $8, X4, X2
        PUNPCKLLQ    X2, X3
        MOVO         X3, X8
        MOVO         X9, X7
        MOVO         X8, X6
        MOVOU        X7, t24-497(SP)
        PADDL        X6, X7
        MOVO         X7, X4
        MOVO         X13, X2
        MOVO         X10, X1
        MOVOU        X2, t28-545(SP)
        PSUBL        X1, X2
        MOVO         X2, X0
        MOVOU        X9, t32-593(SP)
        MOVOU        X8, t33-609(SP)
        MOVOU        t33-609(SP), X1
        MOVOU        t32-593(SP), X2
        PSUBL        X1, X2
        MOVOU        X2, t31-112(SP)
        MOVOU        X0, t36-641(SP)
        MOVOU        X0, t37-657(SP)
        MOVOU        X0, t27-96(SP)
        MOVOU        t36-641(SP), X1
        MOVOU        t37-657(SP), X2
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
        MOVOU        t31-112(SP), X3
        MOVO         X3, X2
        MOVO         X3, X1
        MOVO         X1, X0
        PMULULQ      X2, X0
        MOVOU        X2, t40-689(SP)
        PSRLO        $4, X2
        MOVOU        X1, t41-705(SP)
        PSRLO        $4, X1
        MOVO         X1, X3
        PMULULQ      X2, X3
        PSHUFD       $8, X0, X5
        PSHUFD       $8, X3, X7
        PUNPCKLLQ    X7, X5
        MOVO         X5, X7
        MOVO         X6, X3
        MOVO         X7, X2
        MOVOU        X3, t44-737(SP)
        PADDL        X2, X3
        MOVO         X3, X1
        MOVOU        t27-96(SP), X0
        MOVOU        X0, t48-785(SP)
        MOVOU        t31-112(SP), X0
        MOVOU        X0, t49-801(SP)
        MOVOU        t49-801(SP), X0
        MOVOU        t48-785(SP), X2
        PSUBL        X0, X2
        MOVOU        X2, t47-176(SP)
        MOVOU        X6, t52-833(SP)
        MOVOU        X7, t53-849(SP)
        MOVOU        t53-849(SP), X0
        MOVOU        t52-833(SP), X2
        PSUBL        X0, X2
        MOVOU        X2, t51-192(SP)
        MOVOU        t47-176(SP), X0
        MOVOU        X0, t56-881(SP)
        MOVOU        X0, t57-897(SP)
        MOVOU        t56-881(SP), X2
        MOVOU        t57-897(SP), X3
        MOVO         X3, X0
        PMULULQ      X2, X0
        PSRLO        $4, X2
        PSRLO        $4, X3
        MOVO         X3, X5
        PMULULQ      X2, X5
        MOVOU        X6, t35-128(SP)
        PSHUFD       $8, X0, X6
        MOVOU        X7, t39-144(SP)
        PSHUFD       $8, X5, X7
        PUNPCKLLQ    X7, X6
        MOVO         X6, X7
        MOVOU        t51-192(SP), X5
        MOVO         X5, X3
        MOVO         X5, X2
        MOVO         X2, X0
        PMULULQ      X3, X0
        MOVOU        X3, t60-929(SP)
        PSRLO        $4, X3
        MOVOU        X2, t61-945(SP)
        PSRLO        $4, X2
        MOVO         X2, X5
        PMULULQ      X3, X5
        PSHUFD       $8, X0, X6
        MOVOU        X8, t19-64(SP)
        PSHUFD       $8, X5, X8
        PUNPCKLLQ    X8, X6
        MOVO         X6, X8
        MOVO         X7, X5
        MOVO         X8, X3
        MOVOU        X5, t64-977(SP)
        PADDL        X3, X5
        MOVO         X5, X2
        MOVOU        X13, t3-16(SP)
        MOVQ         $0, R9
        IMUL3Q       $4, R9, R9
        LEAQ         t3-16(SP), R10
        ADDQ         R9, R10
        MOVL         (R10), R9
        MOVL         R9, t68-1021(SP)
        MOVQ         $1, R8
        IMUL3Q       $4, R8, R8
        LEAQ         t3-16(SP), R9
        ADDQ         R8, R9
        MOVL         (R9), R8
        MOVL         R8, t70-1033(SP)
        MOVLQZX      t68-1021(SP), R9
        MOVLQZX      t70-1033(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, BX
        IMUL3Q       $4, BX, BX
        LEAQ         t3-16(SP), BP
        ADDQ         BX, BP
        MOVL         (BP), BX
        MOVL         BX, t73-1049(SP)
        MOVL         R8, t71-1037(SP)
        MOVLQZX      t71-1037(SP), R9
        MOVLQZX      t73-1049(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, DI
        IMUL3Q       $4, DI, DI
        LEAQ         t3-16(SP), BX
        ADDQ         DI, BX
        MOVL         (BX), DI
        MOVL         DI, t76-1065(SP)
        MOVL         R8, t74-1053(SP)
        MOVLQZX      t74-1053(SP), R9
        MOVLQZX      t76-1065(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t27-96(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t78-1077(SP)
        MOVQ         t78-1077(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t79-1081(SP)
        MOVL         R8, t77-1069(SP)
        MOVLQZX      t77-1069(SP), R9
        MOVLQZX      t79-1081(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t27-96(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t81-1093(SP)
        MOVQ         t81-1093(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t82-1097(SP)
        MOVL         R8, t80-1085(SP)
        MOVLQZX      t80-1085(SP), R9
        MOVLQZX      t82-1097(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t27-96(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t84-1109(SP)
        MOVQ         t84-1109(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t85-1113(SP)
        MOVL         R8, t83-1101(SP)
        MOVLQZX      t83-1101(SP), R9
        MOVLQZX      t85-1113(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t27-96(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t87-1125(SP)
        MOVQ         t87-1125(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t88-1129(SP)
        MOVL         R8, t86-1117(SP)
        MOVLQZX      t86-1117(SP), R9
        MOVLQZX      t88-1129(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t47-176(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t90-1141(SP)
        MOVQ         t90-1141(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t91-1145(SP)
        MOVL         R8, t89-1133(SP)
        MOVLQZX      t89-1133(SP), R9
        MOVLQZX      t91-1145(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t47-176(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t93-1157(SP)
        MOVQ         t93-1157(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t94-1161(SP)
        MOVL         R8, t92-1149(SP)
        MOVLQZX      t92-1149(SP), R9
        MOVLQZX      t94-1161(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t47-176(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t96-1173(SP)
        MOVQ         t96-1173(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t97-1177(SP)
        MOVL         R8, t95-1165(SP)
        MOVLQZX      t95-1165(SP), R9
        MOVLQZX      t97-1177(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t47-176(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t99-1189(SP)
        MOVQ         t99-1189(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t100-1193(SP)
        MOVL         R8, t98-1181(SP)
        MOVLQZX      t98-1181(SP), R9
        MOVLQZX      t100-1193(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X10, t9-32(SP)
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t9-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t102-1205(SP)
        MOVQ         t102-1205(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t103-1209(SP)
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t9-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t104-1217(SP)
        MOVQ         t104-1217(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t105-1221(SP)
        MOVL         R8, t101-1197(SP)
        MOVLQZX      t103-1209(SP), R9
        MOVLQZX      t105-1221(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t9-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t107-1233(SP)
        MOVQ         t107-1233(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t108-1237(SP)
        MOVL         R8, t106-1225(SP)
        MOVLQZX      t106-1225(SP), R9
        MOVLQZX      t108-1237(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t9-32(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t110-1249(SP)
        MOVQ         t110-1249(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t111-1253(SP)
        MOVL         R8, t109-1241(SP)
        MOVLQZX      t109-1241(SP), R9
        MOVLQZX      t111-1253(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t31-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t113-1265(SP)
        MOVQ         t113-1265(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t114-1269(SP)
        MOVL         R8, t112-1257(SP)
        MOVLQZX      t112-1257(SP), R9
        MOVLQZX      t114-1269(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t31-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t116-1281(SP)
        MOVQ         t116-1281(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t117-1285(SP)
        MOVL         R8, t115-1273(SP)
        MOVLQZX      t115-1273(SP), R9
        MOVLQZX      t117-1285(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t31-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t119-1297(SP)
        MOVQ         t119-1297(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t120-1301(SP)
        MOVL         R8, t118-1289(SP)
        MOVLQZX      t118-1289(SP), R9
        MOVLQZX      t120-1301(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t31-112(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t122-1313(SP)
        MOVQ         t122-1313(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t123-1317(SP)
        MOVL         R8, t121-1305(SP)
        MOVLQZX      t121-1305(SP), R9
        MOVLQZX      t123-1317(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t51-192(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t125-1329(SP)
        MOVQ         t125-1329(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t126-1333(SP)
        MOVL         R8, t124-1321(SP)
        MOVLQZX      t124-1321(SP), R9
        MOVLQZX      t126-1333(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t51-192(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t128-1345(SP)
        MOVQ         t128-1345(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t129-1349(SP)
        MOVL         R8, t127-1337(SP)
        MOVLQZX      t127-1337(SP), R9
        MOVLQZX      t129-1349(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t51-192(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t131-1361(SP)
        MOVQ         t131-1361(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t132-1365(SP)
        MOVL         R8, t130-1353(SP)
        MOVLQZX      t130-1353(SP), R9
        MOVLQZX      t132-1365(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVQ         $3, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t51-192(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t134-1377(SP)
        MOVQ         t134-1377(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t135-1381(SP)
        MOVL         R8, t133-1369(SP)
        MOVLQZX      t133-1369(SP), R9
        MOVLQZX      t135-1381(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVL         R8, t136-1385(SP)
        MOVLQZX      t101-1197(SP), R9
        MOVLQZX      t136-1385(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X4, t23-80(SP)
        MOVQ         $0, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t23-80(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t138-1397(SP)
        MOVQ         t138-1397(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t139-1401(SP)
        MOVL         R8, t137-1389(SP)
        MOVLQZX      t137-1389(SP), R9
        MOVLQZX      t139-1401(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X1, t43-160(SP)
        MOVQ         $1, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t43-160(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t141-1413(SP)
        MOVQ         t141-1413(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t142-1417(SP)
        MOVL         R8, t140-1405(SP)
        MOVLQZX      t140-1405(SP), R9
        MOVLQZX      t142-1417(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVOU        X2, t63-240(SP)
        MOVQ         $2, SI
        IMUL3Q       $4, SI, SI
        LEAQ         t63-240(SP), DI
        ADDQ         SI, DI
        MOVQ         DI, t144-1429(SP)
        MOVQ         t144-1429(SP), BX
        MOVL         (BX), SI
        MOVL         SI, t145-1433(SP)
        MOVL         R8, t143-1421(SP)
        MOVLQZX      t143-1421(SP), R9
        MOVLQZX      t145-1433(SP), R10
        MOVL         R9, R8
        ADDL         R10, R8
        MOVL         R8, ret0+48(FP)
        RET

