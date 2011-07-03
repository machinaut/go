# Copyright 2010 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# Include this after Make.cmd to make a bootable disk.
# Only actually useful for GOOS=tiny, of course.

$(TARG).asm: _go_.8
	$(LD) -d -a _go_.8 > $(TARG).asm

disk: $(TARG) $(TARG).asm
	dd if=/dev/zero of=disk count=10000
	cat $(GOROOT)/src/pkg/runtime/$(GOOS)/bootblock $(TARG) | \
		dd of=disk conv=notrunc

run: disk
	bochs

# do nothing on install
install:

CLEANFILES+=$(TARG).asm disk
