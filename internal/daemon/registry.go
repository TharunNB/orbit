package daemon

import (
	"orbit/internal/models"
	"sync"
)

type Registry struct {
	mu     sync.RWMutex
	Agents map[string]*models.RuntimeMetaData
}

func NewRegistry() *Registry {
	return &Registry{
		Agents: make(map[string]*models.RuntimeMetaData),
	}
}

func (r *Registry) Update(meta *models.RuntimeMetaData) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Agents[meta.Name] = meta
}

func (r *Registry) Get(name string) (*models.RuntimeMetaData, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	meta, ok := r.Agents[name]
	return meta, ok
}

func (r *Registry) GetAll() []*models.RuntimeMetaData {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]*models.RuntimeMetaData, 0, len(r.Agents))
	for _, v := range r.Agents {
		list = append(list, v)
	}
	return list
}
