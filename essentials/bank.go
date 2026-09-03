package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func initBanking() {
	// Definitions
	balance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
	}

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

func writeBalanceToFile(balance float64) {
	// Write balance to file
	balanceText := fmt.Sprintf("%.2f", balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	fmt.Printf("Balance of %.2f written to file: %s\n", balance, accountBalanceFile)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0, errors.New("Failed to read balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
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
		writeBalanceToFile(*balance)
	} else {
		*balance -= transactionAmount
		writeBalanceToFile(*balance)
	}

	fmt.Println("Balance updated! Your new balance is:", *balance)

	// Return true for successful update
	return true
}
