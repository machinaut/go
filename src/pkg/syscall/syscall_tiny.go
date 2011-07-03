// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Tiny system calls (i.e. none)

// This file is compiled as ordinary Go code,
// but it is also input to mksyscall,
// which parses the //sys lines and generates system call stubs.
// Note that sometimes we use a lowercase //sys name and
// wrap it in our own nicer implementation.

package syscall

const OS = "tiny"

// prototypes for things in assembly
func Outw(port, val uint16)
func Inw(port uint16) uint16
func Outb(port uint16, val uint8)
func Inb(port uint16) uint8

const ImplementsGetwd = false

func Sleep(nsec int64) (errno int) {
	//tv := NsecToTimeval(nsec)
	//_, err := Select(0, nil, nil, nil, &tv)
	//return err
	return 0
}

type WaitStatus uint32

// Wait status is 7 bits at bottom, either 0 (exited),
// 0x7F (stopped), or a signal number that caused an exit.
// The 0x80 bit is whether there was a core dump.
// An extra number (exit code, signal causing a stop)
// is in the high bits.  At least that's the idea.
// There are various irregularities.  For example, the
// "continued" status is 0xFFFF, distinguishing itself
// from stopped via the core dump bit.

const (
	mask    = 0x7F
	core    = 0x80
	exited  = 0x00
	stopped = 0x7F
	shift   = 8
)

// needed by os/error.go
func Errstr(errno int) string {
	if errno < 0 || errno >= int(len(errors)) {
		return "error " + str(errno)
	}
	return errors[errno]
}
