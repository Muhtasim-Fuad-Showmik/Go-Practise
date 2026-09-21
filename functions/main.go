package main

import "fmt"

// For complex functions defining custom function types can save characters
// from being repeatedly typed and maintained
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	transformedNumbers := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformedNumbers)
}

// Passing functions as values into the parameters for another function
func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{}
	for _, val := range *numbers {
		tNumbers = append(tNumbers, transform(val))
	}

	return tNumbers
}