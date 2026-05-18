//go:build windows

package process_test

import (
	"syscall"
	"testing"

	"orbit/internal/process"
)

func TestDetachProcessWindows(t *testing.T) {
	var sysAttr syscall.SysProcAttr
	process.DetachProcess(&sysAttr)

	if sysAttr.CreationFlags != syscall.CREATE_NEW_PROCESS_GROUP {
		t.Errorf("expected CreationFlags to have CREATE_NEW_PROCESS_GROUP, got %d", sysAttr.CreationFlags)
	}
}
