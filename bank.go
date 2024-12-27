package main

import "fmt"

func main() {
	var accountBalance float64 = 1000
	fmt.Println("Welcome to Go Bank")

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		// wantsCheckBalance := choice == 1

		fmt.Print("Your choice:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Your deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			accountBalance += depositAmount

			if depositAmount <= 0 {
				fmt.Println("Invalid amount!")
				continue
			}

			fmt.Println("Balance updated! New Amount:", accountBalance)
		case 3:
			fmt.Print("Your withdrawal:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 || withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount!")
				continue
			}

			accountBalance -= withdrawalAmount

			fmt.Println("Balance updated! New Amount:", accountBalance)

		default:
			fmt.Println("Exiting...")
			fmt.Println("Thanks for using Go Bank")
			return
		}
	}
}
