// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

import (
	"os"
)

// OpenReader will open the Zip file specified by name and return a ReaderCloser.
func OpenReader(name string) (*ReadCloser, os.Error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}
	r := new(ReadCloser)
	if err := r.init(f, fi.Size); err != nil {
		f.Close()
		return nil, err
	}
	return r, nil
}
