package main

import (
	"fmt"
	// "strings"
)

func main() {

	var age int
	var vipPass string

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Print("Do you have a VIP pass (yes/no): ")
	fmt.Scanln(&vipPass)

	if age >= 18 {
		fmt.Println("Allow")

	} else if age < 18 && vipPass == "yes" {
		fmt.Println("Allow")

	} else {
		fmt.Println("Deny")
	}

	// var age int
	// var vipPass string

	// fmt.Print("Enter your age: ")
	// fmt.Scanln(&age)

	// if age >= 18 {
	// 	fmt.Println("Allow")

	// } else {

	// 	fmt.Print("Do you have a VIP pass (yes/no): ")
	// 	fmt.Scanln(&vipPass)
	// 	vipPass = strings.ToLower(vipPass)

	// 	if vipPass == "yes" {
	// 		fmt.Println("Allow")
	// 	} else {
	// 		fmt.Println("Deny")
	// 	}
	// }

}

// if age >= 18 (
// 	//allow
// ) else if age < 18 && vipPass == "yes" {
// 	// allow
// } else {
// 	// deny
// }
