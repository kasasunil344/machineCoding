package main

import (
	"fmt"
	"github.com/kasasunil344/machineCoding/cache/policy"
	"github.com/kasasunil344/machineCoding/cache/storage"
)

type Cache struct {
	policy  policy.Policy
	storage storage.Storage
	size    int
}

func NewCache(size int) *Cache {
	return &Cache{
		policy:  policy.NewLRU(),
		storage: storage.NewMapStorage(),
		size:    size,
	}
}

func (c *Cache) Put(key string, val interface{}) error {
	if c.storage == nil {
		return fmt.Errorf("storage not initialized properly :(")
	}

	if c.policy == nil {
		return fmt.Errorf("policy not initialized properly :(")
	}

	if c.storage.GetCurLength() < c.size {
		ok := c.storage.Add(key, val)
		if !ok {
			return fmt.Errorf("error while adding into storage")
		}
	} else {
		fmt.Println("Evicting key!!!")
		keyToEvict, err := c.policy.EvictKey()
		if err != nil {
			return fmt.Errorf("error while evicting key")
		}
		// Remove from storage as well
		c.storage.Remove(keyToEvict)

		fmt.Println("Evicted key : ", keyToEvict)
		c.storage.Add(key, val)
	}

	err := c.policy.ProcessKey(key)
	if err != nil {
		return fmt.Errorf("error while processing key")
	}

	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	if c.storage == nil {
		return nil, fmt.Errorf("storage not initialized properly :(")
	}

	if c.policy == nil {
		return nil, fmt.Errorf("policy not initialized properly :(")
	}

	err := c.policy.ProcessKey(key)
	if err != nil {
		return nil, fmt.Errorf("key not processed properly")
	}

	value, ok := c.storage.Get(key)
	if !ok {
		return nil, fmt.Errorf("error getting value")
	}

	return value, nil
}
