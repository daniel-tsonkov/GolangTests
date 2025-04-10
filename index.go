package main

import "fmt"

func main() {
	var text string = "aaa"
	var text1 string = "bbb"
	// var asd float64
	// fmt.Scan(&asd)

	// fmt.Printf(`%5f`, asd)
	myText(text)
	myText(text1)

}

func printText() {
	var sss = myText("sd")
	fmt.Println(sss)
}

func myText(str string) {
	return str
}
