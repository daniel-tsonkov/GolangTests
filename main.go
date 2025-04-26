package main

import (
	"errors"
	"fmt"
)

func main() {
}

func getNOdeData(string, string) {
	title, err := getUserInput("enter")

	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := getUserInput("enter")

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("Invalit input")
	}
	return value, nil
}
