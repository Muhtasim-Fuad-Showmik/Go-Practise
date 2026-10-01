package main

import (
	// "example.com/price-calculator/cmdmanager"
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		// For file processing
		fm := filemanager.New("prices.txt", fmt.Sprintf("tax_included_prices_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		// For terminal processing
		// cmdm := cmdmanager.New()
		// priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)

		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index, taxRate := range taxRates {
		// `select` waits for whichever channels earlier but does not wait for the other channel
		select {
		case err := <-errorChans[index]:
			if err != nil {
				panic("Could not process prices: " + err.Error())
			}
		case <-doneChans[index]:
			fmt.Println("Done processing prices for tax rate:", taxRate)
		}
	}
}
