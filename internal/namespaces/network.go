package namespaces

import (
	"syscall"
)

func CreateNewNetNamespace() error {
	_, _, errno := syscall.RawSyscall(syscall.SYS_CLONE, syscall.CLONE_NEWNET, 0, 0)
	if errno != 0 {
		return errno
	}
	// Ici, vous devez configurer le réseau dans le nouveau namespace
	return nil
}
