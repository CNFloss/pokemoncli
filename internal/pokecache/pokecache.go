package pokecache

import (
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	cache map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if _, ok := c.cache[key]; !ok {
		return nil, false
	}
	return c.cache[key].val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	t := time.Now().UTC().Add(-interval)
	for i, entry := range c.cache {
		if entry.createdAt.Before(t) {
			delete(c.cache, i)
		}
	}
}