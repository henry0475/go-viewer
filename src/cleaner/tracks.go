package cleaner

import (
	"container/list"
	"sync"
	"time"

	"github.com/henry0475/go-viewer/src/structures"
)

type nodeMetadata struct {
	NodeID int
	NumCPU int64
}

type trackedInfo struct {
	NodeMetadata nodeMetadata

	Containers *list.List
	count      int
	mu         sync.Mutex
}

type tracker struct {
	UUID    int64
	Buckets *structures.OrderedMap
}

type buckets *structures.OrderedMap

type bucketDetail struct {
	name  string
	hash  string
	num   int
	items []*bucketItem
}

type bucketItem struct {
	FuncName  string
	FilePath  string
	Line      int
	Duration  int64
	Timestamp int64
	Remarks   string
}

func (t *trackedInfo) containerCreator() {
	ticker := time.NewTicker(time.Second)
	t.Containers = list.New()

	go func(ticker *time.Ticker) {
		for {
			<-ticker.C

			t.mu.Lock()
			if t.count >= maxSec {
				t.Containers.Remove(t.Containers.Front())
				t.count = t.count - 1
			}
			t.Containers.PushBack(structures.NewOrderedMap()) // trackers
			t.count = t.count + 1
			t.mu.Unlock()
		}
	}(ticker)
}
