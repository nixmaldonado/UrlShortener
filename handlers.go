package main

import (
	"github.com/gin-gonic/gin"
	"log"
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

	log.Printf("Received shortCode: %q", shortCode)

	entry, err := storage.Get(shortCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if entry.URL == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	if err := storage.IncrementCounter(shortCode); err != nil {
		log.Printf("Error incrementing counter: %v", err)
	}

	c.Redirect(http.StatusFound, entry.URL)
}
