package services

import (
	"net/url"
	"strings"
)

func CleanService(href string) string {
	parsedHref, err := url.Parse(href)
	if err != nil || parsedHref.Scheme == "" {
		return strings.TrimPrefix(href, "/")
	}
	return parsedHref.String()
}
