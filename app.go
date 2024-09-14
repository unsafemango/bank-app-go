package main

import "fmt"

func main() {
	var accountBalance = 1000.00

	fmt.Println("Welcome to Mango's Bank!")

	// execute the code indefinitely
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		newLine()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		if choice == 1 { // view user's balance
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 { // deposit amount into user's balance
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
		} else if choice == 3 { // withdraw amount from user's balance
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
			newLine()
		} else { // exit program
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thanks for choosing our bank!")

}

func newLine() {
	fmt.Println("")
}
