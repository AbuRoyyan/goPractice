package main 

import (
	"fmt"
)

func joinwithpunctuation(tokens []string)string {

	result := ""

	for i, token := range tokens {
		if isPunctuation(token) {
			result += token
		} else {
			if i > 0 {
				result += " "
			}
			result += token
		}
	}
	return result
}

func isPunctuation(s string) bool {
	punctuations := []string{".", ",", "!", "?", ":", ";"}

	for _, p := range punctuations {
		if s == p {
			return true
		}
	}
	return false
}

func main() {
	tokens := []string{"hello", ",", "world", "!"}
	result := joinwithpunctuation(tokens)

	fmt.Println(result)
}