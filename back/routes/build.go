package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/handlers"
)

// TODO: Add finer routes
func AddRankingRoutes(rg *gin.RouterGroup) {
	ranking := rg.Group("/build")
	ranking.POST("/quickRanking", handlers.StructWrapper(handlers.GetQuickBuild))
	// TODO: Remove test
	ranking.POST("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, c.Request.Body)
	})
}
