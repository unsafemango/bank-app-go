package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"unsafemango.com/bank-app/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Mango's Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	newLine()

	// for loop to execute the code indefinitely
	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1\
		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
			newLine()
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 { // check for invalid user input
				fmt.Println("Invalid amount! Must be greater than 0.")
				newLine()
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			newLine()
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)

		case 3:
			fmt.Print("Your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance { // check that balance is greater than withdrawal amount
				fmt.Println("Not enough balance for this action! Balance:", accountBalance)
				newLine()
				continue
			} else if withdrawalAmount <= 0 { // check for invalid user input
				fmt.Println("Invalid amount! Must be greater than 0.")
				newLine()
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")

			return
		}

	}

}

func newLine() {
	fmt.Println("")
}
