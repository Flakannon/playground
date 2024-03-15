package cache

import "container/list"

type cacheKey int

type LRUCache struct {
	Capacity int
	Cache    map[cacheKey]*list.Element
	Order    *list.List
}

func Initialise(capacity int) *LRUCache {
	return &LRUCache{
		Capacity: capacity,
		Cache:    make(map[cacheKey]*list.Element),
		Order:    list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if element, found := c.Cache[cacheKey(key)]; found {
		c.Order.MoveToFront(element)
		return element.Value.(CacheItem).Value
	}
	return -1
}

func (c *LRUCache) Add(key int, value int) {
	// prep the cache item
	cacheItem := CacheItem{
		Key:   cacheKey(key),
		Value: value,
	}

	// if cache key exists in the cache already bring it to the front and exit early
	if item, exists := c.Cache[cacheItem.Key]; exists {
		c.Order.MoveToFront(item)
		return
	}

	// check if cache is at capacity
	// do some delete of the last element
	if len(c.Cache) == c.Capacity {
		lastElement := c.Order.Back()
		c.Order.Remove(lastElement)
		delete(c.Cache, lastElement.Value.(CacheItem).Key)
	}

	element := c.Order.PushFront(cacheItem)
	c.Cache[cacheItem.Key] = element
}
