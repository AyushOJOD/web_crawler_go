package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"webcrawler/services"
)

var baseURL = "https://usf-cs272-s25.github.io/top10/"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a word to search (or type 'exit' to quit): ")
		scanner.Scan()
		word := scanner.Text()

		if word == "exit" {
			fmt.Println("Exiting...")
			break
		}

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			services.CrawlService(baseURL, 3, word) 
		}()

		wg.Wait()

		services.SearchService(word)
	}
}
