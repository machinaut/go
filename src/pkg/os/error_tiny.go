// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import syscall "syscall"

// An Error can represent any printable error condition.
type Error interface {
	String() string
}

// A helper type that can be embedded or wrapped to simplify satisfying
// Error.
type ErrorString string

func (e ErrorString) String() string  { return string(e) }
func (e ErrorString) Temporary() bool { return false }
func (e ErrorString) Timeout() bool   { return false }

// Note: If the name of the function NewError changes,
// pkg/go/doc/doc.go should be adjusted since it hardwires
// this name in a heuristic.

// NewError converts s to an ErrorString, which satisfies the Error interface.
func NewError(s string) Error { return ErrorString(s) }

// Errno is the Unix error number.  Names such as EINVAL are simple
// wrappers to convert the error number into an Error.
type Errno int64

func (e Errno) String() string { return syscall.Errstr(int(e)) }

func (e Errno) Temporary() bool {
	return false
}

func (e Errno) Timeout() bool {
	return false
}

// these are needed by io and strconv

var (
	EINVAL Error = Errno(0x16)
	EPIPE  Error = Errno(0x20)
	ERANGE Error = Errno(0x22)
)

// PathError records an error and the operation and file path that caused it.
type PathError struct {
	Op    string
	Path  string
	Error Error
}

func (e *PathError) String() string { return e.Op + " " + e.Path + ": " + e.Error.String() }

// SyscallError records an error from a specific system call.
type SyscallError struct {
	Syscall string
	Errno   Errno
}

func (e *SyscallError) String() string { return e.Syscall + ": " + e.Errno.String() }

// Note: If the name of the function NewSyscallError changes,
// pkg/go/doc/doc.go should be adjusted since it hardwires
// this name in a heuristic.

// NewSyscallError returns, as an Error, a new SyscallError
// with the given system call name and error number.
// As a convenience, if errno is 0, NewSyscallError returns nil.
func NewSyscallError(syscall string, errno int) Error {
	if errno == 0 {
		return nil
	}
	return &SyscallError{syscall, Errno(errno)}
}
