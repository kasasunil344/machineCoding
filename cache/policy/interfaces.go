package policy

type Policy interface {
	ProcessKey(key string) error
	EvictKey() (string, error)
}
