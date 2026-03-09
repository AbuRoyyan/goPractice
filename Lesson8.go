package main

import "fmt"

func main() {
	// These control how loops behave.

	// break

	// Stops a loop completely.
	// Example

	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// continue

	// Skips the current loop step.
	// // Example:

	for i := 1; i <= 5; i++ {

		if i == 3 {
			continue
		}
		fmt.Println(i)
	}



}