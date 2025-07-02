package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const ShortCode = "short_code"

type Payload struct {
	LongURL string `json:"long_url"`
}

func handlerShorten(c *gin.Context, storage *URLStorage) {
	p := new(Payload)
	err := c.ShouldBindJSON(&p)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if isValidUrl(p.LongURL) {
		shortCode := generateShortCode(p.LongURL)

		shortCode, err := storage.Store(shortCode, p.LongURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"short_url": shortCode,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid url"})
	return
}

func handleRedirect(c *gin.Context, storage *URLStorage) {
	shortCode := c.Param(ShortCode)

	longUrl, err := storage.Get(shortCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if longUrl == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusFound, longUrl)
}
