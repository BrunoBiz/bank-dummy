package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	// wantsCheckBalance := choice == 1

	fmt.Print("Your choice:")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit:")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		accountBalance += depositAmount

		fmt.Println("Balance updated! New Amount:", accountBalance)
	} else if choice == 3 {
		fmt.Print("Your withdrawal:")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)

		accountBalance -= withdrawalAmount

		fmt.Println("Balance updated! New Amount:", accountBalance)
	} else {

	}

	fmt.Println("Your choice: ", choice)
}
