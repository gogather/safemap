package safemap

import "sync"

// SafeMap safe map struct
type SafeMap struct {
	sync.RWMutex
	m map[string]interface{}
}

// New new a SafeMap
func New() *SafeMap {
	return &SafeMap{
		m: make(map[string]interface{}),
	}
}

// Put put element into safemap
func (sm *SafeMap) Put(key string, value interface{}) {
	sm.Lock()
	sm.m[key] = value
	sm.Unlock()
}

// Remove remove element from safemap
func (sm *SafeMap) Remove(key string) {
	sm.Lock()
	delete(sm.m, key)
	sm.Unlock()
}

// Get get element from safemap
func (sm *SafeMap) Get(key string) (interface{}, bool) {
	defer func() {
		sm.RUnlock()
	}()
	sm.RLock()
	v, ok := sm.m[key]
	return v, ok
}
