// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

TEXT _rt0_386_tiny(SB), 7, $0
	// Disable interrupts.
	CLI


	// setup floating point
	// from 8a/l.s, the rt0 for Plan9
	MOVL  CR0,AX
	ANDL  $~0x4,AX  /* EM=0 */
	MOVL  AX,CR0
	FINIT
	WAIT
	PUSHW $0x0330 /* ignore underflow/precision, signal others */
	FLDCW 0(SP)
	POPW  AX
	WAIT
	
	// Establish stack.
	MOVL	$0x10000, AX
	MOVL	AX, SP
	
	// Set up memory hardware.
	CALL	runtime·msetup(SB)

	// _rt0_386 expects to find argc, argv, envv on stack.
	// Set up argv=["kernel"] and envv=[].
	SUBL	$64, SP
	MOVL	$1, 0(SP)
	MOVL	$runtime·kernel(SB), 4(SP)
	MOVL	$0, 8(SP)
	MOVL	$0, 12(SP)
	JMP	_rt0_386(SB)

DATA runtime·kernel(SB)/7, $"kernel\z"
GLOBL runtime·kernel(SB), $7
