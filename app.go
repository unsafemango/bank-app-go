package main

import "fmt"

func main(){
	fmt.Println("Welcome to Mango's Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	
	if choice == 1 {
		
	}

	fmt.Println("Your choice:", choice)
}