package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/handler"
)

func ApiAddRoutes(rg *gin.RouterGroup) {
	config := rg.Group("/")
	config.GET("/version", handler.StructWrapper(handler.GetVersion))
	config.GET("/Wversion", handler.StructWrapper(handler.GetWVersion))
}
