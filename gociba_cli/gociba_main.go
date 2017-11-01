package main

import (
	"gociba"
	"log"
	"os"
	"sync"
)

func main() {
	words := os.Args[1:]
	wordCount := len(words)
	group := &sync.WaitGroup{}
	group.Add(wordCount)
	results := make(chan *gociba.WordExplanation, wordCount)

	for _, word := range words {
		go func(w string) {
			defer group.Done()
			explain, err := gociba.LookupWord(w)
			if err != nil {
				log.Printf("Failed to get explanation for %v", w)
				return
			}
			results <- explain
		}(word)
	}

	group.Wait()
	close(results)

	for explain := range results {
		log.Printf("%v", explain)
	}
}
