package http

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/henry0475/go-viewer/src/cleaner"
)

// Server is ...
type Server struct {
	server *gin.Engine

	cleanerHandler *cleaner.Cleaner
}

// NewHTTPServer is the constructor
func NewHTTPServer(ctx context.Context, cleanerHandler *cleaner.Cleaner) *Server {
	var h = new(Server)
	h.cleanerHandler = cleanerHandler

	go func(ctx context.Context) {
		h.server = gin.New()

		h.server.Use(gin.Recovery())

		h.binding()

		select {
		case <-ctx.Done():
			return
		default:
			h.server.Run(":4000")
		}
	}(ctx)

	return h
}
