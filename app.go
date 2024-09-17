package main

import "fmt"

func main() {
	var accountBalance = 1000.00

	fmt.Println("Welcome to Mango's Bank!")

	// for loop to execute the code indefinitely
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
