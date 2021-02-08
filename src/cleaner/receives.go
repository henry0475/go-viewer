package cleaner

import (
	"log"

	collector "github.com/henry0475/go-viewer/servers/rpc/protos"
	"github.com/henry0475/go-viewer/src/structures"
)

// ReceivedRecord is ....
func (c *Cleaner) ReceivedRecord(request *collector.RecordTrackerRequest) {
	nodeData := c.getNodeDataFromMap(int(request.GetAgent().GetNodeID()))
	nodeData.NodeMetadata.NodeID = int(request.GetAgent().GetNodeID())
	nodeData.NodeMetadata.NumCPU = request.GetAgent().GetNumCPU()

	containerElement := nodeData.Containers.Back()
	if containerElement == nil {
		log.Println("containerElement is nil")
		return
	}
	if trackers, ok := containerElement.Value.(*structures.OrderedMap); ok { // trackers
		if !trackers.Exist(request.GetUUID()) {
			trackers.Set(request.GetUUID(), &tracker{
				UUID:    request.GetUUID(),
				Buckets: structures.NewOrderedMap(),
			})
		}

		if ti, ok := trackers.GetVal(request.GetUUID()); ok {

			var t = ti.(*tracker)

			for _, info := range request.GetInfo() {

				if !t.Buckets.Exist(info.GetBucket().GetHash()) {
					v := new(bucketDetail)
					v.name = info.Bucket.Name
					v.hash = info.GetBucket().GetHash()
					v.items = make([]*bucketItem, 0)
					t.Buckets.Set(info.GetBucket().GetHash(), v)
				}

				if bDetail, ok := t.Buckets.GetVal(info.GetBucket().GetHash()); ok {
					bd := bDetail.(*bucketDetail)
					bd.num = bd.num + 1
					bd.items = append(bd.items, &bucketItem{
						FuncName:  info.GetFuncName(),
						FilePath:  info.GetFilePath(),
						Line:      int(info.GetLine()),
						Duration:  info.GetDuration(),
						Timestamp: info.GetTimestamp(),
						Remarks:   info.GetRemarks(),
					})
				}
			}
		}
	} else {
		log.Println("00000011")
	}
}

// InitNodeCollector should be used for initializing the node map
func (c *Cleaner) InitNodeCollector(nodeID int) {
	if _, exist := c.nodeMap[nodeID]; !exist {
		c.mu.Lock()
		defer c.mu.Unlock()

		c.nodeMap[nodeID] = new(trackedInfo)
		c.nodeMap[nodeID].containerCreator()
		return
	}
}

func (c *Cleaner) getNodeDataFromMap(nodeID int) *trackedInfo {
	if m, exist := c.nodeMap[nodeID]; exist {
		return m
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.nodeMap[nodeID] = new(trackedInfo)
	c.nodeMap[nodeID].containerCreator()
	return c.nodeMap[nodeID]
}
