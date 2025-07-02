package main

import (
	"fmt"
	"hash/crc32"
	"net/url"
	"strings"
)

const CodeSize = 7

func isValidUrl(longUrl string) bool {
	u, err := url.ParseRequestURI(longUrl)
	if err != nil {
		return false
	}

	if u.Scheme == "" || u.Host == "" {
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
