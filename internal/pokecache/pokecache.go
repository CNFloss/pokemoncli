package pokecache

import (
	"pokemoncli/internal/pokeapi"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type pokemonEntry struct {
	createdAt time.Time
	val pokeapi.Pokemon
}

type Cache struct {
	cache map[string]cacheEntry
}

type PokemonCache struct {
	Cache map[string]pokemonEntry
}

func NewPokemonCache(interval time.Duration) PokemonCache {
	c := PokemonCache{
		Cache: make(map[string]pokemonEntry),
	}
	go c.reapLoop(interval)
	return c
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if _, ok := c.cache[key]; !ok {
		return nil, false
	}
	return c.cache[key].val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	t := time.Now().UTC().Add(-interval)
	for i, entry := range c.cache {
		if entry.createdAt.Before(t) {
			delete(c.cache, i)
		}
	}
}

func (c *PokemonCache) Add(key string, val pokeapi.Pokemon) {
	c.Cache[key] = pokemonEntry{
		createdAt: time.Now().UTC(),
		val: val,
	}
}

func (c *PokemonCache) Get(key string) (pokeapi.Pokemon, bool) {
	if _, ok := c.Cache[key]; !ok {
		return pokeapi.Pokemon{}, false
	}
	return c.Cache[key].val, true
}

func (c *PokemonCache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *PokemonCache) reap(interval time.Duration) {
	t := time.Now().UTC().Add(-interval)
	for i, entry := range c.Cache {
		if entry.createdAt.Before(t) {
			delete(c.Cache, i)
		}
	}
}