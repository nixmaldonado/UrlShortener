package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	//TODO ADD MIDDLEWARE RATE LIMITING r.Use(rateLimiter)

	r.POST("/v1/shorten", func(c *gin.Context) {
		handlerShorten(c)
	})

	r.GET("/v1/:short_url", func(c *gin.Context) {
		handleRedirect(c)
	})

	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
