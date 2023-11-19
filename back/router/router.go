package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/routes"
)

func InitRoutes(router *gin.RouterGroup) {
	api := router.Group("/api")
	routes.AddVersionRoutes(api)
	routes.AddWdataRoutes(api)
	routes.AddDataRoutes(api)
	routes.AddDataCategoriesRoutes(api)
	routes.AddRankingRoutes(api)
}
