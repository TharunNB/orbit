//go:build !windows

package process

import "syscall"

func DetachProcess(attr *syscall.SysProcAttr) {
	attr.Setsid = true
}
