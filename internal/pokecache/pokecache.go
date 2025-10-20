package pokecache

import (
	"time"
)

type Cache struct {
	mu	sync.Mutex
	entries	map[string]cacheEntry
}

type cacheEntry struct {
	createdAt	time.Time
	val	[]byte
}