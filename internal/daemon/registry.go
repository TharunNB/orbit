package daemon

import (
	"sync"
	"time"
)

type RuntimeMetaData struct {
	PID       int       `json:"pid"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	StartedAt time.Time `json:"started_at"`
}

type Registry struct {
	mu     sync.RWMutex
	Agents map[string]*RuntimeMetaData
}

func NewRegistry() *Registry {
	return &Registry{
		Agents: make(map[string]*RuntimeMetaData),
	}
}

func (r *Registry) Update(meta *RuntimeMetaData) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Agents[meta.Name] = meta
}

func (r *Registry) GetAll() []*RuntimeMetaData {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]*RuntimeMetaData, 0, len(r.Agents))
	for _, v := range r.Agents {
		list = append(list, v)
	}
	return list
}
