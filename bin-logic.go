// Task 1 (Binary → Decimal)

// Write a Go program that converts this binary string to decimal:

// 101

// Expected result:

// 5

// Use Go's standard library.

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	num, err := strconv.ParseInt("101",2,64)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(num)
// }

// Task 2 (Finding the Word Before (bin))

// Create a slice of words like this:

// ["It","has","been","10","(bin)","years"]

// Write code that prints the word before (bin).

// Expected output:

// 10

// package main

// import (
// 	"fmt"
// 	"strings")

// func main() {
// 	text := "It has been 10 (bin) years"
// 	words := strings.Fields(text)

// 	fmt.Println(words[3])
// }

// Task 3 (Detecting (bin))

// Modify the slice so that (bin) is detected.

// Print a message like:

// Binary command found

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	text := "It has been (bin) years"
// 	words := strings.Fields(text)
// 	for i := 0; i < len(words); i++ {
// 		//Detect
// 		if words[i] == "(bin)" {
// 			fmt.Println("Binary command found")
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	text := "It has been (hex) years"
// 	words := strings.Fields(text)
// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(hex)" {
// 			fmt.Println("Hexa found.")
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	text := "It has been 1E (hex) years"

// 	// spliting of words
// 	words := strings.Fields(text)

// 	// loop
// 	for i := 0; i < len(words); i++ {

// 		// Dectecting  of hex
// 		if words[i] == "(hex)" {

// 			// checking the previous word
// 			prevWord := words[i-1]

// 			// fmt.Println("Our previous words is", prevWord)

// 			// convert from hexa to decimal
// 			num, err := strconv.ParseInt(prevWord, 16, 64)

// 			// relace the word
// 			if err == nil {
// 				words[i-1] = strconv.FormatInt(num, 10)

// 				// remove hex
// 				words = append(words[:i], words[i+1:]...)
// i--
// 			}
// 		}

// 	}
// 	result := strings.Join(words, " ")
// 	fmt.Println(result)
// }

//

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"strconv"
// )

// func main() {
// 	text := "It has been 10 (bin) years"

// 	// spliting
// 	words := strings.Fields(text)

// 	// looping
// 	for i := 0; i < len(words); i++ {

// 		// Dectecting bin

// 		// if words[i] == "(bin)"
// 		if i > 0 && words[i] == "(bin)" {

// 			// checking previous word
// 			prevWord := words[i-1]

// 			num, err := strconv.ParseInt(prevWord,2,64)

// 			if err == nil {

// 				// replacing the previous word
// 				words[i-1] = strconv.FormatInt(num, 10)

// 				// removing (bin)
// 				words = append(words[:i], words[i+1:]...)
// 				i--
// 			}

// 		}

// 	}
// 	result := strings.Join(words, " ")
// 	fmt.Println(result)
// }

 package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {

	text := "Simply add 42 (hex) and 10 (bin)"
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			num, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--

			// bin
		} else if words[i] == "(bin)" && i > 0 {
			num, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	result := strings.Join(words, " ")
	fmt.Println(result)
}
