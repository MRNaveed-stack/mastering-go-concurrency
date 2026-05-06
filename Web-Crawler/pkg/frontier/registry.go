package frontier

import "sync"

type Registry struct {
	mu      sync.RWMutex
	visited map[string]bool
}

func NewRegistry() *Registry {
	return &Registry{
		visited: make(map[string]bool),
	}
}

func (r *Registry) Visit(url string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.visited[url] {
		return false
	}
	r.visited[url] = true
	return true
}

func (r *Registry) Count() int {
	r.mu.RLock()
	defer r.mu.Unlock()
	return len(r.visited)
}
