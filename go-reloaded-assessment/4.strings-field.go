package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "hello world go"
	words := strings.Fields(text)

	fmt.Println(words)
}