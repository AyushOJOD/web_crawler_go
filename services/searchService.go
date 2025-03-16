package services

import (
	"fmt"
	"strings"
)

func SearchService(word string) {
	word = strings.ToLower(word)
	mu.Lock()
	defer mu.Unlock()

	if urls, found := index[word]; found {
		fmt.Println("Results for:", word)
		for url, count := range urls {
			fmt.Printf(" - %s (%d times)\n", url, count)
		}
	} else {
		fmt.Println("No results found for:", word)
	}
}
