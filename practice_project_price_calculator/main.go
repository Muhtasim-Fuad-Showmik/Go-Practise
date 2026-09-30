package main

import (
	// "example.com/price-calculator/cmdmanager"
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// For file processing
		fm := filemanager.New("prices.txt", fmt.Sprintf("tax_included_prices_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		// For terminal processing
		// cmdm := cmdmanager.New()
		// priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)

		priceJob.Process()
	}
}
