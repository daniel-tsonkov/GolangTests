package main

import "fmt"

func main() {
	// var accauntBalance float64 = 1000

	fmt.Println("Поздравление")
	fmt.Println("Меню:")
	for {
		fmt.Println("1. Баланс по сметката")
		fmt.Println("2. Депозит")
		fmt.Println("3. Теглене")
		fmt.Println("4. Изход")

		var choice int
		fmt.Print("Избери вариант: ")
		fmt.Scan(&choice)

		fmt.Println("Избор:", choice)

		if choice == 4 {
			return
		}
	}
}
