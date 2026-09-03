package main

import "fmt"

func calculateProfit() {
	var revenue, expenses, taxRate float64

	revenue = getInput("Revenue: ")
	expenses = getInput("Expenses: ")
	taxRate = getInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func getInput(label string) (value float64) {
	fmt.Print(label)
	fmt.Scan(&value)
	return
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
