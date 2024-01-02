package namespace

import "syscall"

func NewPIDNameSpace() (err error) {
	r1, r2, errno := syscall.RawSyscall(syscall.SYS_UNSHARE, syscall.CLONE_NEWPID, 0, 0)
	if errno != 0 {
		err = errno
		return
	}

	if r1 != 0 || r2 != 0 {
		err = syscall.EINVAL
		return
	}

	return
}
