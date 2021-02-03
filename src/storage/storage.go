package storage

import "sync"

// Storage is ...
type Storage struct {
	mu sync.Mutex
}

// NewStorage is the constructor
func NewStorage() *Storage {
	storageHandler = new(Storage)

	return storageHandler
}

// GetStorage should be used to return the handle of storage
func GetStorage() *Storage {
	if storageHandler == nil {
		panic("This package should be initialized before calling")
	}
	return storageHandler
}
