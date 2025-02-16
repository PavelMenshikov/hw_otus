package main

import (
	"fmt"
	"strings"
	"unicode"
)

// countWords принимает текст и возвращает мапу, где каждому слову соответствует число его вхождений.
func countWords(text string) map[string]int {
	counts := make(map[string]int)
	lower := strings.ToLower(text)
	words := strings.FieldsFunc(lower, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	for _, word := range words {
		if word != "" {
			counts[word]++
		}
	}
	return counts
}

func main() {
	text := "Hello, world! Hello, Go. Go is great, and Go is fun."
	wordsCount := countWords(text)
	for word, count := range wordsCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
