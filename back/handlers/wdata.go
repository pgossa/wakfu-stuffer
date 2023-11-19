package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/types/wakfuTypes"
)

func GetWItems(c *gin.Context) (any, error) {
	return wakfuTypes.WItemData, nil
}

func GetWActions(c *gin.Context) (any, error) {
	return wakfuTypes.WActionData, nil
}

func GetWItemProperties(c *gin.Context) (any, error) {
	return wakfuTypes.WItemPropertiesData, nil
}

func GetWItemTypes(c *gin.Context) (any, error) {
	return wakfuTypes.WItemTypesData, nil
}
