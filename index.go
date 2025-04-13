package main

import (
	"fmt"
	"os"
)

func writeBalanceToFIle(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accauntBalance float64 = 1000

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

		switch choice {
		case 1:
			fmt.Println("Баланс по сметка:", accauntBalance)
		case 2:
			writeBalanceToFIle(accauntBalance)
		case 3:
		case 4:
			fmt.Println("Чао!!!")
			return
		}
	}
}
