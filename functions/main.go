package main

import "fmt"

func main() {
	sum := sumUp(1, 10, 15)
	numbers := []int{1, 10, 15, 40}
	anotherSum := sumUp(numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// Variadic Function: A function where the number of parameters is dynamic and therefore can vary
func sumUp(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
