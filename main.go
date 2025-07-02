package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	storage, err := NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	//TODO ADD MIDDLEWARE RATE LIMITING r.Use(rateLimiter)

	r.POST("/v1/shorten", func(c *gin.Context) {
		handlerShorten(c, storage)
	})

	r.GET("/v1/:short_code", func(c *gin.Context) {
		handleRedirect(c, storage)
	})

	log.Println("Server starting on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
