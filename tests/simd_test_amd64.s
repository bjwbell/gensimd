// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·addi8x16s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PADDB        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·subi8x16s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PSUBB        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·addu8x16s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PADDB        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·subu8x16s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PSUBB        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·addi16x8s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PADDW        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·subi16x8s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PSUBW        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·muli16x8s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PMULLW       X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·shli16x8s(SB),$24-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
block0:
        MOVB         shift+16(FP), R15
        MOVQ         R15, X15
        MOVOU        x+0(FP), X14
        PSLLW        X15, X14
        MOVOU        X14, ret0+24(FP)
        RET

TEXT ·shri16x8s(SB),$24-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
block0:
        MOVB         shift+16(FP), R15
        MOVQ         R15, X15
        MOVOU        x+0(FP), X14
        PSRAW        X15, X14
        MOVOU        X14, ret0+24(FP)
        RET

TEXT ·addu16x8s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PADDW        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·subu16x8s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PSUBW        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·mulu16x8s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PMULLW       X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·shlu16x8s(SB),$24-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
block0:
        MOVB         shift+16(FP), R15
        MOVQ         R15, X15
        MOVOU        x+0(FP), X14
        PSLLW        X15, X14
        MOVOU        X14, ret0+24(FP)
        RET

TEXT ·shru16x8s(SB),$24-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
block0:
        MOVOU        x+0(FP), X15
        MOVB         shift+16(FP), R15
        PEXTRW       $0, X15, R14
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
        PINSRW       $0, R14, X15
        PEXTRW       $1, X15, R14
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
        PINSRW       $1, R14, X15
        PEXTRW       $2, X15, R14
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
        PINSRW       $2, R14, X15
        PEXTRW       $3, X15, R14
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
        PINSRW       $3, R14, X15
        PEXTRW       $4, X15, R14
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
        PINSRW       $4, R14, X15
        PEXTRW       $5, X15, R14
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
        PINSRW       $5, R14, X15
        PEXTRW       $6, X15, R14
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
        PINSRW       $6, R14, X15
        PEXTRW       $7, X15, R14
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
        PINSRW       $7, R14, X15
        MOVOU        X15, ret0+24(FP)
        RET

TEXT ·addi32x4s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PADDL        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·subi32x4s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PSUBL        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·muli32x4s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        x+0(FP), X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X15
        PMULULQ      X14, X15
        PSRLO        $4, X14
        PSRLO        $4, X13
        MOVO         X13, X12
        PMULULQ      X14, X12
        PSHUFD       $8, X15, X11
        PSHUFD       $8, X12, X10
        PUNPCKLLQ    X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·shli32x4s(SB),$24-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
block0:
        MOVB         shift+16(FP), R15
        MOVQ         R15, X15
        MOVOU        x+0(FP), X14
        PSLLL        X15, X14
        MOVOU        X14, ret0+24(FP)
        RET

TEXT ·shri32x4s(SB),$24-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
block0:
        MOVB         shift+16(FP), R15
        MOVQ         R15, X15
        MOVOU        x+0(FP), X14
        PSRAL        X15, X14
        MOVOU        X14, ret0+24(FP)
        RET

TEXT ·addu32x4s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PADDL        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·subu32x4s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PSUBL        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·mulu32x4s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        x+0(FP), X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X15
        PMULULQ      X14, X15
        PSRLO        $4, X14
        PSRLO        $4, X13
        MOVO         X13, X12
        PMULULQ      X14, X12
        PSHUFD       $8, X15, X11
        PSHUFD       $8, X12, X10
        PUNPCKLLQ    X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

TEXT ·shlu32x4s(SB),$24-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
block0:
        MOVB         shift+16(FP), R15
        MOVQ         R15, X15
        MOVOU        x+0(FP), X14
        PSLLL        X15, X14
        MOVOU        X14, ret0+24(FP)
        RET

TEXT ·shru32x4s(SB),$24-40
        MOVQ         $0, ret0+24(FP)
        MOVQ         $0, ret0+32(FP)
block0:
        MOVB         shift+16(FP), R15
        MOVQ         R15, X15
        MOVOU        x+0(FP), X14
        PSRLL        X15, X14
        MOVOU        X14, ret0+24(FP)
        RET

TEXT ·addi64x2s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PADDQ        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·subi64x2s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PSUBQ        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·addu64x2s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PADDQ        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·subu64x2s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVOU        y+16(FP), X15
        MOVOU        x+0(FP), X14
        PSUBQ        X15, X14
        MOVOU        X14, ret0+32(FP)
        RET

TEXT ·addf32x4s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVUPS       y+16(FP), X15
        MOVUPS       x+0(FP), X14
        ADDPS        X15, X14
        MOVUPS       X14, ret0+32(FP)
        RET

TEXT ·subf32x4s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVUPS       y+16(FP), X15
        MOVUPS       x+0(FP), X14
        SUBPS        X15, X14
        MOVUPS       X14, ret0+32(FP)
        RET

TEXT ·mulf32x4s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVUPS       y+16(FP), X15
        MOVUPS       x+0(FP), X14
        MULPS        X15, X14
        MOVUPS       X14, ret0+32(FP)
        RET

TEXT ·divf32x4s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVUPS       y+16(FP), X15
        MOVUPS       x+0(FP), X14
        DIVPS        X15, X14
        MOVUPS       X14, ret0+32(FP)
        RET

TEXT ·addf64x2s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVUPD       y+16(FP), X15
        MOVUPD       x+0(FP), X14
        ADDPD        X15, X14
        MOVUPD       X14, ret0+32(FP)
        RET

TEXT ·subf64x2s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVUPD       y+16(FP), X15
        MOVUPD       x+0(FP), X14
        SUBPD        X15, X14
        MOVUPD       X14, ret0+32(FP)
        RET

TEXT ·mulf64x2s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVUPD       y+16(FP), X15
        MOVUPD       x+0(FP), X14
        MULPD        X15, X14
        MOVUPD       X14, ret0+32(FP)
        RET

TEXT ·divf64x2s(SB),$24-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
block0:
        MOVUPD       y+16(FP), X15
        MOVUPD       x+0(FP), X14
        DIVPD        X15, X14
        MOVUPD       X14, ret0+32(FP)
        RET

