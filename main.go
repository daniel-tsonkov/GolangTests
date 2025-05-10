package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	number := []int{1, 2, 3, 4}

	double := createTransformer(2)

	transformed := transformNumbers(&number, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&number, double)

	fmt.Println(transformed)
	fmt.Println(doubled)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

func getTranformenFunction() transformFn {
	return double
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
