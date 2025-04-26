package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/note/note"
)

func main() {
	tittle, content := getNOdeData()

	userNote, err := note.New(tittle, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
}

func getNOdeData() (string, string) {
	title := getUserInput("enter")

	content := getUserInput("enter")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')

	return value
}
