package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------")
		//panic("CANT CONTINUE")
	}

	fmt.Println("Welcome to Go Bank")
	fmt.Println(randomdata.PhoneNumber())

	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Exiting...")
			fmt.Println("Thanks for using Go Bank")
			return
		}
	}
}
