// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·uint8_t0_simd(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVBQZX      x+0(FP), R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·uint8_t1_simd(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVBQZX      x+0(FP), R14
        MOVB         $1, R13
        MOVB         R14, R15
        ADDB         R13, R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·uint8_t2_simd(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVBQZX      x+0(FP), R14
        MOVB         $2, R13
        MOVB         R14, R15
        MOVB         R15, AX
        MULB         R13
        MOVB         AX, R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·uint8_t3_simd(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVBQZX      x+0(FP), R14
        MOVB         $3, R13
        XORQ         AX, AX
        MOVB         R14, AX
        DIVB         R13
        MOVB         AX, R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·uint8_t4_simd(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVBQZX      x+0(FP), R14
        MOVB         R14, R15
        MOVB         R15, AX
        MULB         R14
        MOVB         AX, R15
        MOVB         R15, ret0+8(FP)
        RET

