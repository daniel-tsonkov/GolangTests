package main

import "fmt"

func main() {
	websites := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(websites["key1"])
	websites["key3"] = "value3"
	fmt.Println(websites["key3"])
}
