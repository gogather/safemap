package safemap

import (
	"encoding/json"
	"fmt"
	"sync"
)

// SafeMap safe map struct
type SafeMap struct {
	sync.RWMutex
	m    map[string]interface{}
	keys []string
}

// New new a SafeMap
func New() *SafeMap {
	return &SafeMap{
		m:    make(map[string]interface{}),
		keys: []string{},
	}
}

// Put put element into safemap
func (sm *SafeMap) Put(key string, value interface{}) {
	sm.Lock()
	sm.m[key] = value
	sm.keys = append(sm.keys, key)
	sm.Unlock()
}

// Keys sorted keys
func (sm *SafeMap) Keys() []string {
	return sm.keys
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

func (sm *SafeMap) String() string {
	defer func() {
		sm.RUnlock()
	}()
	sm.RLock()
	return fmt.Sprintf("%v", sm.m)
}

// JSON convert map to json string
func (sm *SafeMap) JSON() string {
	defer func() {
		sm.RUnlock()
	}()
	sm.RLock()
	c, _ := json.Marshal(sm.m)
	return string(c)
}

// GetMap get original map
func (sm *SafeMap) GetMap() map[string]interface{} {
	return sm.m
}

// Clear clear the map
func (sm *SafeMap) Clear() {
	sm.m = make(map[string]interface{})
	sm.keys = make([]string, 0)
}
