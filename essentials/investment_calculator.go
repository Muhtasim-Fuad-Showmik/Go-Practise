package main

import (
	"fmt"
	"math"
)

func calculateInvestment() {
	const inflationRate = 6.5
	fmt.Printf("Inflation Rate: %f\n", inflationRate)

	var investmentAmount, years, expectedReturnRate float64

	fmt.Println("Enter the required informaiton:")
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Future Real Value: %.2f\n", futureRealValue)
}
