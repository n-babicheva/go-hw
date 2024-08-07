package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/example/hello/reverse"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println("Input: ", scanner.Text())
		fmt.Println("Output:", reverse.String(scanner.Text()))
	}

	//var s string
	//
	//for {
	//	_, err := fmt.Scanln(&s)
	//	if err != nil {
	//		break
	//	}
	//
	//	fmt.Println("Input: ", s)
	//	fmt.Println("Output:", reverse.String(s))
	//}
}
