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

type RedisClientWrapper struct {
}

func (w RedisClientWrapper) fetch(key string) string {
	fmt.Print("hitme\n")
	return "redis_val"
}
