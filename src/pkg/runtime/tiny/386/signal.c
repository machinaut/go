// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file. 

#include "runtime.h"
#include "tiny/386/ureg.h"
#include "tiny/386/signals.h"

void
runtime·gettime(int64*, int32*) 
{
}

String
runtime·signame(int32)
{
	return runtime·emptystring;
}

typedef struct Segdesc {
	uint32 d0;
	uint32 d1;
} Segdesc;

Segdesc ilt[256];

#define SELGDT (0<<3)  /* selector is in gdt */
#define SEGP (1<<15)   /* segment present */
#define SEGPL(x) ((x)<<13) /* priority level */

#define  KESEG 2 /* kernel executable */
#define SEGTG  (0x0F<<8) /* task gate */
#define SELECTOR(i, t, p)  (((i)<<3) | (t) | (p))
#define KESEL  SELECTOR(KESEG, SELGDT, 0)

void
runtime·sethvec(uint32 v, void (*r)(void), uint32 type, uint32 pri)
{
	ilt[v].d0 = ((uint32)r)&0xFFFF|(KESEL<<16);
	ilt[v].d1 = ((uint32)r)&0xFFFF0000|SEGP|SEGPL(pri)|type;
}

void runtime·putidt(Segdesc *p, uint16 len);

void
runtime·trapinit(void)
{
	int32 i;

	/* set all to bad, then fix the ones we want */
	for(i = 0; i < 256; i++)
		runtime·sethvec(i, intrbad, SEGTG, 0);

	runtime·sethvec(6, intr6, SEGTG, 0);
}

void
runtime·dumpregs(Ureg *r)
{
	runtime·printf("eax     %x\t", r->ax);
	runtime·printf("ebx     %x\n", r->bx);
	runtime·printf("ecx     %x\t", r->cx);
	runtime·printf("edx     %x\n", r->dx);
	runtime·printf("edi     %x\t", r->di);
	runtime·printf("esi     %x\n", r->si);
	runtime·printf("ebp     %x\t", r->bp);
	runtime·printf("esp     %x\n", r->nsp);
	runtime·printf("eip     %x\t", r->pc);
	runtime·printf("eflags  %x\n", r->flags);
	runtime·printf("cs      %x\t", r->cs);
	runtime·printf("fs      %x\n", r->fs);
	runtime·printf("gs      %x\n", r->gs);
}

void runtime·arch_exit();

void
runtime·trap(Ureg *ur)
{
	runtime·printf("Trap %x:\n", ur->trap);
	runtime·dumpregs(ur);

	runtime·arch_exit();
}

