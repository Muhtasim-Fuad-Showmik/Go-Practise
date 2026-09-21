package main

import "fmt"

// For complex functions defining custom function types can save characters
// from being repeatedly typed and maintained
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	double := createTransformer(2)
	triple := createTransformer(3)

	transformedNumbers := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformedNumbers)

	// Closeure: Closes created function via variable passed into factory method
	// Generated functions remain unaffected by multiple calls of the factory functions
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

// Passing functions as values into the parameters for another function
func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{}
	for _, val := range *numbers {
		tNumbers = append(tNumbers, transform(val))
	}

	return tNumbers
}

// Factory function for producing transformer functions
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
