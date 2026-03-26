package main

import "fmt"

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
	fmt.Println(isPunctuation(","))
	fmt.Println(isPunctuation("x"))
}