package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/type/wakfuType"
)

type version struct {
	Version string `string:"Version"`
}

func GetVersion(c *gin.Context) (any, error) {
	v := version{
		Version: "1.0.0",
	}
	return v, nil
}

func GetWVersion(c *gin.Context) (any, error) {
	return wakfuType.WVersionData, nil
}
