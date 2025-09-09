package main

import (
	"crypto/sha256"
	"fmt"
	"go.uber.org/zap"
	"net/url"
	"strings"
)

func isValidUrl(longUrl string) bool {
	u, err := url.ParseRequestURI(longUrl)
	if err != nil {
		log.Error(ErrorParsingURL, zap.String("long_url", longUrl), zap.Error(err))
		return false
	}

	if u.Scheme == "" || u.Host == "" {
		log.Info(EventMissingSchemeOrHost, zap.String("url", u.String()))
		return false
	}

	host := strings.Trim(u.Host, ".")
	if host == "" || strings.HasPrefix(u.Host, ".") || strings.HasSuffix(u.Host, ".") {
		return false
	}

	return true
}

func generateShortCode(longURL string) string {
	hash := sha256.Sum256([]byte(longURL))
	hashStr := fmt.Sprintf("%x", hash)

	return hashStr[:conf.CodeSize]
}
