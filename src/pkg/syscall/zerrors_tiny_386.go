package syscall

const (
	// for time/sleep.go
	EINTR = 0x4
)

var errors = [...]string{
	4: "interrupted system call",
}
