package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/type/customType"
)

func GetItems(c *gin.Context) (any, error) {
	return customType.CustomItemsData, nil
}
