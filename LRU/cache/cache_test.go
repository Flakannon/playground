package cache

import (
	"container/list"
	"log"
	"testing"
)

func TestGetItemFromCache(t *testing.T) {
	c := LRUCache{
		Capacity: 1,
		Cache:    make(map[cacheKey]*list.Element),
		Order:    list.New(),
	}

	cacheItem := CacheItem{
		Key:   1,
		Value: 100,
	}

	log.Print(cacheItem)

	element := c.Order.PushFront(cacheItem)
	c.Cache[cacheItem.Key] = element

	actual := c.Get(1)

	expected := 100

	if expected != actual {
		t.Errorf("test failed: expected %v, but got %v", expected, actual)
	}
}

func TestAddItemToCache(t *testing.T) {
	c := LRUCache{
		Capacity: 3,
		Cache:    make(map[cacheKey]*list.Element),
		Order:    list.New(),
	}

	tests := []struct {
		testName      string
		cache         LRUCache
		cacheItems    []CacheItem
		expectedFront int
		expectedBack  int
	}{
		{
			testName: "Add simple Item to cache",
			cache:    c,
			cacheItems: []CacheItem{
				{
					Key:   1,
					Value: 100,
				},
			},
			expectedFront: 100,
			expectedBack:  100,
		},
		{
			testName: "Max out cache capacity and fetch most recent entry",
			cache:    c,
			cacheItems: []CacheItem{
				{
					Key:   1,
					Value: 100,
				},
				{
					Key:   2,
					Value: 200,
				},
				{
					Key:   3,
					Value: 300,
				},
			},
			expectedFront: 300,
			expectedBack:  100,
		},
		{
			testName: "Cache doesnt bust and least recent entry is dropped out ",
			cache:    c,
			cacheItems: []CacheItem{
				{
					Key:   1,
					Value: 100,
				},
				{
					Key:   2,
					Value: 200,
				},
				{
					Key:   3,
					Value: 300,
				},
				{
					Key:   4,
					Value: 400,
				},
			},
			expectedFront: 400,
			expectedBack:  200,
		},
	}

	for _, tt := range tests {
		for _, cacheItem := range tt.cacheItems {
			tt.cache.Add(int(cacheItem.Key), cacheItem.Value)
		}

		frontOfCache := tt.cache.Order.Front().Value.(CacheItem)
		backOfCache := tt.cache.Order.Back().Value.(CacheItem)

		if frontOfCache.Value != tt.expectedFront {
			t.Errorf("Expected most recent cache entry t be %v, but it actually is %v for test %s", tt.expectedFront, frontOfCache.Value, tt.testName)
		}
		if backOfCache.Value != tt.expectedBack {
			t.Errorf("Expected least recent cache entry t be %v, but it actually is %v for test %s", tt.expectedBack, backOfCache.Value, tt.testName)
		}
	}
}
