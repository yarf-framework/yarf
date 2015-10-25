package yarf

import (
	"net/url"
	"sync"
)

// routeCache stores previously matched and parsed routes
type routeCache struct {
	route  Router
	params url.Values
}

// Cache is the service handler for route caching
type Cache struct {
	// Cache data storage
	storage map[string]routeCache

	// Sync Mutex
	mtx sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		storage: make(map[string]routeCache),
	}
}

func (c *Cache) Get(k string) (rc routeCache, ok bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	rc, ok = c.storage[k]
	return
}

func (c *Cache) Set(k string, r routeCache) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.storage[k] = r
}
