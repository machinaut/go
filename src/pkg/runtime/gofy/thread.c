// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "runtime.h"
#include "defs.h"
#include "os.h"
#include "stack.h"

void
runtime·lock(Lock *l)
{
	if(m->locks < 0)
		runtime·throw("lock count");
	m->locks++;
	if(l->key != 0)
		runtime·throw("deadlock");
	l->key = 1;
}

void
runtime·unlock(Lock *l)
{
	m->locks--;
	if(m->locks < 0)
		runtime·throw("lock count");
	if(l->key != 1)
		runtime·throw("unlock of unlocked lock");
	l->key = 0;
}

void
runtime·destroylock(Lock*)
{
}


// One-time notifications.
//
// Since the lock/unlock implementation already
// takes care of sleeping in the kernel, we just reuse it.
// (But it's a weird use, so it gets its own interface.)
//
// We use a lock to represent the event:
// unlocked == event has happened.
// Thus the lock starts out locked, and to wait for the
// event you try to lock the lock.  To signal the event,
// you unlock the lock.

void
runtime·noteclear(Note *n)
{
	n->lock.key = 0;	// memset(n, 0, sizeof *n)
}

void
runtime·notewakeup(Note *n)
{
  n->lock.key = 1;
}

void
runtime·notesleep(Note *n)
{
   if(n->lock.key != 1)
   runtime·throw("notesleep");
}

void
runtime·newosproc(M *m, G *g, void *stk, void (*fn)(void))
{
   USED(m, g, stk, fn);
   runtime·throw("newosproc");
}

void
runtime·osinit(void)
{
}

void
runtime·goenvs(void)
{
}

// Called to initialize a new m (including the bootstrap m).
void
runtime·minit(void)
{
}
