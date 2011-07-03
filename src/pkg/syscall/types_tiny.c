// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <sys/types.h>
typedef struct timespec $Timespec;
typedef struct timeval $Timeval;

#include <sys/user.h>
typedef struct user_regs_struct $PtraceRegs;

// for os/stat_tiny.go (required by os/file.go)
#include <sys/stat.h>
typedef struct stat $Stat_t;

// for os/exec.go
#include <sys/resource.h>
typedef struct rusage $Rusage;
