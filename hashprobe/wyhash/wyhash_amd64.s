//go:build !purego

#include "textflag.h"

#define m1 0xa0761d6478bd642f
#define m2 0xe7037ed1a0b428db
#define m3 0x8ebc6af09c88c6e3
#define m4 0x589965cc75374cc3
#define m5 0x1d8e4e27c47d124f

// func MultiHash32(hashes []uintptr, values []uint32, seed uintptr)
TEXT ·MultiHash32(SB), NOSPLIT, $0-56
    MOVQ hashes_base+0(FP), R12
    MOVQ values_base+24(FP), DI
    MOVQ values_len+32(FP), CX
    MOVQ seed+48(FP), R11

    MOVQ $m1, R8
    MOVQ $m2, R9
    MOVQ $m5^4, R10

    XORQ SI, SI
    JMP test
loop:
    MOVQ R11, R13
hash:
    MOVL (DI)(SI*4), AX

    MOVQ R8, BX
    XORQ R13, BX

    XORQ AX, BX
    XORQ R9, AX

    MULQ BX
    XORQ DX, AX

    MULQ R10
    XORQ DX, AX

    CMPQ AX, $0
    JE next

    MOVQ AX, (R12)(SI*8)
    INCQ SI
test:
    CMPQ SI, CX
    JNE loop
    RET
next:
    INCQ R13
    JMP hash

// func MultiHash64(hashes []uintptr, values []uint64, seed uintptr)
TEXT ·MultiHash64(SB), NOSPLIT, $0-56
    MOVQ hashes_base+0(FP), R12
    MOVQ values_base+24(FP), DI
    MOVQ values_len+32(FP), CX
    MOVQ seed+48(FP), R11

    MOVQ $m1, R8
    MOVQ $m2, R9
    MOVQ $m5^8, R10

    XORQ SI, SI
    JMP test
loop:
    MOVQ R11, R13
hash:
    MOVQ (DI)(SI*8), AX

    MOVQ R8, BX
    XORQ R13, BX

    XORQ AX, BX
    XORQ R9, AX

    MULQ BX
    XORQ DX, AX

    MULQ R10
    XORQ DX, AX

    CMPQ AX, $0
    JE next

    MOVQ AX, (R12)(SI*8)
    INCQ SI
test:
    CMPQ SI, CX
    JNE loop
    RET
next:
    INCQ R13
    JMP hash

// func MultiHash128(hashes []uintptr, values [][16]byte, seed uintptr)
TEXT ·MultiHash128(SB), NOSPLIT, $0-56
    MOVQ hashes_base+0(FP), R12
    MOVQ values_base+24(FP), DI
    MOVQ values_len+32(FP), CX
    MOVQ seed+48(FP), R11

    MOVQ $m1, R8
    MOVQ $m2, R9
    MOVQ $m5^16, R10

    XORQ SI, SI
    JMP test
loop:
    MOVQ R11, R13
hash:
    MOVQ 0(DI), AX
    MOVQ 8(DI), DX

    MOVQ R8, BX
    XORQ R13, BX

    XORQ DX, BX
    XORQ R9, AX

    MULQ BX
    XORQ DX, AX

    MULQ R10
    XORQ DX, AX

    MOVQ AX, (R12)(SI*8)
    ADDQ $16, DI
    INCQ SI
test:
    CMPQ SI, CX
    JNE loop
    RET
next:
    INCQ R13
    JMP hash
