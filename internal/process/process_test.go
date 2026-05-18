package process_test

import (
	"os/exec"
	"runtime"
	"testing"
	"time"

	"orbit/internal/process"
)

func TestStopProcess(t *testing.T) {
	t.Run("StopProcess", func(t *testing.T) {
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			// Ping local loopback 5 times (takes ~4 seconds if not interrupted)
			cmd = exec.Command("ping", "127.0.0.1", "-n", "5")
		} else {
			// Sleep for 5 seconds
			cmd = exec.Command("sleep", "5")
		}

		if err := cmd.Start(); err != nil {
			t.Fatalf("failed to start test process: %v", err)
		}

		// Wait briefly to ensure it is running
		time.Sleep(100 * time.Millisecond)

		// Stop the process
		startTime := time.Now()
		if err := process.StopProcess(cmd.Process); err != nil {
			t.Fatalf("failed to stop process: %v", err)
		}

		// Wait for process to exit
		err := cmd.Wait()
		duration := time.Since(startTime)

		if duration > 2*time.Second {
			t.Errorf("process was not stopped quickly, took %v", duration)
		}

		t.Logf("Process terminated successfully in %v. Exit error (expected): %v", duration, err)
	})
}
