package process

import (
	"orbit/internal/models"
	"os/exec"
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
	cmd := exec.Command("orbit", "runtime", "execute", "--name", name)
	cmd.Dir = workspacePath

	if err := cmd.Start(); err != nil {
		return err
	}

	meta := &models.RuntimeMetaData{
		PID:       cmd.Process.Pid,
		Name:      name,
		Status:    "Running",
		StartedAt: time.Now(),
	}
	m.Updater.Update(meta)

	go func() {
		err := cmd.Wait()
		status := "Stopped"
		if err != nil {
			status = "Crashed"
		}
		meta.Status = status
		m.Updater.Update(meta)
	}()

	return nil
}
