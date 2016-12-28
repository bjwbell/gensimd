// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·addi8x16s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PADDB        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·subi8x16s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PSUBB        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·addu8x16s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PADDB        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·subu8x16s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PSUBB        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·addi16x8s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PADDW        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·subi16x8s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PSUBW        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·muli16x8s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PMULLW       X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·shli16x8s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVO         X14, X13
        MOVBQZX      shift+16(FP), R15
        MOVQ         R15, X12
        MOVOU        X13, t1-32(SP)
        PSLLW        X12, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shri16x8s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVO         X14, X13
        MOVBQZX      shift+16(FP), R15
        MOVQ         R15, X12
        MOVOU        X13, t1-32(SP)
        PSRAW        X12, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·addu16x8s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PADDW        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·subu16x8s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PSUBW        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·mulu16x8s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PMULLW       X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·shlu16x8s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVO         X14, X13
        MOVBQZX      shift+16(FP), R15
        MOVQ         R15, X12
        MOVOU        X13, t1-32(SP)
        PSLLW        X12, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shru16x8s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVO         X14, X13
        MOVBQZX      shift+16(FP), R15
        MOVQ         R15, X12
        MOVOU        X13, t1-32(SP)
        PSRLW        X12, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·addi32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PADDL        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·subi32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PSUBL        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·muli32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVO         X10, X9
        PMULULQ      X11, X9
        MOVOU        X11, t2-48(SP)
        PSRLO        $4, X11
        MOVOU        X10, t3-64(SP)
        PSRLO        $4, X10
        MOVO         X10, X8
        PMULULQ      X11, X8
        PSHUFD       $8, X9, X7
        PSHUFD       $8, X8, X6
        PUNPCKLLQ    X6, X7
        MOVOU        X7, ret0+32(FP)
        RET

TEXT ·shli32x4s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVO         X14, X13
        MOVBQZX      shift+16(FP), R15
        MOVQ         R15, X12
        MOVOU        X13, t1-32(SP)
        PSLLL        X12, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shri32x4s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVO         X14, X13
        MOVBQZX      shift+16(FP), R15
        MOVQ         R15, X12
        MOVOU        X13, t1-32(SP)
        PSRAL        X12, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·addu32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PADDL        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·subu32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PSUBL        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·mulu32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVO         X10, X9
        PMULULQ      X11, X9
        MOVOU        X11, t2-48(SP)
        PSRLO        $4, X11
        MOVOU        X10, t3-64(SP)
        PSRLO        $4, X10
        MOVO         X10, X8
        PMULULQ      X11, X8
        PSHUFD       $8, X9, X7
        PSHUFD       $8, X8, X6
        PUNPCKLLQ    X6, X7
        MOVOU        X7, ret0+32(FP)
        RET

TEXT ·shlu32x4s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVO         X14, X13
        MOVBQZX      shift+16(FP), R15
        MOVQ         R15, X12
        MOVOU        X13, t1-32(SP)
        PSLLL        X12, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shru32x4s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVO         X14, X13
        MOVBQZX      shift+16(FP), R15
        MOVQ         R15, X12
        MOVOU        X13, t1-32(SP)
        PSRLL        X12, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·addi64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PADDQ        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·subi64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PSUBQ        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·addu64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PADDQ        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·subu64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PSUBQ        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·addf32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVO         X15, X14
        MOVUPS       y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVUPS       X11, t2-48(SP)
        ADDPS        X10, X11
        MOVUPS       X11, ret0+32(FP)
        RET

TEXT ·subf32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVO         X15, X14
        MOVUPS       y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVUPS       X11, t2-48(SP)
        SUBPS        X10, X11
        MOVUPS       X11, ret0+32(FP)
        RET

TEXT ·mulf32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVO         X15, X14
        MOVUPS       y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVUPS       X11, t2-48(SP)
        MULPS        X10, X11
        MOVUPS       X11, ret0+32(FP)
        RET

TEXT ·divf32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVO         X15, X14
        MOVUPS       y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVUPS       X11, t2-48(SP)
        DIVPS        X10, X11
        MOVUPS       X11, ret0+32(FP)
        RET

TEXT ·addf64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVUPD       x+0(FP), X15
        MOVO         X15, X14
        MOVUPD       y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVUPD       X11, t2-48(SP)
        ADDPD        X10, X11
        MOVUPD       X11, ret0+32(FP)
        RET

TEXT ·subf64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVUPD       x+0(FP), X15
        MOVO         X15, X14
        MOVUPD       y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVUPD       X11, t2-48(SP)
        SUBPD        X10, X11
        MOVUPD       X11, ret0+32(FP)
        RET

TEXT ·mulf64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVUPD       x+0(FP), X15
        MOVO         X15, X14
        MOVUPD       y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVUPD       X11, t2-48(SP)
        MULPD        X10, X11
        MOVUPD       X11, ret0+32(FP)
        RET

TEXT ·divf64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVUPD       x+0(FP), X15
        MOVO         X15, X14
        MOVUPD       y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVUPD       X11, t2-48(SP)
        DIVPD        X10, X11
        MOVUPD       X11, ret0+32(FP)
        RET

