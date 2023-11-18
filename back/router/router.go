package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/routes"
)

func InitRoutes(router *gin.RouterGroup) {
	api := router.Group("/api")
	routes.ApiAddRoutes(api)
	routes.WDataAddRoutes(api)
	routes.DataAddRoutes(api)
}
