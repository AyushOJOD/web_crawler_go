package services

import (
	"io"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func ExtractService(body io.Reader, baseURL string) ([]string, string) {
	var links []string
	var textContent strings.Builder

	tokenizer := html.NewTokenizer(body)

	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			return links, textContent.String()
		case html.TextToken:
			text := strings.TrimSpace(html.UnescapeString(string(tokenizer.Text())))
			if text != "" {
				textContent.WriteString(text + " ")
			}
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						href := attr.Val
						absoluteURL := ResolveURL(baseURL, href)
						links = append(links, absoluteURL)
					}
				}
			}
		}
	}
}

func ResolveURL(base, href string) string {
	baseParsed, err := url.Parse(base)
	if err != nil {
		return ""
	}
	hrefParsed, err := url.Parse(href)
	if err != nil {
		return ""
	}
	return baseParsed.ResolveReference(hrefParsed).String()
}
