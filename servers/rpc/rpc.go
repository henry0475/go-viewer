package rpc

import (
	"context"
	"log"
	"net"

	collector "github.com/henry0475/go-viewer/servers/rpc/protos"
	"github.com/henry0475/go-viewer/src/cleaner"
	"google.golang.org/grpc"
)

// GRPCServer is ...
type GRPCServer struct {
	grpc *grpc.Server
}

// NewRPCServer is the constructor
func NewRPCServer(ctx context.Context, cleanerHandler *cleaner.Cleaner) {
	var g = new(GRPCServer)

	go func(ctx context.Context) {
		// altsTC := alts.NewServerCreds(alts.DefaultServerOptions())
		// g.grpc = grpc.NewServer(grpc.Creds(altsTC))

		g.grpc = grpc.NewServer()

		collector.RegisterCollectorServer(g.grpc, &collectorServer{
			cleanerHandler: cleanerHandler,
		})

		listener, err := net.Listen("tcp", ":3000")
		if err != nil {
			log.Println("Error for listening the port of 3000")
		}
		select {
		case <-ctx.Done():
			return
		default:
			g.grpc.Serve(listener)
		}
	}(ctx)
}
