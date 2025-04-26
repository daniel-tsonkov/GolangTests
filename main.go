package main

import "fmt"

func main() {

}

func getUserInput(prompt string) {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)
}
