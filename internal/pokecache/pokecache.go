package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data     map[string]cacheEntry
	mux      sync.RWMutex
	interval time.Duration
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// Read lock
	c.mux.RLock()
	defer c.mux.RUnlock()

	val, succ := c.data[key]
	if !succ {
		return nil, false
	}

	return val.val, true
}

func (c *Cache) Add(key string, val []byte) {
	// Write lock
	c.mux.Lock()
	defer c.mux.Unlock()

	c.data[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		data:     make(map[string]cacheEntry),
		mux:      sync.RWMutex{},
		interval: interval,
	}

	go c.reapLoop()

	return &c
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.data {
			if c.interval <= time.Since(entry.createdAt) {
				delete(c.data, key)
			}
		}
		c.mux.Unlock()
	}
}
