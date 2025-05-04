package main

import "fmt"

func main() {
	// kreate key; value
	websites := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(websites["key1"])

	// add key; value
	websites["key3"] = "value3"
	fmt.Println(websites["key3"])

	// delete element
	delete(websites, "key1")
	fmt.Println(websites)
}
