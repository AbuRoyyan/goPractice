package main 

import (
	"fmt"
	"strings"
)

func capitalizeWord(word string) string {
	if word == "" {
		return ""
	}

	lower := strings.ToLower(word)

	return strings.ToUpper(string(lower[0])) + lower[1:]
}

func main() {
	fmt.Println(capitalizeWord("hELLo"))
}