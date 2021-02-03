package storage

var storageHandler *Storage

func init() {
	storageHandler = NewStorage()
}
