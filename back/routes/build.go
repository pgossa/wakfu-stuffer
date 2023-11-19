package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/handlers"
)

// TODO: Add finer routes
func AddRankingRoutes(rg *gin.RouterGroup) {
	ranking := rg.Group("/build")
	ranking.GET("/:level", handlers.StructWrapper(handlers.GetLevelItems))
	ranking.POST("/ranking1", handlers.StructWrapper(handlers.GetBetterBuild))
	ranking.POST("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, c.Request.Body)
	})
}