// Conversion to ToUpper

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	text := "This is so exciting (up)"

// 	// spliting
// 	words := strings.Fields(text)

// 	// loop
// 	for i := 0; i < len(words); i++ {
// 		// Detect
// 		if words[i] == "(up)" && i > 0 {
// 			words[i-1]= strings.ToUpper(words[i-1])
// 			words = append(words[:i], words[i+1:]...)
// 			i--
// 		}

// 	}

// 	fmt.Println(words, " ")
// }


// Conversion to ToLower

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	text := "I AM SHOUTING (low)"
// 	words := strings.Fields(text)

// 	for i := 0; i < len(words); i++ {

// 		if words[i] == "(low)" && i > 0 {

// 			words[i-1] = strings.ToLower(words[i-1])

// 			words = append(words[:i], words[i+1:]...)
// 			i--

// 		}
// 	}
// 	strings.Join(words, " ")
// 	fmt.Println(words, " ")

// }

// Conversion to To CAP

package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "I AM shouting (cap)"
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		word := words[i-1]

		if len(word)> 1 {
			words[i-1] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		} else {
			words[i-1] = strings.ToUpper(word)
		}

		words = append(words[:i], words[i+1:]...)
		i--
	}
	strings.Join(words, " ")
	fmt.Println(words, " ")

}
