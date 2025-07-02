package main

import (
	"net/url"
)

func isValidUrl(longUrl string) bool {
	_, err := url.ParseRequestURI(longUrl)
	return err == nil
}
