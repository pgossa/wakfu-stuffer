package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StructWrapper(f func(c *gin.Context) (any, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := f(c)
		if err != nil {
			c.JSON(503, gin.H{"status": err})
			return
		}
		c.JSON(http.StatusOK, body)
	}
}

func MessageWrapper(f func(c *gin.Context) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := f(c)
		if err != nil {
			c.JSON(503, gin.H{"status": err})
			return
		}
		c.JSON(http.StatusOK, body)
	}
}
