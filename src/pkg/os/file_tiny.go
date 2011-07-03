// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// These are needed for io/*.go:

type eofError int

func (eofError) String() string { return "EOF" }

// EOF is the Error returned by Read when no more input is available.
// Functions should return EOF only to signal a graceful end of input.
// If the EOF occurs unexpectedly in a structured data stream,
// the appropriate error is either io.ErrUnexpectedEOF or some other error
// giving more detail.
var EOF Error = eofError(0)

// Needed for fmt:

// File represents an open file descriptor.
type File struct {
	fd      int
	name    string
	dirinfo *dirInfo
	nepipe  int
}

type dirInfo struct {
	buf  []byte // buffer for directory I/O
	nbuf int    // length of buf; return value from Getdirentries
	bufp int    // location of next record in buf.
}

var (
	Stdin  = &File{0, "stdin", nil, 0}
	Stdout = &File{1, "stdout", nil, 0}
	Stderr = &File{2, "stderr", nil, 0}
)

// implement io.Writer
func (f *File) Write(p []byte) (n int, err Error) {
	print(p)
	return len(p), nil
}

// implement io.Reader
func (f *File) Read(p []byte) (n int, err Error) {
	return 0, nil
}
