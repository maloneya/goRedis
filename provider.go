package main

import (
	"fmt"
)

type Provider interface {
	fetchKey(key string) string
}

type RedisProvider struct {
}

func (p RedisProvider) fetchKey(key string) string {
	return fmt.Sprintf("%v %v", key, "val")
}
