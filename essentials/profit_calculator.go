package main

import (
	"errors"
	"fmt"
	"os"
)

const profitCalculatorFile = "profit_calculator.txt"

func calculateProfit() {
	var revenue, expenses, taxRate float64
	var err error

	for _, label := range []string{"Revenue: ", "Expenses: ", "Tax Rate: "} {
		var value float64
		value, err = getInput(label, true)
		if err != nil {
			panic(err)
		}
		switch label {
		case "Revenue: ":
			revenue = value
		case "Expenses: ":
			expenses = value
		case "Tax Rate: ":
			taxRate = value
		}
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

	writeResultsToFile(ebt, profit, ratio)
}

func getInput(label string, onlyPositive bool) (value float64, err error) {
	fmt.Print(label)
	fmt.Scan(&value)

	if onlyPositive && value <= 0 {
		return 0, errors.New("Value must be greater than 0")
	}

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
