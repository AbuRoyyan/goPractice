package main

import (
	"fmt"
	"strings"
)

func uppercaseLastN(words []string, n int) []string {
	start := len(words) - n 

	if start < 0 {
		start = 0
	}

	for i := start; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}

	return words
}

func main() {
	words := []string{"this", "is", "so", "exciting"}
	fmt.Println(uppercaseLastN(words, 2))
}