package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	hobbies := [3]string{"sports", "Cookies", "Reading"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	mainHobbie := hobbies[0:2]
	fmt.Println(mainHobbie[0])

	courseGoals := []string{"Learn Go!", "Learn all"}
	fmt.Println(courseGoals)

	courseGoals[1] = "Learn details"
	courseGoals = append(courseGoals, "Learn all")
	fmt.Println(courseGoals)
}
