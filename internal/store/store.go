package store

// Store Dynamic code structure implemented to add new features in the future
type Store interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}

// DataManager Dynamic code structure implemented to add new features in the future
type DataManager interface {
	Retrieve(input interface{}) (out interface{}, err error)
}
