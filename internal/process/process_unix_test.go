//go:build !windows

package process_test

import (
	"syscall"
	"testing"

	"orbit/internal/process"
)

func TestDetachProcessUnix(t *testing.T) {
	var sysAttr syscall.SysProcAttr
	process.DetachProcess(&sysAttr)

	if !sysAttr.Setsid {
		t.Errorf("expected Setsid to be true on Unix platform")
	}
}
