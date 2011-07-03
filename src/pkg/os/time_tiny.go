// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import "syscall"

func Time() (sec int64, nsec int64, err Error) {
	var tv syscall.Timeval

	// same as the frozen time on the playground
	tv.Sec = 1257894000
	tv.Usec = 0

	return int64(tv.Sec), int64(tv.Usec) * 1000, nil
}
