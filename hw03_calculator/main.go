package main

import "fmt"

func main() {
	var action string
	var a, b int
	fmt.Scan(&action, &a, &b)
	if action == "+" {
		fmt.Println(a + b)
	} else if action == "-" {
		fmt.Println(a - b)
	} else if action == "*" {
		fmt.Println(a * b)
	} else if action == "/" {
		if b == 0 {
			fmt.Println("Error: b mustn't be zero")
		} else {
			fmt.Println(a / b)
		}
	} else {
		fmt.Println("Error: enter correct data")
	}

}
