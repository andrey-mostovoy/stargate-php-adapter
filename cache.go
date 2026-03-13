package main

import "sync"

type PreparedCache struct {
	mu sync.RWMutex
	data map[string]interface{}
}

func NewCache() *PreparedCache {
	return &PreparedCache{
		data: make(map[string]interface{}),
	}
}

func (c *PreparedCache) Get(q string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[q]
	return v, ok
}

func (c *PreparedCache) Set(q string, stmt interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[q] = stmt
}