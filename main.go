package main

import "fmt"

func main() {
	result := multiply(5)

	fmt.Println(result)
}

func multiply(number int) int {
	return number * multiply(number-1)
}
