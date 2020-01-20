package storage

// Ensure interface (protocol) conformance
var _ StoreInterface = (*Store)(nil)

func NewStore() StoreInterface {
	var s StoreInterface = Store(make(map[string]interface{}))
	return s
}

func (s Store) Clone() StoreInterface {
	var (
		clone StoreInterface = NewStore()
		cld   Store          = clone.(Store)
	)
	for k, v := range s {
		cld[k] = v
	}
	return clone
}

func (s Store) Add(id string, value interface{}) {
	s[id] = value
}

func (s Store) Get(id string) (value interface{}, ok bool) {
	value, ok = s[id]
	return value, ok
}

func (s Store) Delete(id string) {
	delete(s, id)
}

func (s Store) Exists(id string) (ok bool) {
	_, ok = s[id]
	return ok
}

func (s Store) Values() (values []interface{}) {
	if len(s) == 0 {
		return nil
	}
	values = make([]interface{}, 0, len(s))
	for _, v := range s {
		values = append(values, v)
	}
	return values
}

func (s Store) Keys() (keys []string) {
	if len(s) == 0 {
		return nil
	}
	keys = make([]string, 0, len(s))
	for k, _ := range s {
		keys = append(keys, k)
	}
	return keys
}

func (s Store) Size() int {
	return len(s)
}

func (s Store) Drain() {
	if len(s) == 0 {
		return
	}
	var (
		keys []string = s.Keys()
	)
	for _, v := range keys {
		delete(s, v)
	}
	return
}
