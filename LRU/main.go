package main

import "github.com/Flakannon/sandbox/LRU/cache"

func main() {
	cache := cache.Initialise(3)

	cache.Add(1, 1)       // Cache is {1=1}
	cache.Add(2, 2)       // Cache is {1=1, 2=2}
	println(cache.Get(1)) // Returns 1, Cache is {2=2, 1=1}
	cache.Add(3, 3)       // Evicts key 2, Cache is {1=1, 3=3}
	println(cache.Get(2)) // Returns -1 (not found)
	cache.Add(4, 4)       // Evicts key 1, Cache is {3=3, 4=4}
	println(cache.Get(1)) // Returns -1 (not found)
	println(cache.Get(3)) // Returns 3
	println(cache.Get(4)) // Returns 4
}
