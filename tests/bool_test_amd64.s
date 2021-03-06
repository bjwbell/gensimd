// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·boolt0s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVBQZX      x+0(FP), R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·boolt1s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVBQZX      x+0(FP), R15
        XORQ         $1, R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·boolt2s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVBQZX      x+0(FP), R15
        MOVB         $0, R14
        MOVB         R14, t0-1(SP)
        CMPB         R15, $0
        JEQ          block2
        JMP          block1
block1:
        MOVBQZX      y+1(FP), R15
        MOVB         R15, t0-1(SP)
        JMP block2
block2:
        MOVBQZX      t0-1(SP), R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·boolt3s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVBQZX      x+0(FP), R15
        CMPB         R15, $0
        JEQ          block1
        MOVB         $1, R15
        MOVB         R15, t0-1(SP)
        JMP          block2
block1:
        MOVBQZX      y+1(FP), R15
        MOVB         R15, t0-1(SP)
        JMP block2
block2:
        MOVBQZX      t0-1(SP), R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·boolt4s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVBQZX      x+0(FP), R15
        CMPB         R15, $0
        JEQ          block1
        MOVB         $1, R15
        MOVB         R15, t0-1(SP)
        JMP          block2
block1:
        MOVBQZX      y+1(FP), R15
        MOVB         R15, t0-1(SP)
        JMP block2
block2:
        MOVBQZX      t0-1(SP), R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·boolt5s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVBQZX      x+0(FP), R15
        CMPB         R15, $0
        JEQ          block1
        MOVB         $1, R15
        MOVB         R15, t0-1(SP)
        JMP          block2
block1:
        MOVBQZX      y+1(FP), R15
        MOVB         R15, t0-1(SP)
        JMP block2
block2:
        MOVBQZX      t0-1(SP), R15
        MOVB         R15, ret0+8(FP)
        RET

