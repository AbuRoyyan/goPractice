package main

import (
	"fmt"
	"strings"
)

func aOrAn(nextWord string) string {
	nextWord = strings.ToLower(nextWord)

	if nextWord == "" {
		return "a"
	}

	firstLetter := string(nextWord[0])
	vowelsAndH :=[]string{"a", "e", "i", "o", "u", "h"}

	if strings.HasPrefix(nextWord, "honest") ||
	strings.HasPrefix(nextWord, "hour") ||
	strings.HasPrefix(nextWord, "heir") {
		return "an"
	}

	for _, v := range vowelsAndH {
		if firstLetter == v {
			return "an"
		}
	}

	return "a"
}

func main() {
	fmt.Println(aOrAn("apple"))
	fmt.Println(aOrAn("book"))
	fmt.Println(aOrAn("honest"))
}