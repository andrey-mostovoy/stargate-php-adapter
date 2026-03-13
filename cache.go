package main

import (
	lru "github.com/hashicorp/golang-lru/v2"
)

type PreparedCache struct {
	cache *lru.Cache[string, string]
}

func NewPreparedCache(size int) *PreparedCache {

	c, _ := lru.New[string, string](size)

	return &PreparedCache{
		cache: c,
	}
}

func (p *PreparedCache) Get(q string) (string, bool) {

	return p.cache.Get(q)
}

func (p *PreparedCache) Set(q string, id string) {

	p.cache.Add(q, id)
}