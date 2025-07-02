package main

import (
	"fmt"
	"hash/crc32"
	"net/url"
)

const CodeSize = 7

func isValidUrl(longUrl string) bool {
	_, err := url.ParseRequestURI(longUrl)
	return err == nil
}
func generateShortCode(longURL string) string {
	hash := crc32.ChecksumIEEE([]byte(longURL))
	hashString := fmt.Sprintf("%010d", hash)

	return hashString[:CodeSize]
}
