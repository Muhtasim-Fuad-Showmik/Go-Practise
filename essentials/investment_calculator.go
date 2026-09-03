package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func calculateInvestment() {
	fmt.Printf("Inflation Rate: %f\n", inflationRate)

	var investmentAmount, years, expectedReturnRate float64

	fmt.Println("Enter the required informaiton:")
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(
		investmentAmount,
		expectedReturnRate,
		years,
	)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)
	fmt.Print(formattedFV)
	fmt.Print(formattedFRV)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue

}
