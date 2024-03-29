// Code generated by command: go run addconst_asm.go -out ../../matfuncs/addconst_amd64.s -stubs ../../matfuncs/addconst_amd64_stubs.go -pkg matfuncs. DO NOT EDIT.

//go:build amd64 && gc && !purego

#include "textflag.h"

// func AddConstAVX32(c float32, x []float32, y []float32)
// Requires: AVX, AVX2, SSE
TEXT ·AddConstAVX32(SB), NOSPLIT, $0-56
	MOVSS        c+0(FP), X0
	MOVQ         x_base+8(FP), AX
	MOVQ         y_base+32(FP), CX
	MOVQ         x_len+16(FP), DX
	VBROADCASTSS X0, Y1

unrolledLoop14:
	CMPQ    DX, $0x00000070
	JL      unrolledLoop8
	VADDPS  (AX), Y1, Y2
	VADDPS  32(AX), Y1, Y3
	VADDPS  64(AX), Y1, Y4
	VADDPS  96(AX), Y1, Y5
	VADDPS  128(AX), Y1, Y6
	VADDPS  160(AX), Y1, Y7
	VADDPS  192(AX), Y1, Y8
	VADDPS  224(AX), Y1, Y9
	VADDPS  256(AX), Y1, Y10
	VADDPS  288(AX), Y1, Y11
	VADDPS  320(AX), Y1, Y12
	VADDPS  352(AX), Y1, Y13
	VADDPS  384(AX), Y1, Y14
	VADDPS  416(AX), Y1, Y15
	VMOVUPS Y2, (CX)
	VMOVUPS Y3, 32(CX)
	VMOVUPS Y4, 64(CX)
	VMOVUPS Y5, 96(CX)
	VMOVUPS Y6, 128(CX)
	VMOVUPS Y7, 160(CX)
	VMOVUPS Y8, 192(CX)
	VMOVUPS Y9, 224(CX)
	VMOVUPS Y10, 256(CX)
	VMOVUPS Y11, 288(CX)
	VMOVUPS Y12, 320(CX)
	VMOVUPS Y13, 352(CX)
	VMOVUPS Y14, 384(CX)
	VMOVUPS Y15, 416(CX)
	ADDQ    $0x000001c0, AX
	ADDQ    $0x000001c0, CX
	SUBQ    $0x00000070, DX
	JMP     unrolledLoop14

unrolledLoop8:
	CMPQ    DX, $0x00000040
	JL      unrolledLoop4
	VADDPS  (AX), Y1, Y2
	VADDPS  32(AX), Y1, Y3
	VADDPS  64(AX), Y1, Y4
	VADDPS  96(AX), Y1, Y5
	VADDPS  128(AX), Y1, Y6
	VADDPS  160(AX), Y1, Y7
	VADDPS  192(AX), Y1, Y8
	VADDPS  224(AX), Y1, Y9
	VMOVUPS Y2, (CX)
	VMOVUPS Y3, 32(CX)
	VMOVUPS Y4, 64(CX)
	VMOVUPS Y5, 96(CX)
	VMOVUPS Y6, 128(CX)
	VMOVUPS Y7, 160(CX)
	VMOVUPS Y8, 192(CX)
	VMOVUPS Y9, 224(CX)
	ADDQ    $0x00000100, AX
	ADDQ    $0x00000100, CX
	SUBQ    $0x00000040, DX
	JMP     unrolledLoop8

unrolledLoop4:
	CMPQ    DX, $0x00000020
	JL      unrolledLoop1
	VADDPS  (AX), Y1, Y2
	VADDPS  32(AX), Y1, Y3
	VADDPS  64(AX), Y1, Y4
	VADDPS  96(AX), Y1, Y5
	VMOVUPS Y2, (CX)
	VMOVUPS Y3, 32(CX)
	VMOVUPS Y4, 64(CX)
	VMOVUPS Y5, 96(CX)
	ADDQ    $0x00000080, AX
	ADDQ    $0x00000080, CX
	SUBQ    $0x00000020, DX
	JMP     unrolledLoop4

unrolledLoop1:
	CMPQ    DX, $0x00000008
	JL      tailLoop
	VADDPS  (AX), Y1, Y2
	VMOVUPS Y2, (CX)
	ADDQ    $0x00000020, AX
	ADDQ    $0x00000020, CX
	SUBQ    $0x00000008, DX
	JMP     unrolledLoop1

tailLoop:
	CMPQ  DX, $0x00000000
	JE    end
	MOVSS (AX), X1
	ADDSS X0, X1
	MOVSS X1, (CX)
	ADDQ  $0x00000004, AX
	ADDQ  $0x00000004, CX
	DECQ  DX
	JMP   tailLoop

end:
	RET

// func AddConstAVX64(c float64, x []float64, y []float64)
// Requires: AVX, AVX2, SSE2
TEXT ·AddConstAVX64(SB), NOSPLIT, $0-56
	MOVSD        c+0(FP), X0
	MOVQ         x_base+8(FP), AX
	MOVQ         y_base+32(FP), CX
	MOVQ         x_len+16(FP), DX
	VBROADCASTSD X0, Y1

unrolledLoop14:
	CMPQ    DX, $0x00000038
	JL      unrolledLoop8
	VADDPD  (AX), Y1, Y2
	VADDPD  32(AX), Y1, Y3
	VADDPD  64(AX), Y1, Y4
	VADDPD  96(AX), Y1, Y5
	VADDPD  128(AX), Y1, Y6
	VADDPD  160(AX), Y1, Y7
	VADDPD  192(AX), Y1, Y8
	VADDPD  224(AX), Y1, Y9
	VADDPD  256(AX), Y1, Y10
	VADDPD  288(AX), Y1, Y11
	VADDPD  320(AX), Y1, Y12
	VADDPD  352(AX), Y1, Y13
	VADDPD  384(AX), Y1, Y14
	VADDPD  416(AX), Y1, Y15
	VMOVUPD Y2, (CX)
	VMOVUPD Y3, 32(CX)
	VMOVUPD Y4, 64(CX)
	VMOVUPD Y5, 96(CX)
	VMOVUPD Y6, 128(CX)
	VMOVUPD Y7, 160(CX)
	VMOVUPD Y8, 192(CX)
	VMOVUPD Y9, 224(CX)
	VMOVUPD Y10, 256(CX)
	VMOVUPD Y11, 288(CX)
	VMOVUPD Y12, 320(CX)
	VMOVUPD Y13, 352(CX)
	VMOVUPD Y14, 384(CX)
	VMOVUPD Y15, 416(CX)
	ADDQ    $0x000001c0, AX
	ADDQ    $0x000001c0, CX
	SUBQ    $0x00000038, DX
	JMP     unrolledLoop14

unrolledLoop8:
	CMPQ    DX, $0x00000020
	JL      unrolledLoop4
	VADDPD  (AX), Y1, Y2
	VADDPD  32(AX), Y1, Y3
	VADDPD  64(AX), Y1, Y4
	VADDPD  96(AX), Y1, Y5
	VADDPD  128(AX), Y1, Y6
	VADDPD  160(AX), Y1, Y7
	VADDPD  192(AX), Y1, Y8
	VADDPD  224(AX), Y1, Y9
	VMOVUPD Y2, (CX)
	VMOVUPD Y3, 32(CX)
	VMOVUPD Y4, 64(CX)
	VMOVUPD Y5, 96(CX)
	VMOVUPD Y6, 128(CX)
	VMOVUPD Y7, 160(CX)
	VMOVUPD Y8, 192(CX)
	VMOVUPD Y9, 224(CX)
	ADDQ    $0x00000100, AX
	ADDQ    $0x00000100, CX
	SUBQ    $0x00000020, DX
	JMP     unrolledLoop8

unrolledLoop4:
	CMPQ    DX, $0x00000010
	JL      unrolledLoop1
	VADDPD  (AX), Y1, Y2
	VADDPD  32(AX), Y1, Y3
	VADDPD  64(AX), Y1, Y4
	VADDPD  96(AX), Y1, Y5
	VMOVUPD Y2, (CX)
	VMOVUPD Y3, 32(CX)
	VMOVUPD Y4, 64(CX)
	VMOVUPD Y5, 96(CX)
	ADDQ    $0x00000080, AX
	ADDQ    $0x00000080, CX
	SUBQ    $0x00000010, DX
	JMP     unrolledLoop4

unrolledLoop1:
	CMPQ    DX, $0x00000004
	JL      tailLoop
	VADDPD  (AX), Y1, Y2
	VMOVUPD Y2, (CX)
	ADDQ    $0x00000020, AX
	ADDQ    $0x00000020, CX
	SUBQ    $0x00000004, DX
	JMP     unrolledLoop1

tailLoop:
	CMPQ  DX, $0x00000000
	JE    end
	MOVSD (AX), X1
	ADDSD X0, X1
	MOVSD X1, (CX)
	ADDQ  $0x00000008, AX
	ADDQ  $0x00000008, CX
	DECQ  DX
	JMP   tailLoop

end:
	RET

// func AddConstSSE32(c float32, x []float32, y []float32)
// Requires: SSE
TEXT ·AddConstSSE32(SB), NOSPLIT, $0-56
	MOVSS  c+0(FP), X0
	MOVQ   x_base+8(FP), AX
	MOVQ   y_base+32(FP), CX
	MOVQ   x_len+16(FP), DX
	SHUFPS $0x00, X0, X0

unrolledLoop14:
	CMPQ   DX, $0x00000038
	JL     unrolledLoop8
	MOVUPS (AX), X1
	MOVUPS 16(AX), X2
	MOVUPS 32(AX), X3
	MOVUPS 48(AX), X4
	MOVUPS 64(AX), X5
	MOVUPS 80(AX), X6
	MOVUPS 96(AX), X7
	MOVUPS 112(AX), X8
	MOVUPS 128(AX), X9
	MOVUPS 144(AX), X10
	MOVUPS 160(AX), X11
	MOVUPS 176(AX), X12
	MOVUPS 192(AX), X13
	MOVUPS 208(AX), X14
	ADDPS  X0, X1
	ADDPS  X0, X2
	ADDPS  X0, X3
	ADDPS  X0, X4
	ADDPS  X0, X5
	ADDPS  X0, X6
	ADDPS  X0, X7
	ADDPS  X0, X8
	ADDPS  X0, X9
	ADDPS  X0, X10
	ADDPS  X0, X11
	ADDPS  X0, X12
	ADDPS  X0, X13
	ADDPS  X0, X14
	MOVUPS X1, (CX)
	MOVUPS X2, 16(CX)
	MOVUPS X3, 32(CX)
	MOVUPS X4, 48(CX)
	MOVUPS X5, 64(CX)
	MOVUPS X6, 80(CX)
	MOVUPS X7, 96(CX)
	MOVUPS X8, 112(CX)
	MOVUPS X9, 128(CX)
	MOVUPS X10, 144(CX)
	MOVUPS X11, 160(CX)
	MOVUPS X12, 176(CX)
	MOVUPS X13, 192(CX)
	MOVUPS X14, 208(CX)
	ADDQ   $0x000000e0, AX
	ADDQ   $0x000000e0, CX
	SUBQ   $0x00000038, DX
	JMP    unrolledLoop14

unrolledLoop8:
	CMPQ   DX, $0x00000020
	JL     unrolledLoop4
	MOVUPS (AX), X1
	MOVUPS 16(AX), X2
	MOVUPS 32(AX), X3
	MOVUPS 48(AX), X4
	MOVUPS 64(AX), X5
	MOVUPS 80(AX), X6
	MOVUPS 96(AX), X7
	MOVUPS 112(AX), X8
	ADDPS  X0, X1
	ADDPS  X0, X2
	ADDPS  X0, X3
	ADDPS  X0, X4
	ADDPS  X0, X5
	ADDPS  X0, X6
	ADDPS  X0, X7
	ADDPS  X0, X8
	MOVUPS X1, (CX)
	MOVUPS X2, 16(CX)
	MOVUPS X3, 32(CX)
	MOVUPS X4, 48(CX)
	MOVUPS X5, 64(CX)
	MOVUPS X6, 80(CX)
	MOVUPS X7, 96(CX)
	MOVUPS X8, 112(CX)
	ADDQ   $0x00000080, AX
	ADDQ   $0x00000080, CX
	SUBQ   $0x00000020, DX
	JMP    unrolledLoop8

unrolledLoop4:
	CMPQ   DX, $0x00000010
	JL     unrolledLoop1
	MOVUPS (AX), X1
	MOVUPS 16(AX), X2
	MOVUPS 32(AX), X3
	MOVUPS 48(AX), X4
	ADDPS  X0, X1
	ADDPS  X0, X2
	ADDPS  X0, X3
	ADDPS  X0, X4
	MOVUPS X1, (CX)
	MOVUPS X2, 16(CX)
	MOVUPS X3, 32(CX)
	MOVUPS X4, 48(CX)
	ADDQ   $0x00000040, AX
	ADDQ   $0x00000040, CX
	SUBQ   $0x00000010, DX
	JMP    unrolledLoop4

unrolledLoop1:
	CMPQ   DX, $0x00000004
	JL     tailLoop
	MOVUPS (AX), X1
	ADDPS  X0, X1
	MOVUPS X1, (CX)
	ADDQ   $0x00000010, AX
	ADDQ   $0x00000010, CX
	SUBQ   $0x00000004, DX
	JMP    unrolledLoop1

tailLoop:
	CMPQ  DX, $0x00000000
	JE    end
	MOVSS (AX), X1
	ADDSS X0, X1
	MOVSS X1, (CX)
	ADDQ  $0x00000004, AX
	ADDQ  $0x00000004, CX
	DECQ  DX
	JMP   tailLoop

end:
	RET

// func AddConstSSE64(c float64, x []float64, y []float64)
// Requires: SSE2
TEXT ·AddConstSSE64(SB), NOSPLIT, $0-56
	MOVSD  c+0(FP), X0
	MOVQ   x_base+8(FP), AX
	MOVQ   y_base+32(FP), CX
	MOVQ   x_len+16(FP), DX
	SHUFPD $0x00, X0, X0

unrolledLoop14:
	CMPQ   DX, $0x0000001c
	JL     unrolledLoop8
	MOVUPD (AX), X1
	MOVUPD 16(AX), X2
	MOVUPD 32(AX), X3
	MOVUPD 48(AX), X4
	MOVUPD 64(AX), X5
	MOVUPD 80(AX), X6
	MOVUPD 96(AX), X7
	MOVUPD 112(AX), X8
	MOVUPD 128(AX), X9
	MOVUPD 144(AX), X10
	MOVUPD 160(AX), X11
	MOVUPD 176(AX), X12
	MOVUPD 192(AX), X13
	MOVUPD 208(AX), X14
	ADDPD  X0, X1
	ADDPD  X0, X2
	ADDPD  X0, X3
	ADDPD  X0, X4
	ADDPD  X0, X5
	ADDPD  X0, X6
	ADDPD  X0, X7
	ADDPD  X0, X8
	ADDPD  X0, X9
	ADDPD  X0, X10
	ADDPD  X0, X11
	ADDPD  X0, X12
	ADDPD  X0, X13
	ADDPD  X0, X14
	MOVUPD X1, (CX)
	MOVUPD X2, 16(CX)
	MOVUPD X3, 32(CX)
	MOVUPD X4, 48(CX)
	MOVUPD X5, 64(CX)
	MOVUPD X6, 80(CX)
	MOVUPD X7, 96(CX)
	MOVUPD X8, 112(CX)
	MOVUPD X9, 128(CX)
	MOVUPD X10, 144(CX)
	MOVUPD X11, 160(CX)
	MOVUPD X12, 176(CX)
	MOVUPD X13, 192(CX)
	MOVUPD X14, 208(CX)
	ADDQ   $0x000000e0, AX
	ADDQ   $0x000000e0, CX
	SUBQ   $0x0000001c, DX
	JMP    unrolledLoop14

unrolledLoop8:
	CMPQ   DX, $0x00000010
	JL     unrolledLoop4
	MOVUPD (AX), X1
	MOVUPD 16(AX), X2
	MOVUPD 32(AX), X3
	MOVUPD 48(AX), X4
	MOVUPD 64(AX), X5
	MOVUPD 80(AX), X6
	MOVUPD 96(AX), X7
	MOVUPD 112(AX), X8
	ADDPD  X0, X1
	ADDPD  X0, X2
	ADDPD  X0, X3
	ADDPD  X0, X4
	ADDPD  X0, X5
	ADDPD  X0, X6
	ADDPD  X0, X7
	ADDPD  X0, X8
	MOVUPD X1, (CX)
	MOVUPD X2, 16(CX)
	MOVUPD X3, 32(CX)
	MOVUPD X4, 48(CX)
	MOVUPD X5, 64(CX)
	MOVUPD X6, 80(CX)
	MOVUPD X7, 96(CX)
	MOVUPD X8, 112(CX)
	ADDQ   $0x00000080, AX
	ADDQ   $0x00000080, CX
	SUBQ   $0x00000010, DX
	JMP    unrolledLoop8

unrolledLoop4:
	CMPQ   DX, $0x00000008
	JL     unrolledLoop1
	MOVUPD (AX), X1
	MOVUPD 16(AX), X2
	MOVUPD 32(AX), X3
	MOVUPD 48(AX), X4
	ADDPD  X0, X1
	ADDPD  X0, X2
	ADDPD  X0, X3
	ADDPD  X0, X4
	MOVUPD X1, (CX)
	MOVUPD X2, 16(CX)
	MOVUPD X3, 32(CX)
	MOVUPD X4, 48(CX)
	ADDQ   $0x00000040, AX
	ADDQ   $0x00000040, CX
	SUBQ   $0x00000008, DX
	JMP    unrolledLoop4

unrolledLoop1:
	CMPQ   DX, $0x00000002
	JL     tailLoop
	MOVUPD (AX), X1
	ADDPD  X0, X1
	MOVUPD X1, (CX)
	ADDQ   $0x00000010, AX
	ADDQ   $0x00000010, CX
	SUBQ   $0x00000002, DX
	JMP    unrolledLoop1

tailLoop:
	CMPQ  DX, $0x00000000
	JE    end
	MOVSD (AX), X1
	ADDSD X0, X1
	MOVSD X1, (CX)
	ADDQ  $0x00000008, AX
	ADDQ  $0x00000008, CX
	DECQ  DX
	JMP   tailLoop

end:
	RET
