package test

import (
	. "orbit/internal/daemon"
	"orbit/internal/models"
	"testing"
	"time"
)

func TestRegistry(t *testing.T) {
	t.Run("UpdateAndGet", func(t *testing.T) {
		reg := NewRegistry()
		meta := &models.RuntimeMetaData{
			PID:       1234,
			Name:      "test-agent",
			Status:    "Running",
			StartedAt: time.Now(),
		}

		reg.Update(meta)

		retrieved, ok := reg.Get("test-agent")
		if !ok {
			t.Errorf("expected to find agent 'test-agent'")
		}
		if retrieved.PID != 1234 {
			t.Errorf("expected PID 1234, got %d", retrieved.PID)
		}
	})

	t.Run("GetAll", func(t *testing.T) {
		reg := NewRegistry()
		meta1 := &models.RuntimeMetaData{Name: "agent-1"}
		meta2 := &models.RuntimeMetaData{Name: "agent-2"}

		reg.Update(meta1)
		reg.Update(meta2)

		agents := reg.GetAll()
		if len(agents) != 2 {
			t.Errorf("expected 2 agents, got %d", len(agents))
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		reg := NewRegistry()
		_, ok := reg.Get("non-existent")
		if ok {
			t.Errorf("expected ok=false for non-existent agent")
		}
	})
}
