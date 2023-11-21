package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StructWrapper(f func(c *gin.Context) (any, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
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

// May not be usefull
// func StructWrapperWithParameter(f func(c *gin.Context, parameter string) (any, error)) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		body, err := f(c, c.Param("parameter"))
// 		if err != nil {
// 			c.JSON(503, gin.H{"status": err})
// 			return
// 		}
// 		c.JSON(http.StatusOK, body)
// 	}
// }
