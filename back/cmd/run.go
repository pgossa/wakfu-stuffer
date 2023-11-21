package cmd

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	customRouter "github.com/pgossa/wakfu-stuffer/router"
	"github.com/pgossa/wakfu-stuffer/utils"
)

var r = gin.Default()
var s = gocron.NewScheduler(time.UTC)

func Run() {
	_, err := s.Every(1).Days().Do(utils.CheckForNewVersion)
	if err != nil {
		log.Fatal(err)
	}
	s.StartAsync()
	// TODO: MODIFY THIS IS FOR TESTING PURPOSE ONLY
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins:    []string{"http://localhost:5173"},
		AllowMethods:  []string{"POST", "PUT", "PATCH"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	v1 := r.Group("/v1")
	customRouter.InitRoutes(v1)
	r.Run(":5000")
}
