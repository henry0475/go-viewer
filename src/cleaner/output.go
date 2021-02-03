package cleaner

import (
	"container/list"

	"github.com/henry0475/go-viewer/src/structures"
)

// Output is ...
type Output struct {
	NodeMetadata nodeMetadata
	Sec          int
	Containers   []*container
}

type container struct {
	Trackers []*exportedTracker
}

type exportedTracker struct {
	UUID    int64
	Buckets []*exportedBucket
}

type exportedBucket struct {
	Name  string
	Count int
	Items []*bucketItem
}

// GetCleanedDataFor returns the cleaned data set
func (c *Cleaner) GetCleanedDataFor(nodeID int, sec int) (out *Output) {
	if sec > maxSec || sec == 0 || nodeID == 0 {
		return
	}

	var exportedElements []*list.Element
	exportedElements = make([]*list.Element, 0, sec)
	out = new(Output)

	if info := c.getNodeDataFromMap(nodeID); info != nil {
		backElement := info.Containers.Back()
		for i := 0; i < sec && backElement != nil; i++ {
			exportedElements = append(exportedElements, backElement)
			backElement = backElement.Prev()
		}

		out.Sec = sec
		out.NodeMetadata = info.NodeMetadata
		out.Containers = make([]*container, 0, len(exportedElements))

		for _, element := range exportedElements {
			if element != nil {

				var trackers = element.Value.(*structures.OrderedMap)

				var nContainer = new(container)
				nContainer.Trackers = make([]*exportedTracker, 0, trackers.Len())

				if keys, ok := trackers.GetKeys(); ok {
					for _, key := range keys {

						if v, ok := trackers.GetVal(key); ok {
							vTracker := v.(*tracker)
							var nTracker = new(exportedTracker)
							nTracker.UUID = vTracker.UUID
							nTracker.Buckets = make([]*exportedBucket, 0)

							if hashKeys, ok := vTracker.Buckets.GetKeys(); ok {
								for _, hashKey := range hashKeys {
									if vk, ok := vTracker.Buckets.GetVal(hashKey); ok {
										bDetail := vk.(*bucketDetail)
										var eBucket = new(exportedBucket)
										eBucket.Name = bDetail.name
										eBucket.Count = bDetail.num
										eBucket.Items = bDetail.items

										nTracker.Buckets = append(nTracker.Buckets, eBucket)
									}
								}
							}

							nContainer.Trackers = append(nContainer.Trackers, nTracker)
						}

					}
				}

				out.Containers = append(out.Containers, nContainer)
			}
		}
	}

	return
}
