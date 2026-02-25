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
	entry map[string]cacheEntry
	mutex sync.Mutex
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mutex.Lock()
		for key, val := range c.entry {
			if val.createdAt.Before(time.Now().Add(-interval)) {
				delete(c.entry, key)
			}
		}
		c.mutex.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{entry: make(map[string]cacheEntry)}
	cPtr := &c

	go cPtr.reapLoop(interval)

	return cPtr
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entry[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.entry[key]
	return val.val, ok
}
