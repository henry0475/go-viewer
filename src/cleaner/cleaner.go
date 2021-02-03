package cleaner

import (
	"sync"
)

// Cleaner is ...
type Cleaner struct {
	nodeMap map[int]*trackedInfo
	mu      sync.Mutex
}

var maxSec = 60 * 15

// NewCleaner is the constructor
func NewCleaner() *Cleaner {
	var h = new(Cleaner)
	h.nodeMap = make(map[int]*trackedInfo)

	return h
}
