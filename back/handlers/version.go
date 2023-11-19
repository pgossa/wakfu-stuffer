package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/types/wakfuTypes"
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
	return wakfuTypes.WVersionData, nil
}
