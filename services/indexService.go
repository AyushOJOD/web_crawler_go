package services

import "strings"

func IndexText(url, text string) {
	mu.Lock()
	defer mu.Unlock()

	words := strings.Fields(strings.ToLower(text))
	for _, word := range words {
		if len(word) < 3 {
			continue
		}
		if _, found := index[word]; !found {
			index[word] = make(map[string]int)
		}
		index[word][url]++
	}
}
