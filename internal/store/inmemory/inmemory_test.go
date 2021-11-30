package inmemory

import (
	"testing"
)

func TestGetMemory(t *testing.T) {
	store := New()

	err := store.Set("example-key", "example-value")
	if err != nil {
		return
	}
	v, e := store.Get("example-key")
	if e != nil {
		t.Fail()
	}

	if v != "example-value" {
		t.Fail()
	}
}

func TestSetMemory(t *testing.T) {
	store := New()
	err := store.Set("example-key", "example-value")
	if err != nil {
		return
	}
	if k, ok := store.HoldMap["example-key"]; !ok {
		t.Fail()
	} else if k != "example-value" {
		t.Fail()
	}
}
