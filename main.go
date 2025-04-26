package main

import (
	"errors"
	"fmt"
)

func main() {
	tittle, content, err := getNOdeData()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNOdeData() (string, string, error) {
	title, err := getUserInput("enter")

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	content, err := getUserInput("enter")

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return title, content, nil
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
