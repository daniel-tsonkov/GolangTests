package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//var appUser = user

	appUser, error := user.New(userFirstName, userLastName, userBirthdate)

	if error != nil {
		fmt.Println(error)
		return
	}

	admin := user.NewAdmin("asd@asd.asd", "asdasd")

	admin.OutputUserDetails()
	admin.ClearUserName()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
