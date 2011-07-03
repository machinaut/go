// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// For tiny, just return the same values as "file not found"
// for Unix.

package time

func lookupTimezone(sec int64) (zone string, offset int) {
	return "UTC", 0
}

func lookupByName(name string) (off int, found bool) {
	return 0, false
}
