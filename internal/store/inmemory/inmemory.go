package inmemory

import (
	"errors"
	"sync"
)

// DataHolder : DataHolder struct
type DataHolder struct {
	HoldMap map[string]string
	Mutex   *sync.Mutex
}

func New() *DataHolder {
	return &DataHolder{
		HoldMap: make(map[string]string),
		Mutex:   &sync.Mutex{},
	}
}

func (d *DataHolder) Set(key string, val string) error {
	d.Mutex.Lock()
	d.HoldMap[key] = val
	d.Mutex.Unlock()
	return nil
}

func (d *DataHolder) Get(key string) (string, error) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	if val, ok := d.HoldMap[key]; ok {
		return val, nil
	}
	return "", errors.New("key not found in DataHolder")
}
