package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "This is so exciting (up)"

	// spliting
	words := strings.Fields(text)

	// loop
	for i := 0; i < len(words); i++ {
		// Detect
		if words[i] == "(up)" && i > 0 {
			words[i-1]= strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
		
	}
	
	fmt.Println(words, " ")
}
