package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/handler"
)

func WDataAddRoutes(rg *gin.RouterGroup) {
	config := rg.Group("/Wdata")
	config.GET("/items", handler.StructWrapper(handler.GetWItems))
	config.GET("/actions", handler.StructWrapper(handler.GetWActions))
	config.GET("/itemProperties", handler.StructWrapper(handler.GetWItemProperties))
	config.GET("/itemTypes", handler.StructWrapper(handler.GetWItemTypes))
}
