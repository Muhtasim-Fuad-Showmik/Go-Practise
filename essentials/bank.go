package main

import (
	"fmt"

	"example.com/essentials/fileops"
)

const accountBalanceFile = "balance.txt"

func initBanking() {
	// Definitions
	balance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
	}

	fmt.Println("Welcome to Go Bank!")

MainLoop:
	for {
		presentOptions()

		// Input operations
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		// Logic operations
		switch choice {
		case 1:
			fmt.Println("Your balance is:", balance)
		case 2:
			success := updateBalance("deposit", &balance)
			if !success {
				continue MainLoop
			}
		case 3:
			success := updateBalance("withdraw", &balance)
			if !success {
				continue MainLoop
			}
		case 4:
			fmt.Println("Thank you for using Go Bank!")
			break MainLoop
		default:
			fmt.Println("Invalid choice")
			continue MainLoop
		}
	}
}

func updateBalance(typeOfTransaction string, balance *float64) bool {
	// Get transaction amount
	fmt.Print("How much do you want to ", typeOfTransaction, "?: ")
	var transactionAmount float64
	fmt.Scan(&transactionAmount)

	// Validate entered amount
	if transactionAmount <= 0 {
		fmt.Println("Invalid amount. Please enter a positive number.")
		return false
	}
	if typeOfTransaction == "withdraw" && transactionAmount > *balance {
		fmt.Println("Insufficient funds. Your balance is:", *balance)
		return false
	}

	// Update balance
	if typeOfTransaction == "deposit" {
		*balance += transactionAmount
		fileops.WriteFloatToFile(*balance, accountBalanceFile)
	} else {
		*balance -= transactionAmount
		fileops.WriteFloatToFile(*balance, accountBalanceFile)
	}

	fmt.Println("Balance updated! Your new balance is:", *balance)

	// Return true for successful update
	return true
}
