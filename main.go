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

	r := gin.Default()

	//TODO ADD MIDDLEWARE RATE LIMITING r.Use(rateLimiter) or delegate to API Gateway

	r.POST("/v1/shorten", func(c *gin.Context) { handlerShorten(c) })

	r.GET("/v1/:short_code", func(c *gin.Context) { handleRedirect(c) })

	log.Info(EventServerStart, zap.String("port", Port))
	if err := r.Run(Port); err != nil {
		log.Fatal(ErrorRunningServer, zap.Error(err))
	}
}
