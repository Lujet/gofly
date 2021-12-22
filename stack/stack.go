package main

/* go 的函数调用栈 */
// go 底层原理分析

func mul(a, b int) int {
	return a * b
}

func main() {
	mul(3, 4)
}

// "".mul STEXT nosplit size=29 args=0x18 locals=0x0
// 	0x0000 00000 (stack.go:3)	TEXT	"".mul(SB), NOSPLIT|ABIInternal, $0-24
// 	0x0000 00000 (stack.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 	0x0000 00000 (stack.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 	0x0000 00000 (stack.go:3)	MOVQ	$0, "".~r2+24(SP)
// 	0x0009 00009 (stack.go:4)	MOVQ	"".b+16(SP), AX
// 	0x000e 00014 (stack.go:4)	MOVQ	"".a+8(SP), CX
// 	0x0013 00019 (stack.go:4)	IMULQ	AX, CX
// 	0x0017 00023 (stack.go:4)	MOVQ	CX, "".~r2+24(SP)
// 	0x001c 00028 (stack.go:4)	RET
// 	0x0000 48 c7 44 24 18 00 00 00 00 48 8b 44 24 10 48 8b  H.D$.....H.D$.H.
// 	0x0010 4c 24 08 48 0f af c8 48 89 4c 24 18 c3           L$.H...H.L$..
// "".main STEXT size=71 args=0x0 locals=0x20
// 	0x0000 00000 (stack.go:7)	TEXT	"".main(SB), ABIInternal, $32-0
// 	0x0000 00000 (stack.go:7)	MOVQ	(TLS), CX
// 	0x0009 00009 (stack.go:7)	CMPQ	SP, 16(CX)
// 	0x000d 00013 (stack.go:7)	PCDATA	$0, $-2
// 	0x000d 00013 (stack.go:7)	JLS	64
// 	0x000f 00015 (stack.go:7)	PCDATA	$0, $-1
// 	0x000f 00015 (stack.go:7)	SUBQ	$32, SP
// 	0x0013 00019 (stack.go:7)	MOVQ	BP, 24(SP)
// 	0x0018 00024 (stack.go:7)	LEAQ	24(SP), BP
// 	0x001d 00029 (stack.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 	0x001d 00029 (stack.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 	0x001d 00029 (stack.go:8)	MOVQ	$3, (SP)
// 	0x0025 00037 (stack.go:8)	MOVQ	$4, 8(SP)
// 	0x002e 00046 (stack.go:8)	PCDATA	$1, $0
// 	0x002e 00046 (stack.go:8)	CALL	"".mul(SB)
// 	0x0033 00051 (stack.go:9)	MOVQ	24(SP), BP
// 	0x0038 00056 (stack.go:9)	ADDQ	$32, SP
// 	0x003c 00060 (stack.go:9)	RET
// 	0x003d 00061 (stack.go:9)	NOP
// 	0x003d 00061 (stack.go:7)	PCDATA	$1, $-1
// 	0x003d 00061 (stack.go:7)	PCDATA	$0, $-2
// 	0x003d 00061 (stack.go:7)	NOP
// 	0x0040 00064 (stack.go:7)	CALL	runtime.morestack_noctxt(SB)
// 	0x0045 00069 (stack.go:7)	PCDATA	$0, $-1
// 	0x0045 00069 (stack.go:7)	JMP	0
// 	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 31 48  dH..%....H;a.v1H
// 	0x0010 83 ec 20 48 89 6c 24 18 48 8d 6c 24 18 48 c7 04  .. H.l$.H.l$.H..
// 	0x0020 24 03 00 00 00 48 c7 44 24 08 04 00 00 00 e8 00  $....H.D$.......
// 	0x0030 00 00 00 48 8b 6c 24 18 48 83 c4 20 c3 0f 1f 00  ...H.l$.H.. ....
// 	0x0040 e8 00 00 00 00 eb b9                             .......
// 	rel 5+4 t=17 TLS+0
// 	rel 47+4 t=8 "".mul+0
// 	rel 65+4 t=8 runtime.morestack_noctxt+0
// go.cuinfo.packagename. SDWARFINFO dupok size=0
// 	0x0000 6d 61 69 6e                                      main
// ""..inittask SNOPTRDATA size=24
// 	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
// 	0x0010 00 00 00 00 00 00 00 00                          ........
// gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
// 	0x0000 01 00 00 00 00 00 00 00                          ........
