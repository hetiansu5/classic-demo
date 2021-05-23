#include "textflag.h"

// 分别获取当前伪SP、硬件SP、伪FP地址
// 思路：使用LEA把各个寄存器的地址传出来即可
// 当栈帧不为0的时候，伪SP = 硬件SP + 16（栈帧大小）  伪FP = 伪SP + 16 (return addr + caller BP的大小)
// 硬件SP  0xc000110e88  函数栈真实栈顶地址（如果当前栈帧中没有callee，则是最后一个局部变量的开始位置）
// 伪的SP  0xc000110e98  当前栈帧第一个局部变量的结束位置（等价于caller BP的开始位置， BP位置存储的为：进入函数前的栈顶基址）
// 伪的FP  0xc000110ea8  指向caller调用callee时传递的第一个参数的开始位置
//    BP  0xc00009af78  caller BP存储的地址
// 开始位置和结束位置，内存从低位到高位，低位为开始位置，高位为结束位置。比如$a占用8字节，地址0x10~0x17，开始位置为0x10，结束位置为0x18
TEXT ·SpFp(SB),NOSPLIT,$16-48
    LEAQ (SP), AX
	LEAQ c+0(SP), BX
	LEAQ b+0(FP), CX
	MOVQ AX, r0+8(FP)
	MOVQ BX, r1+16(FP)
	MOVQ CX, r2+24(FP)
	MOVQ 32(SP), AX
	MOVQ AX, r3+32(FP)
	MOVQ a+16(SP), AX
	MOVQ AX, r4+40(FP)
	RET

// 栈帧为0的时候， 伪SP = 硬件SP
// 硬件SP  0xc00005ee30  caller栈帧return addr的开始位置
// 伪的SP  0xc00005ee30  caller栈帧return addr的开始位置
// 伪的FP  0xc00005ee38  指向caller调用callee时传递的第一个参数的开始位置
TEXT ·SpZero(SB),NOSPLIT,$0-48
    LEAQ (SP), AX
	LEAQ c+0(SP), BX
	LEAQ b+0(FP), CX
	MOVQ AX, r0+8(FP)
	MOVQ BX, r1+16(FP)
	MOVQ CX, r2+24(FP)
	MOVQ 8(SP), AX
	MOVQ AX, r3+32(FP)
	MOVQ a+8(SP), AX
	MOVQ AX, r4+40(FP)
	RET
