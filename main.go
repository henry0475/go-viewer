package main

import (
	"context"

	"github.com/henry0475/go-viewer/servers/http"
	"github.com/henry0475/go-viewer/servers/rpc"
	"github.com/henry0475/go-viewer/src/cleaner"
	_ "github.com/henry0475/go-viewer/src/storage"
)

type handlers struct {
	cleanerHandler *cleaner.Cleaner
}

var allHandlers = &handlers{
	cleanerHandler: cleaner.NewCleaner(),
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// Start the gRPC server
	rpc.NewRPCServer(ctx, allHandlers.cleanerHandler)
	http.NewHTTPServer(ctx, allHandlers.cleanerHandler)

	select {
	case <-ctx.Done():
		cancel()
	}
}
