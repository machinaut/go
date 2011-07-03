// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// func Outw(port, val uint16)
TEXT	·Outw(SB),7,$0
	MOVL	4(SP), DX	// port
	MOVL	8(SP), AX // word
	OUTW
	RET

// func Inw(port uint16) uint16
TEXT	·Inw(SB),7,$0
	MOVL	4(SP), DX
	INW
	MOVL	AX, 8(SP)

// func Outb(port uint16, val uint8)
TEXT	·Outb(SB),7,$0
	MOVL	4(SP), DX
	MOVL	8(SP), AX
	OUTB
	RET

// func Inb(port uint16) uint8
TEXT	·Inb(SB),7,$0
	MOVL	4(SP), DX
	INB
	MOVL	AX, 8(SP)
	RET

