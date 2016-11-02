package yarf

import (
	"sync"
)

/*
// RouteCache stores previously matched and parsed routes
type RouteCache struct {
	route  []Router
	params Params
}
*/

// Cache is the service handler for route caching
type Cache struct {
	// Cache data storage
	//storage map[string]RouteCache
	storage map[string]Router

	// Sync Mutex
	sync.RWMutex
}

// NewCache creates and initializes a new Cache service object.
func NewCache() *Cache {
	return &Cache{
		//storage: make(map[string]RouteCache),
		storage: make(map[string]Router),
	}
}

// Get retrieves a routeCache object by key name.
//func (c *Cache) Get(k string) (rc RouteCache, ok bool) {
func (c *Cache) Get(k string) (rc Router, ok bool) {
	c.RLock()
	defer c.RUnlock()

	rc, ok = c.storage[k]
	return
}

// Set stores a routeCache object under a key name.
//func (c *Cache) Set(k string, r RouteCache) {
func (c *Cache) Set(k string, r Router) {
	c.Lock()
	defer c.Unlock()

	c.storage[k] = r
}
