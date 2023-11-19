package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/handlers"
)

func AddVersionRoutes(rg *gin.RouterGroup) {
	config := rg.Group("/")
	config.GET("/version", handlers.StructWrapper(handlers.GetVersion))
	config.GET("/Wversion", handlers.StructWrapper(handlers.GetWVersion))
}
