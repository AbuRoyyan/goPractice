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

// package main

// import (
// 	"fmt"
// 	"strings")

// func main() {
// 	text := "It has been 10 (bin) years"
// 	words := strings.Fields(text)

// 	fmt.Println(words[3])
// }

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

