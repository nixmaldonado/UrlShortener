package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

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

	if validURL(p.LongUrl) {
		// collision checking
		//database saving
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid url"})
	return
}

func handleRedirect(c *gin.Context) {

}

func validURL(u string) bool {
	_, err := url.ParseRequestURI(u)

	if err != nil {
		return false
	}

	return true
}
