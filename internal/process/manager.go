package process

import (
	"orbit/internal/models"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type StatusUpdater interface {
	Update(meta *models.RuntimeMetaData)
}

type Manager struct {
	Updater StatusUpdater
}

func NewManager(updater StatusUpdater) *Manager {
	return &Manager{Updater: updater}
}

func (m *Manager) SpawnAgent(name string, workspacePath string) error {
	execPath, err := os.Executable()
	if err != nil {
		execPath = "orbit"
	}
	cmd := exec.Command(execPath, "runtime", "execute", "--name", name)
	cmd.Dir = workspacePath

	// Ensure the logs subdirectory exists
	logDir := filepath.Join(workspacePath, "logs")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	logFilePath := filepath.Join(logDir, "agent.log")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	cmd.Stdout = logFile
	cmd.Stderr = logFile

	if err := cmd.Start(); err != nil {
		logFile.Close()
		return err
	}
	// Close parent's write handle so we don't leak open file descriptors/handles
	logFile.Close()

	meta := &models.RuntimeMetaData{
		PID:       cmd.Process.Pid,
		Name:      name,
		Status:    models.StatusRunning,
		StartedAt: time.Now(),
	}
	m.Updater.Update(meta)

	go func() {
		err := cmd.Wait()
		status := models.StatusStopped
		if err != nil {
			status = models.StatusFailed
		}
		meta.Status = status
		m.Updater.Update(meta)
	}()

	return nil
}

func (m *Manager) StopAgent(pid int) error {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return StopProcess(proc)
}
