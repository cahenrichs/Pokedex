package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	mu	*sync.Mutex
	cache	map[string]cacheEntry
	interval	time.Duration
}

type cacheEntry struct {
	createdAt	time.Time
	val	[]byte
}

// go
func NewCache(interval time.Duration) *Cache {
    c := &Cache{
        cache:    make(map[string]cacheEntry),
        mu:       &sync.Mutex{},
        interval: interval,
    }
    go c.reapLoop()
    return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:	value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	
}