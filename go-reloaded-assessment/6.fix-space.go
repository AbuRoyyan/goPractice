package main

import (
	"fmt"
	"regexp"
)

func fixSingleQuotes(text string) string {
	result := text

	re := regexp.MustCompile(`'\s*(.*?)\s*'`)

	result = re.ReplaceAllStringFunc(result, func(match string) string {
		inner := re.FinderStringSubmatch(match)[1]
		return "'" + inner + "'"
	})

	return result
}

func main() {
	fmt.Println(fixSingleQuotes("' hello world '"))
}