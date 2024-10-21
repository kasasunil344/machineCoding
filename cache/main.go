package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var size int
	if len(os.Args) < 2 {
		size = 3
	} else {
		size1, err := strconv.Atoi(os.Args[1])
		size = size1
		if err != nil {
			fmt.Println("Invalid cache size:", err)
			return
		}
	}

	// Initialize cache
	cache := NewCache(size)

	// Test cache logic
	err := cache.Put("firstMC", "cache implementation1")
	if err != nil {
		fmt.Println("Error putting key", err)
		return
	}

	err = cache.Put("firstMC1", "cache implementation2")
	if err != nil {
		fmt.Println("Error putting key", err)
		return
	}

	err = cache.Put("firstMC2", "cache implementation3")
	if err != nil {
		fmt.Println("Error putting key", err)
		return
	}

	err = cache.Put("firstMC3", "cache implementation4")
	if err != nil {
		fmt.Println("Error putting key", err)
		return
	}
	err = cache.Put("firstMC4", "cache implementation5")
	if err != nil {
		fmt.Println("Error putting key", err)
		return
	}

	val, err := cache.Get("firstMC")
	if err != nil {
		fmt.Println("Error getting key")
		return
	}
	fmt.Println("key Found : ", val)
}
