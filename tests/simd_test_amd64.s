// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·addi8x16s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PADDB        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·subi8x16s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PSUBB        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·addu8x16s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PADDB        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·subu8x16s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PSUBB        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·addi16x8s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PADDW        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·subi16x8s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PSUBW        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·muli16x8s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PMULLW       X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·shli16x8s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X14
        MOVOU        t1-32(SP), X13
        PSLLW        X14, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shri16x8s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X14
        MOVOU        t1-32(SP), X13
        PSRAW        X14, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·addu16x8s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PADDW        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·subu16x8s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PSUBW        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·mulu16x8s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PMULLW       X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·shlu16x8s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X14
        MOVOU        t1-32(SP), X13
        PSLLW        X14, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shru16x8s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t1-32(SP)
        MOVOU        t1-32(SP), X14
        MOVB         shift+16(FP), R15
        PEXTRW       $0, X14, R14
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
        PINSRW       $0, R14, X14
        PEXTRW       $1, X14, R14
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
        PINSRW       $1, R14, X14
        PEXTRW       $2, X14, R14
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
        PINSRW       $2, R14, X14
        PEXTRW       $3, X14, R14
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
        PINSRW       $3, R14, X14
        PEXTRW       $4, X14, R14
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
        PINSRW       $4, R14, X14
        PEXTRW       $5, X14, R14
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
        PINSRW       $5, R14, X14
        PEXTRW       $6, X14, R14
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
        PINSRW       $6, R14, X14
        PEXTRW       $7, X14, R14
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
        PINSRW       $7, R14, X14
        MOVOU        X14, ret0+24(FP)
        RET

TEXT ·addi32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PADDL        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·subi32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PSUBL        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·muli32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t2-48(SP), X12
        MOVOU        t3-64(SP), X11
        MOVO         X11, X13
        PMULULQ      X12, X13
        PSRLO        $4, X12
        PSRLO        $4, X11
        MOVO         X11, X10
        PMULULQ      X12, X10
        PSHUFD       $8, X13, X9
        PSHUFD       $8, X10, X8
        PUNPCKLLQ    X8, X9
        MOVOU        X9, ret0+32(FP)
        RET

TEXT ·shli32x4s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X14
        MOVOU        t1-32(SP), X13
        PSLLL        X14, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shri32x4s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X14
        MOVOU        t1-32(SP), X13
        PSRAL        X14, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·addu32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PADDL        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·subu32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PSUBL        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·mulu32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t2-48(SP), X12
        MOVOU        t3-64(SP), X11
        MOVO         X11, X13
        PMULULQ      X12, X13
        PSRLO        $4, X12
        PSRLO        $4, X11
        MOVO         X11, X10
        PMULULQ      X12, X10
        PSHUFD       $8, X13, X9
        PSHUFD       $8, X10, X8
        PUNPCKLLQ    X8, X9
        MOVOU        X9, ret0+32(FP)
        RET

TEXT ·shlu32x4s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X14
        MOVOU        t1-32(SP), X13
        PSLLL        X14, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·shru32x4s(SB),$56-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        X15, t1-32(SP)
        MOVB         shift+16(FP), R15
        MOVQ         R15, X14
        MOVOU        t1-32(SP), X13
        PSRLL        X14, X13
        MOVOU        X13, ret0+24(FP)
        RET

TEXT ·addi64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PADDQ        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·subi64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PSUBQ        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·addu64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PADDQ        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·subu64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVOU        y+16(FP), X14
        MOVOU        X15, t2-48(SP)
        MOVOU        X14, t3-64(SP)
        MOVOU        t3-64(SP), X13
        MOVOU        t2-48(SP), X12
        PSUBQ        X13, X12
        MOVOU        X12, ret0+32(FP)
        RET

TEXT ·addf32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVUPS       y+16(FP), X14
        MOVUPS       X15, t2-48(SP)
        MOVUPS       X14, t3-64(SP)
        MOVUPS       t3-64(SP), X13
        MOVUPS       t2-48(SP), X12
        ADDPS        X13, X12
        MOVUPS       X12, ret0+32(FP)
        RET

TEXT ·subf32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVUPS       y+16(FP), X14
        MOVUPS       X15, t2-48(SP)
        MOVUPS       X14, t3-64(SP)
        MOVUPS       t3-64(SP), X13
        MOVUPS       t2-48(SP), X12
        SUBPS        X13, X12
        MOVUPS       X12, ret0+32(FP)
        RET

TEXT ·mulf32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVUPS       y+16(FP), X14
        MOVUPS       X15, t2-48(SP)
        MOVUPS       X14, t3-64(SP)
        MOVUPS       t3-64(SP), X13
        MOVUPS       t2-48(SP), X12
        MULPS        X13, X12
        MOVUPS       X12, ret0+32(FP)
        RET

TEXT ·divf32x4s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVUPS       x+0(FP), X15
        MOVUPS       y+16(FP), X14
        MOVUPS       X15, t2-48(SP)
        MOVUPS       X14, t3-64(SP)
        MOVUPS       t3-64(SP), X13
        MOVUPS       t2-48(SP), X12
        DIVPS        X13, X12
        MOVUPS       X12, ret0+32(FP)
        RET

TEXT ·addf64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVUPD       x+0(FP), X15
        MOVUPD       y+16(FP), X14
        MOVUPD       X15, t2-48(SP)
        MOVUPD       X14, t3-64(SP)
        MOVUPD       t3-64(SP), X13
        MOVUPD       t2-48(SP), X12
        ADDPD        X13, X12
        MOVUPD       X12, ret0+32(FP)
        RET

TEXT ·subf64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVUPD       x+0(FP), X15
        MOVUPD       y+16(FP), X14
        MOVUPD       X15, t2-48(SP)
        MOVUPD       X14, t3-64(SP)
        MOVUPD       t3-64(SP), X13
        MOVUPD       t2-48(SP), X12
        SUBPD        X13, X12
        MOVUPD       X12, ret0+32(FP)
        RET

TEXT ·mulf64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVUPD       x+0(FP), X15
        MOVUPD       y+16(FP), X14
        MOVUPD       X15, t2-48(SP)
        MOVUPD       X14, t3-64(SP)
        MOVUPD       t3-64(SP), X13
        MOVUPD       t2-48(SP), X12
        MULPD        X13, X12
        MOVUPD       X12, ret0+32(FP)
        RET

TEXT ·divf64x2s(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
        MOVQ         $0, t2-48(SP)
        MOVQ         $0, t2-40(SP)
        MOVQ         $0, t3-64(SP)
        MOVQ         $0, t3-56(SP)
        MOVQ         $0, t4-80(SP)
        MOVQ         $0, t4-72(SP)
block0:
        MOVUPD       x+0(FP), X15
        MOVUPD       y+16(FP), X14
        MOVUPD       X15, t2-48(SP)
        MOVUPD       X14, t3-64(SP)
        MOVUPD       t3-64(SP), X13
        MOVUPD       t2-48(SP), X12
        DIVPD        X13, X12
        MOVUPD       X12, ret0+32(FP)
        RET

