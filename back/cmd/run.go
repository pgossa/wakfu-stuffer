package cmd

import (
	"log"
	"time"

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
	v1 := r.Group("/v1")
	customRouter.InitRoutes(v1)
	r.Run(":5000")
}
