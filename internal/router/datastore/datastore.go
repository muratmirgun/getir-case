package datastore

import (
	"github.com/muratmirgun/getir-case/internal/store"
)

type DataStore struct {
	store.Store
}

func New(holder store.Store) *DataStore {
	return &DataStore{
		holder,
	}
}
