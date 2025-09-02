package main

import (
	"fmt"
	"go.uber.org/zap"
	"hash/crc32"
	"net/url"
	"strings"
)

const CodeSize = 7

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
	hash := crc32.ChecksumIEEE([]byte(longURL))
	hashString := fmt.Sprintf("%010d", hash)

	return hashString[:CodeSize]
}
