package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const acountBalanceFile = "balance.txt"

func writeBalanceToFIle(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(acountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(acountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to read file")
	}

	balanceText, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 1000, errors.New("Failed to parce file")
	}

	return balanceText, nil
}

func main() { //fake upload
	//fake upload 15 apr 2025
	//test upload!!!
	var accauntBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("-----------")
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
	}

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
