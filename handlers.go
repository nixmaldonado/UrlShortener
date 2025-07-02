package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const ShortCode = "short_code"

type Payload struct {
	LongUrl string `json:"long_url"`
}

func handlerShorten(c *gin.Context) {
	p := new(Payload)
	err := c.ShouldBindJSON(&p)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if isValidUrl(p.LongUrl) {

		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid url"})
	return
}

func handleRedirect(c *gin.Context) {

}
