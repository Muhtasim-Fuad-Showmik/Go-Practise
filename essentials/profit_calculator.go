package main

import (
	"fmt"
	"os"
)

const profitCalculatorFile = "profit_calculator.txt"

func calculateProfit() {
	var revenue, expenses, taxRate float64

	revenue = getInput("Revenue: ")
	if revenue <= 0 {
		panic("Revenue must be greater than 0")
	}
	expenses = getInput("Expenses: ")
	if expenses <= 0 {
		panic("Expenses must be greater than 0")
	}
	taxRate = getInput("Tax Rate: ")
	if taxRate <= 0 {
		panic("Tax Rate must be greater than 0")
	}

	// Store calculated results into file
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

	writeResultsToFile(ebt, profit, ratio)
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

func writeResultsToFile(ebt, profit, ratio float64) {
	resultsText := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	os.WriteFile(profitCalculatorFile, []byte(resultsText), 0644)
	fmt.Printf("Results written to file: %s\n", profitCalculatorFile)
}
