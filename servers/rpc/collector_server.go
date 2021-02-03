package rpc

import (
	"context"
	"fmt"
	"time"

	collector "github.com/henry0475/go-viewer/servers/rpc/protos"
	"github.com/henry0475/go-viewer/src/cleaner"
)

type collectorServer struct {
	cleanerHandler *cleaner.Cleaner
}

func (c *collectorServer) RecordTracker(ctx context.Context, request *collector.RecordTrackerRequest) (reply *collector.RecordTrackerReply, err error) {
	fmt.Printf(`
	** Got **
	Report UUID: %d
	Depth: %d
	`, request.GetUUID(), request.GetDepth())

	c.cleanerHandler.ReceivedRecord(request)

	reply = &collector.RecordTrackerReply{
		Status:    0,
		Timestamp: time.Now().Unix(),
	}
	return
}

func (c *collectorServer) RegisterNode(ctx context.Context, request *collector.RegisterNodeRequest) (reply *collector.RegisterNodeReply, err error) {
	fmt.Printf(`
	** Got **
	Node ID: %d
	Num CPU: %d
	`, request.GetAgent().GetNodeID(), request.Agent.GetNumCPU())

	c.cleanerHandler.InitNodeCollector(int(request.GetAgent().GetNodeID()))

	reply = &collector.RegisterNodeReply{
		Status:    0,
		Timestamp: time.Now().Unix(),
	}
	return
}
