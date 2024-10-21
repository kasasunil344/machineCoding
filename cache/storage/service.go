package storage

import "fmt"

type MapStorage struct {
	data map[string]interface{}
}

func NewMapStorage() *MapStorage {
	return &MapStorage{
		data: make(map[string]interface{}),
	}
}

func (s *MapStorage) Add(key string, value interface{}) bool {
	s.data[key] = value
	return true
}

func (s *MapStorage) Remove(key string) bool {
	if _, ok := s.data[key]; ok {
		delete(s.data, key)
		return true
	}
	return false
}

func (s *MapStorage) Get(key string) (interface{}, bool) {
	if val, ok := s.data[key]; ok {
		return val, true
	}

	fmt.Println("Key not found")
	return nil, false
}

func (s *MapStorage) GetCurLength() int {
	return len(s.data)
}
