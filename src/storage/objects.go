package storage

import "errors"

// TODO 这里可能存在溢出，当存入的objects逐渐增加，内存占用也在增加，以后处理

var objectsMap map[string]interface{}

// GetObject will return an object
func (s *Storage) GetObject(name string) (obj interface{}, err error) {
	obj, ok := objectsMap[name]
	if !ok {
		err = errors.New("No object found")
		return
	}

	return
}

// SaveObject will save an object
func (s *Storage) SaveObject(name string, obj interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	objectsMap[name] = obj
}
