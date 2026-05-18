//go:build windows

package process

import (
	"os"
	"syscall"
)

func DetachProcess(attr *syscall.SysProcAttr) {
	attr.CreationFlags = syscall.CREATE_NEW_PROCESS_GROUP
}

func StopProcess(proc *os.Process) error {
	return proc.Kill()
}
