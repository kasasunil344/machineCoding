package policy

import (
	"container/list"
	"fmt"
)

type LRU struct {
	data  map[string]list.Element
	queue *list.List
}

func NewLRU() *LRU {
	return &LRU{
		data:  make(map[string]list.Element),
		queue: list.New(),
	}
}

func (p *LRU) ProcessKey(key string) error {
	if val, ok := p.data[key]; ok {
		p.queue.Remove(&val)
		p.queue.PushBack(list.Element{
			Value: key,
		})
	} else {
		p.queue.PushBack(list.Element{
			Value: key,
		})
		p.data[key] = list.Element{
			Value: key,
		}
	}

	return nil
}

func (p *LRU) EvictKey() (string, error) {
	front := p.queue.Front()
	if front == nil {
		return "", fmt.Errorf("empty queue")
	}
	// Remove from queue
	p.queue.Remove(front)

	ele := front.Value

	var keyToEvict string
	if val, ok := ele.(list.Element); ok {
		keyToEvict = val.Value.(string)
	}

	// Remove from map
	_, exists := p.data[keyToEvict]
	if !exists {
		return "", fmt.Errorf("key not found in data")
	}
	delete(p.data, keyToEvict) // Delete using the correct key

	return keyToEvict, nil
}
