package safemap

import "sync"
import "fmt"
import "github.com/gogather/com"

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

func (sm *SafeMap) String() string {
	return fmt.Sprintf("%v", sm.m)
}

// JSON convert map to json string
func (sm *SafeMap) JSON() (json string) {
	json, _ = com.JsonEncode(sm.m)
	return
}

// GetMap get original map
func (sm *SafeMap) GetMap() map[string]interface{} {
	return sm.m
}
