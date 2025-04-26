package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
	tittle, content := getNOdeData()

	userNote, err := note.New(tittle, content)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNOdeData() (string, string) {
	title := getUserInput("enter")

	content := getUserInput("enter")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)

	return value
}
