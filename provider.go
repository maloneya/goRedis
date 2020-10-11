package main

import (
	"fmt"
)

type Provider interface {
	fetchKey(key string) string
}

type ClientWrapper interface {
	fetch(key string) string
}

type Cache interface {
	get(key string) string
	put(key, val string)
}

type RedisProvider struct {
	client ClientWrapper
	cache  Cache
}

func (p RedisProvider) fetchKey(key string) string {
	val := p.cache.get(key)
	if val == "" {
		val = p.client.fetch(key)
		p.cache.put(key, val)
	}
	return fmt.Sprintf("%v %v", key, val)
}

type LRUCache struct {
	values map[string]string
}

func (c LRUCache) get(key string) string {
	val, _ := c.values[key]
	return val
}

func (c LRUCache) put(key, val string) {
	c.values[key] = "cache val"
}

type RedisClientWrapper struct {
}

func (w RedisClientWrapper) fetch(key string) string {
	return "redis_val"
}
