package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	Value     []byte
	createdAt time.Time
}

type Cache struct {
	data map[string]CacheEntry
	interval time.Duration
	mu   sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	currentCache := &Cache{
		data: make(map[string]CacheEntry),
		interval: time.Duration(interval),
	}
	go currentCache.readLoop()
	return  currentCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newCacheEntry := CacheEntry{}
	newCacheEntry.Value = val
	newCacheEntry.createdAt = time.Now()
	c.data[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.data[key]
	return entry.Value, ok
}

func (c *Cache) readLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.data {
			if time.Now().After(entry.createdAt.Add(c.interval)) {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}
