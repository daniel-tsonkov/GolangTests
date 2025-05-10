package main

import "fmt"

type transformFn func(int) int

func main() {
	number := []int{1, 2, 3, 4}

	transformed := transformNumbers(&number, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)
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
