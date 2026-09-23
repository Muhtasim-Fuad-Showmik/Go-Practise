package recursion

import "fmt"

func main() {
	loopFactorialResult := loopFactorial(5)
	recursiveFactorialResult := recursiveFactorial(5)
	fmt.Println(loopFactorialResult)
	fmt.Println(recursiveFactorialResult)
}

func loopFactorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}

func recursiveFactorial(number int) int {
	if number == 1 {
		return 1
	}
	return number * recursiveFactorial(number-1)
}
