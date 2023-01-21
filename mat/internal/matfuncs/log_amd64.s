// Code generated by command: go run log_asm.go -out ../../matfuncs/log_amd64.s -stubs ../../matfuncs/log_amd64_stubs.go -pkg matfuncs. DO NOT EDIT.

//go:build amd64 && gc && !purego

#include "textflag.h"

DATA AVX2_LCPI0_0<>+0(SB)/4, $0x00800000
GLOBL AVX2_LCPI0_0<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_1<>+0(SB)/4, $0x807fffff
GLOBL AVX2_LCPI0_1<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_2<>+0(SB)/4, $0x3f000000
GLOBL AVX2_LCPI0_2<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_3<>+0(SB)/4, $0xffffff81
GLOBL AVX2_LCPI0_3<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_4<>+0(SB)/4, $0x3f800000
GLOBL AVX2_LCPI0_4<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_5<>+0(SB)/4, $0x3f3504f3
GLOBL AVX2_LCPI0_5<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_6<>+0(SB)/4, $0xbf800000
GLOBL AVX2_LCPI0_6<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_7<>+0(SB)/4, $0x3d9021bb
GLOBL AVX2_LCPI0_7<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_8<>+0(SB)/4, $0xbdebd1b8
GLOBL AVX2_LCPI0_8<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_9<>+0(SB)/4, $0x3def251a
GLOBL AVX2_LCPI0_9<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_10<>+0(SB)/4, $0xbdfe5d4f
GLOBL AVX2_LCPI0_10<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_11<>+0(SB)/4, $0x3e11e9bf
GLOBL AVX2_LCPI0_11<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_12<>+0(SB)/4, $0xbe2aae50
GLOBL AVX2_LCPI0_12<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_13<>+0(SB)/4, $0x3e4cceac
GLOBL AVX2_LCPI0_13<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_14<>+0(SB)/4, $0xbe7ffffc
GLOBL AVX2_LCPI0_14<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_15<>+0(SB)/4, $0x3eaaaaaa
GLOBL AVX2_LCPI0_15<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_16<>+0(SB)/4, $0xb95e8083
GLOBL AVX2_LCPI0_16<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_17<>+0(SB)/4, $0xbf000000
GLOBL AVX2_LCPI0_17<>(SB), RODATA|NOPTR, $4

DATA AVX2_LCPI0_18<>+0(SB)/4, $0x3f318000
GLOBL AVX2_LCPI0_18<>(SB), RODATA|NOPTR, $4

// func LogAVX32(x []float32, y []float32)
// Requires: AVX, AVX2
TEXT ·LogAVX32(SB), NOSPLIT, $0-48
	MOVQ         x_base+0(FP), AX
	MOVQ         y_base+24(FP), CX
	VMOVUPS      (AX), Y0
	VXORPS       X1, X1, X1
	VCMPPS       $0x02, Y1, Y0, Y1
	VBROADCASTSS AVX2_LCPI0_0<>+0(SB), Y2
	VMAXPS       Y2, Y0, Y0
	VPSRLD       $0x17, Y0, Y2
	VBROADCASTSS AVX2_LCPI0_1<>+0(SB), Y3
	VANDPS       Y3, Y0, Y0
	VBROADCASTSS AVX2_LCPI0_2<>+0(SB), Y3
	VPBROADCASTD AVX2_LCPI0_3<>+0(SB), Y4
	VORPS        Y3, Y0, Y0
	VPADDD       Y4, Y2, Y2
	VCVTDQ2PS    Y2, Y2
	VBROADCASTSS AVX2_LCPI0_4<>+0(SB), Y3
	VADDPS       Y3, Y2, Y2
	VBROADCASTSS AVX2_LCPI0_5<>+0(SB), Y4
	VCMPPS       $0x01, Y4, Y0, Y4
	VANDPS       Y0, Y4, Y5
	VBROADCASTSS AVX2_LCPI0_6<>+0(SB), Y6
	VADDPS       Y6, Y0, Y0
	VADDPS       Y5, Y0, Y0
	VANDPS       Y3, Y4, Y3
	VSUBPS       Y3, Y2, Y2
	VMULPS       Y0, Y0, Y3
	VBROADCASTSS AVX2_LCPI0_7<>+0(SB), Y4
	VMULPS       Y4, Y0, Y4
	VBROADCASTSS AVX2_LCPI0_8<>+0(SB), Y5
	VADDPS       Y5, Y4, Y4
	VMULPS       Y4, Y0, Y4
	VBROADCASTSS AVX2_LCPI0_9<>+0(SB), Y5
	VADDPS       Y5, Y4, Y4
	VMULPS       Y4, Y0, Y4
	VBROADCASTSS AVX2_LCPI0_10<>+0(SB), Y5
	VADDPS       Y5, Y4, Y4
	VMULPS       Y4, Y0, Y4
	VBROADCASTSS AVX2_LCPI0_11<>+0(SB), Y5
	VADDPS       Y5, Y4, Y4
	VMULPS       Y4, Y0, Y4
	VBROADCASTSS AVX2_LCPI0_12<>+0(SB), Y5
	VADDPS       Y5, Y4, Y4
	VMULPS       Y4, Y0, Y4
	VBROADCASTSS AVX2_LCPI0_13<>+0(SB), Y5
	VADDPS       Y5, Y4, Y4
	VMULPS       Y4, Y0, Y4
	VBROADCASTSS AVX2_LCPI0_14<>+0(SB), Y5
	VADDPS       Y5, Y4, Y4
	VMULPS       Y4, Y0, Y4
	VBROADCASTSS AVX2_LCPI0_15<>+0(SB), Y5
	VADDPS       Y5, Y4, Y4
	VMULPS       Y4, Y0, Y4
	VMULPS       Y4, Y3, Y4
	VBROADCASTSS AVX2_LCPI0_16<>+0(SB), Y5
	VMULPS       Y5, Y2, Y5
	VADDPS       Y5, Y4, Y4
	VBROADCASTSS AVX2_LCPI0_17<>+0(SB), Y5
	VMULPS       Y5, Y3, Y3
	VADDPS       Y3, Y4, Y3
	VBROADCASTSS AVX2_LCPI0_18<>+0(SB), Y4
	VMULPS       Y4, Y2, Y2
	VADDPS       Y3, Y0, Y0
	VADDPS       Y0, Y2, Y0
	VORPS        Y0, Y1, Y0
	VMOVUPS      Y0, (CX)
	RET

DATA SSE_LCPI0_0<>+0(SB)/4, $0x00800000
DATA SSE_LCPI0_0<>+4(SB)/4, $0x00800000
DATA SSE_LCPI0_0<>+8(SB)/4, $0x00800000
DATA SSE_LCPI0_0<>+12(SB)/4, $0x00800000
GLOBL SSE_LCPI0_0<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_1<>+0(SB)/4, $0x807fffff
DATA SSE_LCPI0_1<>+4(SB)/4, $0x807fffff
DATA SSE_LCPI0_1<>+8(SB)/4, $0x807fffff
DATA SSE_LCPI0_1<>+12(SB)/4, $0x807fffff
GLOBL SSE_LCPI0_1<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_2<>+0(SB)/4, $0x3f000000
DATA SSE_LCPI0_2<>+4(SB)/4, $0x3f000000
DATA SSE_LCPI0_2<>+8(SB)/4, $0x3f000000
DATA SSE_LCPI0_2<>+12(SB)/4, $0x3f000000
GLOBL SSE_LCPI0_2<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_3<>+0(SB)/4, $0xffffff81
DATA SSE_LCPI0_3<>+4(SB)/4, $0xffffff81
DATA SSE_LCPI0_3<>+8(SB)/4, $0xffffff81
DATA SSE_LCPI0_3<>+12(SB)/4, $0xffffff81
GLOBL SSE_LCPI0_3<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_4<>+0(SB)/4, $0x3f800000
DATA SSE_LCPI0_4<>+4(SB)/4, $0x3f800000
DATA SSE_LCPI0_4<>+8(SB)/4, $0x3f800000
DATA SSE_LCPI0_4<>+12(SB)/4, $0x3f800000
GLOBL SSE_LCPI0_4<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_5<>+0(SB)/4, $0x3f3504f3
DATA SSE_LCPI0_5<>+4(SB)/4, $0x3f3504f3
DATA SSE_LCPI0_5<>+8(SB)/4, $0x3f3504f3
DATA SSE_LCPI0_5<>+12(SB)/4, $0x3f3504f3
GLOBL SSE_LCPI0_5<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_6<>+0(SB)/4, $0xbf800000
DATA SSE_LCPI0_6<>+4(SB)/4, $0xbf800000
DATA SSE_LCPI0_6<>+8(SB)/4, $0xbf800000
DATA SSE_LCPI0_6<>+12(SB)/4, $0xbf800000
GLOBL SSE_LCPI0_6<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_7<>+0(SB)/4, $0x3d9021bb
DATA SSE_LCPI0_7<>+4(SB)/4, $0x3d9021bb
DATA SSE_LCPI0_7<>+8(SB)/4, $0x3d9021bb
DATA SSE_LCPI0_7<>+12(SB)/4, $0x3d9021bb
GLOBL SSE_LCPI0_7<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_8<>+0(SB)/4, $0xbdebd1b8
DATA SSE_LCPI0_8<>+4(SB)/4, $0xbdebd1b8
DATA SSE_LCPI0_8<>+8(SB)/4, $0xbdebd1b8
DATA SSE_LCPI0_8<>+12(SB)/4, $0xbdebd1b8
GLOBL SSE_LCPI0_8<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_9<>+0(SB)/4, $0x3def251a
DATA SSE_LCPI0_9<>+4(SB)/4, $0x3def251a
DATA SSE_LCPI0_9<>+8(SB)/4, $0x3def251a
DATA SSE_LCPI0_9<>+12(SB)/4, $0x3def251a
GLOBL SSE_LCPI0_9<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_10<>+0(SB)/4, $0xbdfe5d4f
DATA SSE_LCPI0_10<>+4(SB)/4, $0xbdfe5d4f
DATA SSE_LCPI0_10<>+8(SB)/4, $0xbdfe5d4f
DATA SSE_LCPI0_10<>+12(SB)/4, $0xbdfe5d4f
GLOBL SSE_LCPI0_10<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_11<>+0(SB)/4, $0x3e11e9bf
DATA SSE_LCPI0_11<>+4(SB)/4, $0x3e11e9bf
DATA SSE_LCPI0_11<>+8(SB)/4, $0x3e11e9bf
DATA SSE_LCPI0_11<>+12(SB)/4, $0x3e11e9bf
GLOBL SSE_LCPI0_11<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_12<>+0(SB)/4, $0xbe2aae50
DATA SSE_LCPI0_12<>+4(SB)/4, $0xbe2aae50
DATA SSE_LCPI0_12<>+8(SB)/4, $0xbe2aae50
DATA SSE_LCPI0_12<>+12(SB)/4, $0xbe2aae50
GLOBL SSE_LCPI0_12<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_13<>+0(SB)/4, $0x3e4cceac
DATA SSE_LCPI0_13<>+4(SB)/4, $0x3e4cceac
DATA SSE_LCPI0_13<>+8(SB)/4, $0x3e4cceac
DATA SSE_LCPI0_13<>+12(SB)/4, $0x3e4cceac
GLOBL SSE_LCPI0_13<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_14<>+0(SB)/4, $0xbe7ffffc
DATA SSE_LCPI0_14<>+4(SB)/4, $0xbe7ffffc
DATA SSE_LCPI0_14<>+8(SB)/4, $0xbe7ffffc
DATA SSE_LCPI0_14<>+12(SB)/4, $0xbe7ffffc
GLOBL SSE_LCPI0_14<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_15<>+0(SB)/4, $0x3eaaaaaa
DATA SSE_LCPI0_15<>+4(SB)/4, $0x3eaaaaaa
DATA SSE_LCPI0_15<>+8(SB)/4, $0x3eaaaaaa
DATA SSE_LCPI0_15<>+12(SB)/4, $0x3eaaaaaa
GLOBL SSE_LCPI0_15<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_16<>+0(SB)/4, $0xb95e8083
DATA SSE_LCPI0_16<>+4(SB)/4, $0xb95e8083
DATA SSE_LCPI0_16<>+8(SB)/4, $0xb95e8083
DATA SSE_LCPI0_16<>+12(SB)/4, $0xb95e8083
GLOBL SSE_LCPI0_16<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_17<>+0(SB)/4, $0xbf000000
DATA SSE_LCPI0_17<>+4(SB)/4, $0xbf000000
DATA SSE_LCPI0_17<>+8(SB)/4, $0xbf000000
DATA SSE_LCPI0_17<>+12(SB)/4, $0xbf000000
GLOBL SSE_LCPI0_17<>(SB), RODATA|NOPTR, $16

DATA SSE_LCPI0_18<>+0(SB)/4, $0x3f318000
DATA SSE_LCPI0_18<>+4(SB)/4, $0x3f318000
DATA SSE_LCPI0_18<>+8(SB)/4, $0x3f318000
DATA SSE_LCPI0_18<>+12(SB)/4, $0x3f318000
GLOBL SSE_LCPI0_18<>(SB), RODATA|NOPTR, $16

// func LogSSE32(x []float32, y []float32)
// Requires: SSE, SSE2
TEXT ·LogSSE32(SB), NOSPLIT, $0-48
	MOVQ     x_base+0(FP), AX
	MOVQ     y_base+24(FP), CX
	MOVUPS   (AX), X0
	XORPS    X2, X2
	MOVAPS   X0, X1
	CMPPS    X2, X1, $0x02
	MAXPS    SSE_LCPI0_0<>+0(SB), X0
	MOVAPS   X0, X2
	PSRLL    $0x17, X2
	ANDPS    SSE_LCPI0_1<>+0(SB), X0
	ORPS     SSE_LCPI0_2<>+0(SB), X0
	PADDD    SSE_LCPI0_3<>+0(SB), X2
	MOVAPS   X0, X4
	CMPPS    SSE_LCPI0_5<>+0(SB), X4, $0x01
	MOVAPS   X4, X3
	ANDPS    X0, X3
	ADDPS    SSE_LCPI0_6<>+0(SB), X0
	ADDPS    X3, X0
	MOVAPS   SSE_LCPI0_7<>+0(SB), X3
	MULPS    X0, X3
	ADDPS    SSE_LCPI0_8<>+0(SB), X3
	CVTPL2PS X2, X2
	MULPS    X0, X3
	ADDPS    SSE_LCPI0_9<>+0(SB), X3
	MOVAPS   SSE_LCPI0_4<>+0(SB), X5
	MULPS    X0, X3
	ADDPS    SSE_LCPI0_10<>+0(SB), X3
	ADDPS    X5, X2
	MULPS    X0, X3
	ADDPS    SSE_LCPI0_11<>+0(SB), X3
	ANDPS    X5, X4
	MULPS    X0, X3
	ADDPS    SSE_LCPI0_12<>+0(SB), X3
	SUBPS    X4, X2
	MOVAPS   X0, X4
	MULPS    X0, X3
	ADDPS    SSE_LCPI0_13<>+0(SB), X3
	MULPS    X0, X4
	MULPS    X0, X3
	ADDPS    SSE_LCPI0_14<>+0(SB), X3
	MULPS    X0, X3
	ADDPS    SSE_LCPI0_15<>+0(SB), X3
	MULPS    X0, X3
	MULPS    X4, X3
	MOVAPS   SSE_LCPI0_16<>+0(SB), X5
	MULPS    X2, X5
	ADDPS    X3, X5
	MULPS    SSE_LCPI0_17<>+0(SB), X4
	MULPS    SSE_LCPI0_18<>+0(SB), X2
	ADDPS    X5, X4
	ADDPS    X4, X0
	ADDPS    X2, X0
	ORPS     X1, X0
	MOVUPS   X0, (CX)
	RET