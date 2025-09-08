package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var storage *URLStorage

const Port = ":8081"

func main() {
	var err error
	storage, err = NewStorage()
	if err != nil {
		log.Fatal(ErrorCreatingStorage, zap.Error(err))
	}

	InitLogging()

	InitConfig()

	r := gin.Default()

	r.POST("/v1/shorten", func(c *gin.Context) { handlerShorten(c) })

	r.GET("/v1/:short_code", func(c *gin.Context) { handleRedirect(c) })

	log.Info(EventServerStart, zap.String("port", conf.Port))
	if err := r.Run(conf.Port); err != nil {
		log.Fatal(ErrorRunningServer, zap.Error(err))
	}
}
