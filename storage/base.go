package storage

type StoreBuilder func() StoreInterface
type Store map[string]interface{}

type StoreInterface interface {
	Add(string, interface{})
	Get(string) (interface{}, bool)
	Delete(string)
	Exists(string) bool
	Keys() []string
	Values() []interface{}
	Size() int
	Clone() StoreInterface
	Drain()
}
