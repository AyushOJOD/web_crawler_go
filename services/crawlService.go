package services

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func CrawlService(url string, depth int, searchWord string) {
	if depth <= 0 {
		return
	}

	mu.Lock()
	if visited[url] {
		mu.Unlock()
		return
	}
	visited[url] = true
	mu.Unlock()

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Failed to fetch:", url)
		return
	}
	defer resp.Body.Close()

	links, text := ExtractService(resp.Body, url)

	if strings.Contains(strings.ToLower(text), strings.ToLower(searchWord)) {
		IndexText(url, text)
	}

	for _, link := range links {
		time.Sleep(100 * time.Millisecond)
		go CrawlService(link, depth-1, searchWord)
	}
}
