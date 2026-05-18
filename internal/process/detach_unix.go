//go:build !windows

package process

import (
	"os"
	"syscall"
)

func DetachProcess(attr *syscall.SysProcAttr) {
	attr.Setsid = true
}

func StopProcess(proc *os.Process) error {
	return proc.Signal(syscall.SIGTERM)
}
