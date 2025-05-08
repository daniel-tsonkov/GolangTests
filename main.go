package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4}

	doubled := transformNumbers(&number, triple)

	fmt.Println(doubled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
