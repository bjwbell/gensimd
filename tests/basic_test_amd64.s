// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·t0simd(SB),$8-8
        MOVQ         $0, ret0+0(FP)
block0:
        MOVQ         $0, R15
        MOVQ         R15, ret0+0(FP)
        RET

TEXT ·t1simd(SB),$8-8
        MOVQ         $0, ret0+0(FP)
block0:
        MOVQ         $1, R15
        MOVQ         R15, ret0+0(FP)
        RET

TEXT ·t2simd(SB),$8-8
        MOVQ         $0, ret0+0(FP)
block0:
        MOVQ         $2, R15
        MOVQ         R15, ret0+0(FP)
        RET

TEXT ·t3simd(SB),$8-8
        MOVQ         $0, ret0+0(FP)
block0:
        MOVQ         $256, R15
        MOVQ         R15, ret0+0(FP)
        RET

TEXT ·t4simd(SB),$8-8
        MOVQ         $0, ret0+0(FP)
block0:
        MOVQ         $9223372036854775807, R15
        MOVQ         R15, ret0+0(FP)
        RET

