package storage

type Storage interface {
	Add(key string, value interface{}) bool
	Remove(key string) bool
	Get(key string) (interface{}, bool)
	GetCurLength() int
}
