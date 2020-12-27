package main

import (
	"fmt"
	lru "glru-cache/glrucache"
)

func main() {
	c, err := lru.NewCache(5)

	if err != nil {
		panic(err)
	}

	// Push data to cache
	c.Set("0", "a")
	c.Set("1", "b")
	c.Set("2", "c")
	c.Set("3", "d")
	c.Set("4", "e")
	fmt.Println("Cache:", c.Head, c.Tail)

	// Get data from the cache
	c.Get("0")
	fmt.Println("Cache:", c.Head, c.Tail)

	// Update data
	c.Set("1", "f")
	fmt.Println("Cache:", c.Head, c.Tail)

	// Push new data
	c.Set("5", "q")
	fmt.Println("Cache:", c.Head, c.Tail)
	c.Set("6", "w")
	fmt.Println("Cache:", c.Head, c.Tail)
	c.Set("7", "we")
	fmt.Println("Cache:", c.Head, c.Tail)
	c.Set("1", "90")
	fmt.Println("Cache:", c.Head, c.Tail)

	c.Get("6")
	fmt.Println("Cache:", c.Head, c.Tail)
}
