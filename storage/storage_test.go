package storage

import (
	"fmt"
	"testing"
)

func TestDefaultStore(t *testing.T) {
	var (
		s StoreInterface = NewStore()
	)
	for i := 0; i < 10; i++ {
		var uid string = fmt.Sprintf("item_%d", i)
		s.Add(uid, i)
	}
	if s.Size() != 10 {
		t.Fatal("assertion failed.")
	}
	if len(s.Keys()) != 10 || len(s.Values()) != 10 {
		t.Fatal("inconsistent state.")
	}
	if !s.Exists("item_9") {
		t.Fatal("inconsistent state.")
	}
	s.Delete("item_9")
	if s.Exists("item_9") {
		t.Fatal("inconsistent state, key must be removed.")
	}

}
