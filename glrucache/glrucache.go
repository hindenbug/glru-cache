package glrucache

import (
	"errors"
)

type Cache struct {
	capacity int
	Cache    map[interface{}]*CacheEntry
	Head     *CacheEntry
	Tail     *CacheEntry

	CacheOperations
}

type CacheEntry struct {
	nextEntry, prevEntry *CacheEntry
	Value                interface{}
	Key                  interface{}
}

type CacheOperations interface {
	Get(key string) (interface{}, int)
	Set(key string, value interface{}) int
}

func NewCache(capacity int) (*Cache, error) {
	if capacity > 0 {
		return &Cache{
			capacity: capacity,
			Cache:    make(map[interface{}]*CacheEntry, capacity),
		}, nil
	}

	return nil, errors.New("Cache capacity cannot be zero")
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	if entry, ok := c.Cache[key]; ok {
		c.removeCacheEntry(entry)
		c.moveToHead(entry)
		return entry.Value, true
	}

	return nil, false
}

func (c *Cache) Set(key, value interface{}) bool {
	if entry, exists := c.Cache[key]; exists {
		c.removeCacheEntry(entry)
	}

	newCacheEntry := &CacheEntry{Key: key, Value: value, nextEntry: c.Head, prevEntry: nil}
	c.moveToHead(newCacheEntry)

	// remove least recently used entry i.e .rear
	if len(c.Cache) > c.capacity {
		c.removeCacheEntry(c.Tail)
	}

	return true
}

func (c *Cache) moveToHead(entry *CacheEntry) {
	if c.Head != nil {
		c.Head.prevEntry = entry
		entry.nextEntry = c.Head
		c.Head = entry
	}

	if c.Head == nil && c.Tail == nil {
		c.Tail, c.Head = entry, entry
	}

	c.Cache[entry.Key] = entry
}

func (c *Cache) removeCacheEntry(entry *CacheEntry) {
	if entry == nil {
		return
	}

	if entry.prevEntry != nil {
		entry.prevEntry.nextEntry = entry.nextEntry
	}

	if entry.nextEntry != nil {
		entry.nextEntry.prevEntry = entry.prevEntry
	} else {
		c.Tail = entry.prevEntry
	}

	delete(c.Cache, entry.Key)
}
