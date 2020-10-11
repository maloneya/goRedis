package main

import (
	"container/ring"
	"sync"
	"time"
)

type Cache interface {
	get(key string) string
	put(key, val string)
}

type cacheEntry struct {
	put_time time.Time
	key      string
	value    string
}

func newCacheEntry(key, value string) cacheEntry {
	return cacheEntry{time.Now(), key, value}
}

type LRUCache struct {
	values     map[string]cacheEntry
	lru        *ring.Ring
	rwLock     sync.RWMutex
	globalExpr time.Duration
}

func NewCache(globalExpr time.Duration, capacity int) LRUCache {
	return LRUCache{
		make(map[string]cacheEntry),
		ring.New(capacity),
		sync.RWMutex{},
		globalExpr,
	}
}

func (c LRUCache) get(key string) string {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()

	val, _ := c.values[key]
	return val.value
}

func (c LRUCache) put(key, val string) {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()

	new_entry := newCacheEntry(key, val)

	//remove oldest element
	if c.lru.Value != nil {
		old_entry, _ := c.lru.Value.(cacheEntry)
		//Handle case where oldest elm expired and was replaced
		if old_entry.put_time == c.values[old_entry.key].put_time {
			delete(c.values, old_entry.key)
		}
	}

	c.lru.Value = new_entry
	c.lru.Next()

	c.values[key] = new_entry
	go c.expire(key, c.globalExpr)
}

func (c LRUCache) expire(key string, expr time.Duration) {
	time.Sleep(expr)
	c.rwLock.Lock()
	delete(c.values, key)
	c.rwLock.Unlock()
}
