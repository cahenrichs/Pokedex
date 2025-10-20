package pokecache

import (
	"time"
)

type Cache struct {
	mu	*sync.Mutex
	cache	map[string]cacheEntry
}

type cacheEntry struct {
	createdAt	time.Time
	val	[]byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}
}

func (c *cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Lock()

	c.enties[key] = cacheEntry{
		createdAt: time.Now(),
		val:	val,
	}
}