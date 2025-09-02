package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

const ShortCode = "short_code"

type Payload struct {
	LongURL string `json:"long_url"`
}

func handlerShorten(c *gin.Context) {
	log.Info(EventShortRequest)
	p := new(Payload)
	err := c.ShouldBindJSON(&p)

	if err != nil {
		log.Error(ErrorShortRequestPayload, zap.Error(err))
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

func handleRedirect(c *gin.Context) {
	shortCode := c.Param(ShortCode)

	log.Info(EventRedirectRequest, zap.String("short_code", shortCode))

	entry, err := storage.Get(shortCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if entry.URL == "" {
		log.Error(ErrorEmptyURL)
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	if err := storage.IncrementCounter(shortCode); err != nil {
		log.Error(ErrorIncrementingCounter, zap.Error(err))
	}

	c.Redirect(http.StatusFound, entry.URL)
}
