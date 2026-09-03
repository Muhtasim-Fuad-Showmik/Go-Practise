package main

import "fmt"

func initBanking() {
	// Definitions
	balance := 1000.0

	fmt.Println("Welcome to Go Bank!")

MainLoop:
	for {
		// Output operations
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		// Input operations
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		// Logic operations
		switch choice {
		case 1:
			fmt.Println("Your balance is:", balance)
		case 2:
			updateBalance("deposit", &balance)
		case 3:
			updateBalance("withdraw", &balance)
		case 4:
			fmt.Println("Thank you for using Go Bank!")
			break MainLoop
		default:
			fmt.Println("Invalid choice")
			continue
		}
	}
}

func updateBalance(typeOfTransaction string, balance *float64) {
	// Get transaction amount
	fmt.Print("How much do you want to ", typeOfTransaction, "?: ")
	var transactionAmount float64
	fmt.Scan(&transactionAmount)

	// Validate entered amount
	if transactionAmount <= 0 {
		fmt.Println("Invalid amount. Please enter a positive number.")
		return
	}
	if typeOfTransaction == "withdraw" && transactionAmount > *balance {
		fmt.Println("Insufficient funds. Your balance is:", *balance)
		return
	}

	// Update balance
	if typeOfTransaction == "deposit" {
		*balance += transactionAmount
	} else {
		*balance -= transactionAmount
	}

	fmt.Println("Balance updated! Your new balance is:", *balance)
}
