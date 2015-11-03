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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PADDB        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PSUBB        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PADDB        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PSUBB        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PADDW        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PSUBW        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PMULLW       X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
        RET

TEXT ·shli16x8s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t0-16(SP)
        MOVOU        t0-16(SP), X14
        MOVOU        X14, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X13
        MOVOU        t1-32(SP), X12
        PSLLW        X13, X12
        MOVOU        X12, t2-48(SP)
        MOVOU        t2-48(SP), X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shri16x8s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t0-16(SP)
        MOVOU        t0-16(SP), X14
        MOVOU        X14, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X13
        MOVOU        t1-32(SP), X12
        PSRAW        X13, X12
        MOVOU        X12, t2-48(SP)
        MOVOU        t2-48(SP), X13
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PADDW        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PSUBW        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PMULLW       X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
        RET

TEXT ·shlu16x8s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t0-16(SP)
        MOVOU        t0-16(SP), X14
        MOVOU        X14, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X13
        MOVOU        t1-32(SP), X12
        PSLLW        X13, X12
        MOVOU        X12, t2-48(SP)
        MOVOU        t2-48(SP), X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shru16x8s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t0-16(SP)
        MOVOU        t0-16(SP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t1-32(SP), X13
        MOVB         shift+16(FP), R15
        PEXTRW       $0, X13, R14
        MOVQ         R15, CX
        MOVL         $63, R13
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        MOVL         $1, R13
        XORQ         CX, CX
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        PINSRW       $0, R14, X13
        PEXTRW       $1, X13, R14
        MOVQ         R15, CX
        MOVL         $63, R13
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        MOVL         $1, R13
        XORQ         CX, CX
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        PINSRW       $1, R14, X13
        PEXTRW       $2, X13, R14
        MOVQ         R15, CX
        MOVL         $63, R13
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        MOVL         $1, R13
        XORQ         CX, CX
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        PINSRW       $2, R14, X13
        PEXTRW       $3, X13, R14
        MOVQ         R15, CX
        MOVL         $63, R13
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        MOVL         $1, R13
        XORQ         CX, CX
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        PINSRW       $3, R14, X13
        PEXTRW       $4, X13, R14
        MOVQ         R15, CX
        MOVL         $63, R13
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        MOVL         $1, R13
        XORQ         CX, CX
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        PINSRW       $4, R14, X13
        PEXTRW       $5, X13, R14
        MOVQ         R15, CX
        MOVL         $63, R13
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        MOVL         $1, R13
        XORQ         CX, CX
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        PINSRW       $5, R14, X13
        PEXTRW       $6, X13, R14
        MOVQ         R15, CX
        MOVL         $63, R13
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        MOVL         $1, R13
        XORQ         CX, CX
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        PINSRW       $6, R14, X13
        PEXTRW       $7, X13, R14
        MOVQ         R15, CX
        MOVL         $63, R13
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        MOVL         $1, R13
        XORQ         CX, CX
        CMPB         R15, $64
        CMOVQCC      R13, CX
        MOVBQZX      CL, CX
        SHRQ         CX, R14
        PINSRW       $7, R14, X13
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PADDL        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PSUBL        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t2-48(SP), X10
        MOVOU        t3-64(SP), X9
        MOVO         X9, X11
        PMULULQ      X10, X11
        PSRLO        $4, X10
        PSRLO        $4, X9
        MOVO         X9, X8
        PMULULQ      X10, X8
        PSHUFD       $8, X11, X7
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
        MOVOU        X15, t0-16(SP)
        MOVOU        t0-16(SP), X14
        MOVOU        X14, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X13
        MOVOU        t1-32(SP), X12
        PSLLL        X13, X12
        MOVOU        X12, t2-48(SP)
        MOVOU        t2-48(SP), X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shri32x4s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t0-16(SP)
        MOVOU        t0-16(SP), X14
        MOVOU        X14, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X13
        MOVOU        t1-32(SP), X12
        PSRAL        X13, X12
        MOVOU        X12, t2-48(SP)
        MOVOU        t2-48(SP), X13
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PADDL        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PSUBL        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t2-48(SP), X10
        MOVOU        t3-64(SP), X9
        MOVO         X9, X11
        PMULULQ      X10, X11
        PSRLO        $4, X10
        PSRLO        $4, X9
        MOVO         X9, X8
        PMULULQ      X10, X8
        PSHUFD       $8, X11, X7
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
        MOVOU        X15, t0-16(SP)
        MOVOU        t0-16(SP), X14
        MOVOU        X14, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X13
        MOVOU        t1-32(SP), X12
        PSLLL        X13, X12
        MOVOU        X12, t2-48(SP)
        MOVOU        t2-48(SP), X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shru32x4s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t0-16(SP)
        MOVOU        t0-16(SP), X14
        MOVOU        X14, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X13
        MOVOU        t1-32(SP), X12
        PSRLL        X13, X12
        MOVOU        X12, t2-48(SP)
        MOVOU        t2-48(SP), X13
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PADDQ        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PSUBQ        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PADDQ        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVOU        X15, t0-16(SP)
        MOVOU        y+16(FP), X14
        MOVOU        X14, t1-32(SP)
        MOVOU        t0-16(SP), X13
        MOVOU        X13, t2-48(SP)
        MOVOU        t1-32(SP), X12
        MOVOU        X12, t3-64(SP)
        MOVOU        t3-64(SP), X11
        MOVOU        t2-48(SP), X10
        PSUBQ        X11, X10
        MOVOU        X10, t4-80(SP)
        MOVOU        t4-80(SP), X9
        MOVOU        X9, ret0+32(FP)
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
        MOVUPS       X15, t0-16(SP)
        MOVUPS       y+16(FP), X14
        MOVUPS       X14, t1-32(SP)
        MOVUPS       t0-16(SP), X13
        MOVUPS       X13, t2-48(SP)
        MOVUPS       t1-32(SP), X12
        MOVUPS       X12, t3-64(SP)
        MOVUPS       t3-64(SP), X11
        MOVUPS       t2-48(SP), X10
        ADDPS        X11, X10
        MOVUPS       X10, t4-80(SP)
        MOVUPS       t4-80(SP), X9
        MOVUPS       X9, ret0+32(FP)
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
        MOVUPS       X15, t0-16(SP)
        MOVUPS       y+16(FP), X14
        MOVUPS       X14, t1-32(SP)
        MOVUPS       t0-16(SP), X13
        MOVUPS       X13, t2-48(SP)
        MOVUPS       t1-32(SP), X12
        MOVUPS       X12, t3-64(SP)
        MOVUPS       t3-64(SP), X11
        MOVUPS       t2-48(SP), X10
        SUBPS        X11, X10
        MOVUPS       X10, t4-80(SP)
        MOVUPS       t4-80(SP), X9
        MOVUPS       X9, ret0+32(FP)
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
        MOVUPS       X15, t0-16(SP)
        MOVUPS       y+16(FP), X14
        MOVUPS       X14, t1-32(SP)
        MOVUPS       t0-16(SP), X13
        MOVUPS       X13, t2-48(SP)
        MOVUPS       t1-32(SP), X12
        MOVUPS       X12, t3-64(SP)
        MOVUPS       t3-64(SP), X11
        MOVUPS       t2-48(SP), X10
        MULPS        X11, X10
        MOVUPS       X10, t4-80(SP)
        MOVUPS       t4-80(SP), X9
        MOVUPS       X9, ret0+32(FP)
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
        MOVUPS       X15, t0-16(SP)
        MOVUPS       y+16(FP), X14
        MOVUPS       X14, t1-32(SP)
        MOVUPS       t0-16(SP), X13
        MOVUPS       X13, t2-48(SP)
        MOVUPS       t1-32(SP), X12
        MOVUPS       X12, t3-64(SP)
        MOVUPS       t3-64(SP), X11
        MOVUPS       t2-48(SP), X10
        DIVPS        X11, X10
        MOVUPS       X10, t4-80(SP)
        MOVUPS       t4-80(SP), X9
        MOVUPS       X9, ret0+32(FP)
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
        MOVUPD       X15, t0-16(SP)
        MOVUPD       y+16(FP), X14
        MOVUPD       X14, t1-32(SP)
        MOVUPD       t0-16(SP), X13
        MOVUPD       X13, t2-48(SP)
        MOVUPD       t1-32(SP), X12
        MOVUPD       X12, t3-64(SP)
        MOVUPD       t3-64(SP), X11
        MOVUPD       t2-48(SP), X10
        ADDPD        X11, X10
        MOVUPD       X10, t4-80(SP)
        MOVUPD       t4-80(SP), X9
        MOVUPD       X9, ret0+32(FP)
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
        MOVUPD       X15, t0-16(SP)
        MOVUPD       y+16(FP), X14
        MOVUPD       X14, t1-32(SP)
        MOVUPD       t0-16(SP), X13
        MOVUPD       X13, t2-48(SP)
        MOVUPD       t1-32(SP), X12
        MOVUPD       X12, t3-64(SP)
        MOVUPD       t3-64(SP), X11
        MOVUPD       t2-48(SP), X10
        SUBPD        X11, X10
        MOVUPD       X10, t4-80(SP)
        MOVUPD       t4-80(SP), X9
        MOVUPD       X9, ret0+32(FP)
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
        MOVUPD       X15, t0-16(SP)
        MOVUPD       y+16(FP), X14
        MOVUPD       X14, t1-32(SP)
        MOVUPD       t0-16(SP), X13
        MOVUPD       X13, t2-48(SP)
        MOVUPD       t1-32(SP), X12
        MOVUPD       X12, t3-64(SP)
        MOVUPD       t3-64(SP), X11
        MOVUPD       t2-48(SP), X10
        MULPD        X11, X10
        MOVUPD       X10, t4-80(SP)
        MOVUPD       t4-80(SP), X9
        MOVUPD       X9, ret0+32(FP)
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
        MOVUPD       X15, t0-16(SP)
        MOVUPD       y+16(FP), X14
        MOVUPD       X14, t1-32(SP)
        MOVUPD       t0-16(SP), X13
        MOVUPD       X13, t2-48(SP)
        MOVUPD       t1-32(SP), X12
        MOVUPD       X12, t3-64(SP)
        MOVUPD       t3-64(SP), X11
        MOVUPD       t2-48(SP), X10
        DIVPD        X11, X10
        MOVUPD       X10, t4-80(SP)
        MOVUPD       t4-80(SP), X9
        MOVUPD       X9, ret0+32(FP)
        RET

