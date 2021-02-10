package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) binding() {
	trackers := s.server.Group("/trackers")
	{
		trackers.GET("/:nodeID/info/:sec", func(c *gin.Context) {
			nodeIDStr := c.Param("nodeID")
			nodeID, _ := strconv.Atoi(nodeIDStr)

			secStr := c.Param("sec")
			sec, err := strconv.Atoi(secStr)
			if err != nil {
				log.Println(err.Error())
			}
			c.JSON(http.StatusOK, s.cleanerHandler.GetCleanedDataFor(nodeID, sec))
		})
	}

	s.server.StaticFS("/index", http.Dir("/Users/henry/Documents/goWorkSpace/go-viewer/servers/http/web"))
}
