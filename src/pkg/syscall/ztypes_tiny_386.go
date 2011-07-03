// godefs -gsyscall types_tiny.c

// MACHINE GENERATED - DO NOT EDIT.

package syscall

// Constants

// Types

type Timespec struct {
	Sec  int32
	Nsec int32
}

type Timeval struct {
	Sec  int32
	Usec int32
}

type PtraceRegs struct {
	Ebx      int32
	Ecx      int32
	Edx      int32
	Esi      int32
	Edi      int32
	Ebp      int32
	Eax      int32
	Xds      int32
	Xes      int32
	Xfs      int32
	Xgs      int32
	Orig_eax int32
	Eip      int32
	Xcs      int32
	Eflags   int32
	Esp      int32
	Xss      int32
}

type Stat_t struct {
	Dev        uint64
	X__pad1    uint16
	Pad0       [2]byte
	Ino        uint32
	Mode       uint32
	Nlink      uint32
	Uid        uint32
	Gid        uint32
	Rdev       uint64
	X__pad2    uint16
	Pad1       [2]byte
	Size       int32
	Blksize    int32
	Blocks     int32
	Atim       Timespec
	Mtim       Timespec
	Ctim       Timespec
	X__unused4 uint32
	X__unused5 uint32
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int32
	Ixrss    int32
	Idrss    int32
	Isrss    int32
	Minflt   int32
	Majflt   int32
	Nswap    int32
	Inblock  int32
	Oublock  int32
	Msgsnd   int32
	Msgrcv   int32
	Nsignals int32
	Nvcsw    int32
	Nivcsw   int32
}
