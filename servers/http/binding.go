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

	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	s.server.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
}
