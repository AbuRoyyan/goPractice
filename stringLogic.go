package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var operation string

	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Println("Enter operation (add/sub/mult): ")
	fmt.Scanln(&operation)

	if operation == "add" {
		fmt.Println(num1 + num2)
	} else if operation == "sub" {
		fmt.Println(num1 - num2)
	} else if operation == "mult"{
		fmt.Println(num1 * num2)
	} else {
		fmt.Println("Unknown operation")
	}
}