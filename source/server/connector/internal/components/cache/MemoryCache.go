package cache

import "sync"

type MemoryCache struct {
	cache map[string]interface{}
	mux   *sync.RWMutex
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		cache: make(map[string]interface{}),
		mux:   &sync.RWMutex{},
	}
}

func (m *MemoryCache) Set(key string, value interface{}) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.cache[key] = value
}

func (m *MemoryCache) Get(key string) (interface{}, bool) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	v, ok := m.cache[key]
	return v, ok
}
func (m *MemoryCache) Remove(key string) {
	m.mux.Lock()
	defer m.mux.Unlock()
	delete(m.cache, key)
}
