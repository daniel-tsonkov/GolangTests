package main

import "fmt"

func main() {
	result := multiply(5)

	fmt.Println(result)
}

func multiply(number int) int {
	if number == 0 {
		return 1
	}

	return number * multiply(number-1)
}
