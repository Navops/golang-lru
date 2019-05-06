package lru

import (
	"sync"
	"time"

	"github.com/bserdar/golang-lru/simplelru"
)

// Cache is a thread-safe fixed size LRU cache.
type Cache struct {
	lru  simplelru.LRUCache
	lock sync.RWMutex
}

// New creates an LRU of the given size.
func New(size int) (*Cache, error) {
	return NewWithTTLEvict(size, 0, nil)
}

// NewWithTTL creates a LRU cache with the given size limit and a TTL
// for its elements
func NewWithTTL(sizeLimit int, ttl time.Duration) (*Cache, error) {
	return NewWithTTLEvict(sizeLimit, ttl, nil)
}

// NewWithEvict constructs a fixed size cache with the given eviction
// callback.
func NewWithEvict(size int, onEvicted func(key interface{}, value interface{}, size int)) (*Cache, error) {
	return NewWithTTLEvict(size, 0, onEvicted)
}

// NewWithTTLEvict constructs a lru cache with given size limit, ttl for elements, and an onEvicted callback
func NewWithTTLEvict(sizeLimit int, ttl time.Duration, onEvicted func(key interface{}, value interface{}, size int)) (*Cache, error) {
	lru, err := simplelru.NewLRUWithTTL(sizeLimit, ttl, simplelru.EvictCallback(onEvicted))
	if err != nil {
		return nil, err
	}
	c := &Cache{
		lru: lru,
	}
	return c, nil
}

// Purge is used to completely clear the cache.
func (c *Cache) Purge() {
	c.lock.Lock()
	c.lru.Purge()
	c.lock.Unlock()
}

// Add adds a value to the cache.  Returns true if an eviction occurred.
func (c *Cache) Add(key, value interface{}, size int) (evicted bool) {
	c.lock.Lock()
	evicted = c.lru.Add(key, value, size)
	c.lock.Unlock()
	return evicted
}

// Get looks up a key's value from the cache.
func (c *Cache) Get(key interface{}) (value interface{}, ok bool) {
	c.lock.Lock()
	value, ok = c.lru.Get(key)
	c.lock.Unlock()
	return value, ok
}

// Contains checks if a key is in the cache, without updating the
// recent-ness or deleting it for being stale.
func (c *Cache) Contains(key interface{}) bool {
	c.lock.RLock()
	containKey := c.lru.Contains(key)
	c.lock.RUnlock()
	return containKey
}

// Peek returns the key value (or undefined if not found) without updating
// the "recently used"-ness of the key.
func (c *Cache) Peek(key interface{}) (value interface{}, ok bool) {
	c.lock.RLock()
	value, ok = c.lru.Peek(key)
	c.lock.RUnlock()
	return value, ok
}

// ContainsOrAdd checks if a key is in the cache  without updating the
// recent-ness or deleting it for being stale,  and if not, adds the value.
// Returns whether found and whether an eviction occurred.
func (c *Cache) ContainsOrAdd(key, value interface{}, size int) (ok, evicted bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.lru.Contains(key) {
		return true, false
	}
	evicted = c.lru.Add(key, value, size)
	return false, evicted
}

// Remove removes the provided key from the cache.
func (c *Cache) Remove(key interface{}) {
	c.lock.Lock()
	c.lru.Remove(key)
	c.lock.Unlock()
}

// RemoveOldest removes the oldest item from the cache.
func (c *Cache) RemoveOldest() {
	c.lock.Lock()
	c.lru.RemoveOldest()
	c.lock.Unlock()
}

// Keys returns a slice of the keys in the cache, from oldest to newest.
func (c *Cache) Keys() []interface{} {
	c.lock.RLock()
	keys := c.lru.Keys()
	c.lock.RUnlock()
	return keys
}

// Len returns the number of items in the cache.
func (c *Cache) Len() int {
	c.lock.RLock()
	length := c.lru.Len()
	c.lock.RUnlock()
	return length
}

// Size returns the size of the cache
func (c *Cache) Size() int {
	c.lock.RLock()
	size := c.lru.Size()
	c.lock.RUnlock()
	return size
}
