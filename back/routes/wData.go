package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/handlers"
)

func AddWdataRoutes(rg *gin.RouterGroup) {
	config := rg.Group("/Wdata")
	config.GET("/items", handlers.StructWrapper(handlers.GetWItems))
	config.GET("/actions", handlers.StructWrapper(handlers.GetWActions))
	config.GET("/itemProperties", handlers.StructWrapper(handlers.GetWItemProperties))
	config.GET("/itemTypes", handlers.StructWrapper(handlers.GetWItemTypes))
}
