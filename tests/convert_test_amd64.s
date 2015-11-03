// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·U8ToU8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·U8ToU16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVBWZX      R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·U8ToU32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVBLZX      R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·U8ToU64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVBQZX      R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·U8ToI8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·U8ToI16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVBWZX      R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·U8ToI32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVBLZX      R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·U8ToI64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVBQZX      R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·U8ToF32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        CVTSL2SS     R14, X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·U8ToF64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        CVTSL2SD     R14, X15
        MOVQ         X15, ret0+8(FP)
        RET

TEXT ·U16ToU8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·U16ToU16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVW         R15, ret0+8(FP)
        RET

TEXT ·U16ToU32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVWLZX      R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·U16ToU64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVWQZX      R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·U16ToI8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·U16ToI16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVW         R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·U16ToI32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVWLZX      R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·U16ToI64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVWQZX      R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·U16ToF32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        CVTSL2SS     R14, X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·U16ToF64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        CVTSL2SD     R14, X15
        MOVQ         X15, ret0+8(FP)
        RET

TEXT ·U32ToU8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·U32ToU16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVW         R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·U32ToU32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVL         R15, ret0+8(FP)
        RET

TEXT ·U32ToU64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVLQZX      R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·U32ToI8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·U32ToI16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVW         R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·U32ToI32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVL         R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·U32ToI64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVLQZX      R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·U32ToF32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        CVTSQ2SS     R14, X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·U32ToF64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        CVTSQ2SD     R14, X15
        MOVQ         X15, ret0+8(FP)
        RET

TEXT ·U64ToU8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·U64ToU16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVW         R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·U64ToU32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVL         R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·U64ToU64s(SB),$8-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVQ         R15, ret0+8(FP)
        RET

TEXT ·U64ToI8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·U64ToI16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVW         R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·U64ToI32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVL         R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·U64ToI64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVQ         R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·U64ToF32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        //           U64
        CMPQ	     R15, $-1
        // jmp to rounding
        JEQ	     lbl1
        //           U64
        CMPQ	     R15, $-1
        // jmp to no rounding
        JGE	     lbl2
        // rounding label
lbl1:
        //           U64 I64
        MOVQ	     R15, R13
        //           I64
        SHRQ	      $1, R13
        //           U64 TMP
        MOVQ	     R15, R14
        //               TMP
        ANDL	     $1, R14
        //           TMP I64
        ORQ	     R14, R13
        //CVT        I64 XMM
        CVTSQ2SS     R13, X15
        //ADD        XMM, XMM
        ADDSS        X15, X15
        // jmp to end
        JMP          lbl3
        // no rounding label
lbl2:
        //CVT        U64 XMM
        CVTSQ2SS     R15, X15
        // end label
lbl3:
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·U64ToF64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        //           U64
        CMPQ	     R15, $-1
        // jmp to rounding
        JEQ	     lbl1
        //           U64
        CMPQ	     R15, $-1
        // jmp to no rounding
        JGE	     lbl2
        // rounding label
lbl1:
        //           U64 I64
        MOVQ	     R15, R13
        //           I64
        SHRQ	      $1, R13
        //           U64 TMP
        MOVQ	     R15, R14
        //               TMP
        ANDL	     $1, R14
        //           TMP I64
        ORQ	     R14, R13
        //CVT        I64 XMM
        CVTSQ2SD     R13, X15
        //ADD        XMM, XMM
        ADDSD        X15, X15
        // jmp to end
        JMP          lbl3
        // no rounding label
lbl2:
        //CVT        U64 XMM
        CVTSQ2SD     R15, X15
        // end label
lbl3:
        MOVQ         X15, ret0+8(FP)
        RET

TEXT ·I8ToU8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·I8ToU16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVBWSX      R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·I8ToU32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVBLSX      R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·I8ToU64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVBQSX      R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·I8ToI8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·I8ToI16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVBWSX      R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·I8ToI32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVBLSX      R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·I8ToI64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        MOVBQSX      R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·I8ToF32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        CVTSL2SS     R14, X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·I8ToF64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVB         x+0(FP), R15
        CVTSL2SD     R14, X15
        MOVQ         X15, ret0+8(FP)
        RET

TEXT ·I16ToU8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·I16ToU16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVW         R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·I16ToU32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVWLSX      R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·I16ToU64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVWQSX      R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·I16ToI8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·I16ToI16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVW         R15, ret0+8(FP)
        RET

TEXT ·I16ToI32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVWLSX      R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·I16ToI64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        MOVWQSX      R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·I16ToF32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        CVTSL2SS     R14, X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·I16ToF64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVW         x+0(FP), R15
        CVTSL2SD     R14, X15
        MOVQ         X15, ret0+8(FP)
        RET

TEXT ·I32ToU8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·I32ToU16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVW         R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·I32ToU32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVL         R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·I32ToU64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVLQSX      R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·I32ToI8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·I32ToI16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVW         R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·I32ToI32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVL         R15, ret0+8(FP)
        RET

TEXT ·I32ToI64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        MOVLQSX      R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·I32ToF32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        CVTSL2SS     R15, X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·I32ToF64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), R15
        CVTSL2SD     R15, X15
        MOVQ         X15, ret0+8(FP)
        RET

TEXT ·I64ToU8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·I64ToU16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVW         R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·I64ToU32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVL         R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·I64ToU64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVQ         R15, R14
        MOVQ         R14, ret0+8(FP)
        RET

TEXT ·I64ToI8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVB         R15, R14
        MOVB         R14, ret0+8(FP)
        RET

TEXT ·I64ToI16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVW         R15, R14
        MOVW         R14, ret0+8(FP)
        RET

TEXT ·I64ToI32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVL         R15, R14
        MOVL         R14, ret0+8(FP)
        RET

TEXT ·I64ToI64s(SB),$8-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        MOVQ         R15, ret0+8(FP)
        RET

TEXT ·I64ToF32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        CVTSQ2SS     R15, X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·I64ToF64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), R15
        CVTSQ2SD     R15, X15
        MOVQ         X15, ret0+8(FP)
        RET

TEXT ·F32ToU8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), X15
        CVTTSS2SL    X15, R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·F32ToU16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), X15
        CVTTSS2SL    X15, R15
        MOVW         R15, ret0+8(FP)
        RET

TEXT ·F32ToU32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), X15
        CVTTSS2SL    X15, R15
        MOVL         R15, ret0+8(FP)
        RET

TEXT ·F32ToU64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), X15
        CVTTSS2SQ    X15, R15
        MOVQ         R15, ret0+8(FP)
        RET

TEXT ·F32ToI8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), X15
        CVTTSS2SL    X15, R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·F32ToI16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), X15
        CVTTSS2SL    X15, R15
        MOVW         R15, ret0+8(FP)
        RET

TEXT ·F32ToI32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), X15
        CVTTSS2SL    X15, R15
        MOVL         R15, ret0+8(FP)
        RET

TEXT ·F32ToI64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), X15
        CVTTSS2SQ    X15, R15
        MOVQ         R15, ret0+8(FP)
        RET

TEXT ·F32ToF32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), X15
        MOVL         X15, ret0+8(FP)
        RET

TEXT ·F32ToF64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVL         x+0(FP), X15
        CVTSS2SD     X15, X14
        MOVQ         X14, ret0+8(FP)
        RET

TEXT ·F64ToU8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), X15
        CVTTSD2SL    X15, R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·F64ToU16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), X15
        CVTTSD2SL    X15, R15
        MOVW         R15, ret0+8(FP)
        RET

TEXT ·F64ToU32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), X15
        CVTTSD2SL    X15, R15
        MOVL         R15, ret0+8(FP)
        RET

TEXT ·F64ToU64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), X15
        CVTTSD2SQ    X15, R15
        MOVQ         R15, ret0+8(FP)
        RET

TEXT ·F64ToI8s(SB),$8-9
        MOVB         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), X15
        CVTTSD2SL    X15, R15
        MOVB         R15, ret0+8(FP)
        RET

TEXT ·F64ToI16s(SB),$8-10
        MOVW         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), X15
        CVTTSD2SL    X15, R15
        MOVW         R15, ret0+8(FP)
        RET

TEXT ·F64ToI32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), X15
        CVTTSD2SL    X15, R15
        MOVL         R15, ret0+8(FP)
        RET

TEXT ·F64ToI64s(SB),$16-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), X15
        CVTTSD2SQ    X15, R15
        MOVQ         R15, ret0+8(FP)
        RET

TEXT ·F64ToF32s(SB),$8-12
        MOVL         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), X15
        CVTSD2SS     X15, X14
        MOVL         X14, ret0+8(FP)
        RET

TEXT ·F64ToF64s(SB),$8-16
        MOVQ         $0, ret0+8(FP)
block0:
        MOVQ         x+0(FP), X15
        MOVQ         X15, ret0+8(FP)
        RET

