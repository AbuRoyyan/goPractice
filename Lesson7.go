package main

import "fmt"

func main() {

	// 1.Basic for Loop
	// Example

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// 2. Loop with User Input
	// Example:

	// var number int

	// fmt.Println("Enter a number:")
	// fmt.Scanln(&number)

	// for i := 1; i <= number; i++ {
	// 	fmt.Println(i)
	// }

	// 3. Infinte Loop
	// Sometimes programs must run forever.
	// Example:

	/*for {
		fmt.Println("This keeps running")
	}*/

	// Task 1 - Count to 10

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// Task 2 - Even Numbers
	// Print all even numbers from 2 to 20

	// for i := 2; i <= 20; i+=2 {
	// 	fmt.Println(i)
	// }

	// Task 3 - User Countdown

	var number int

	fmt.Println("Enter a number:")
	fmt.Scanln(&number)

	for i := number; i >= 1; i-- {
		fmt.Println(i)
	}

}
