package main

import "fmt"

func main() {
	// These control how loops behave.

	// break

	// Stops a loop completely.
	// Example

	// for i := 1; i <= 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// continue

	// Skips the current loop step.
	// // Example:

	// for i := 1; i <= 5; i++ {

	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// Task 1 - Stop at 7

	// for i := 1; i <= 10; i++ {
	// 	if i == 7 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// Task 2 - Skip Number 5

	// for i := 1; i <= 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// Task 3 - Secret Number Game

	SecretNumber := 7
	var guess int

	for {
		fmt.Println("Guess the secret number:")
		fmt.Scanln(&guess)

		if guess == SecretNumber {
			fmt.Println("correct! You guessed the number.")
			break
		}
		fmt.Println("Wrong guess. Try again.")
	}



}