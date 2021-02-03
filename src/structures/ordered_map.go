package structures

import (
	"sync"
)

// OrderedMap is ...
type OrderedMap struct {
	m map[interface{}]interface{}
	k []interface{}

	mu sync.Mutex
}

// NewOrderedMap is the constructor
func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		m: make(map[interface{}]interface{}),
		k: make([]interface{}, 0),
	}
}

// Len returns the length of current map
func (o *OrderedMap) Len() int {
	return len(o.k)
}

// GetVal returns the value in the map
func (o *OrderedMap) GetVal(key interface{}) (val interface{}, ok bool) {
	if res, exist := o.m[key]; exist {
		return res, exist
	}

	return
}

// GetValByIndex should be used if you want to get the val from the ordered map with an index
func (o *OrderedMap) GetValByIndex(index int) (val interface{}, ok bool) {
	if index >= o.Len() {
		return
	}

	var key = o.k[index]
	if val, ok := o.m[key]; ok {
		return val, ok
	}
	return
}

// Exist determines whether a key does exist or not
func (o *OrderedMap) Exist(key interface{}) bool {
	if _, ok := o.m[key]; ok {
		return true
	}

	return false
}

// GetKeys should be used for getting the keys array
func (o *OrderedMap) GetKeys() (keys []interface{}, ok bool) {
	if len(o.k) == 0 {
		return
	}

	return o.k, true
}

// GetVals should be used for getting all values from the ordered map
func (o *OrderedMap) GetVals() (vals []interface{}, ok bool) {
	if len(o.k) == 0 {
		return
	}
	var outArr = make([]interface{}, o.Len())
	for i := 0; i < o.Len(); i++ {
		if v, ok := o.GetValByIndex(i); ok {
			outArr[i] = v
		}
	}

	return outArr, true
}

// Set will add the val into the map with the key.
// If the key already exists, it will update
func (o *OrderedMap) Set(key interface{}, val interface{}) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.Exist(key) {
		// If the key you inputted already exists
		o.m[key] = val
		return
	}

	// Add the key into the key arrary and map
	o.k = append(o.k, key)
	o.m[key] = val

	return
}
