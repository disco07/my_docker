package namespaces

import (
	"syscall"
)

func CreateNewMountNamespace() error {
	_, _, errno := syscall.RawSyscall(syscall.SYS_CLONE, syscall.CLONE_NEWNS, 0, 0)
	if errno != 0 {
		return errno
	}
	// Ici, configurez les systèmes de fichiers spécifiques au conteneur
	return nil
}
