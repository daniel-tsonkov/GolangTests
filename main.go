package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4}

	doubled := doubleNumbers(&number)

	fmt.Println(doubled)
}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, value*2)
	}

	return dNumbers
}
