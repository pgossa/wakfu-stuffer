package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/handler"
)

func DataAddRoutes(rg *gin.RouterGroup) {
	config := rg.Group("/data")
	config.GET("/items", handler.StructWrapper(handler.GetItems))
}
