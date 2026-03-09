package main

import "fmt"

func main() {
	username := "AbuRoyyan"
	password := "2026"

	var inputUser string
	var inputPass string
	
	fmt.Print("Enter username: ")
	fmt.Scanln(&inputUser)

	fmt.Print("Enter password: ")
	fmt.Scanln(&inputPass)

	if inputUser == username && inputPass == password {
		fmt.Println("Login Successful! Welcome Boss.")
	} else if inputUser == username && inputPass != password {
		fmt.Println("Wrong Passwprd.")
	} else {
		fmt.Println("User not found.")
	}
}
