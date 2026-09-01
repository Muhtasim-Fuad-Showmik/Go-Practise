package main

import "fmt"

func main() {
	var userOption int
	fmt.Println("1. Calculate Investment Value")
	fmt.Println("2. Calculate Profit")
	fmt.Print("Enter your option: ")
	fmt.Scan(&userOption)

	if userOption == 1 {
		calculateInvestment()
	} else if userOption == 2 {
		calculateProfit()
	} else {
		fmt.Println("Invalid option")
	}
}
