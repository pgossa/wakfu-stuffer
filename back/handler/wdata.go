package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/type/wakfuType"
)

func GetWItems(c *gin.Context) (any, error) {
	return wakfuType.WItemData, nil
}

func GetWActions(c *gin.Context) (any, error) {
	return wakfuType.WActionData, nil
}

func GetWItemProperties(c *gin.Context) (any, error) {
	return wakfuType.WItemPropertiesData, nil
}

func GetWItemTypes(c *gin.Context) (any, error) {
	return wakfuType.WItemTypesData, nil
}
