package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num1, _ := strconv.Atoi(os.Args[1])
	num2, _ := strconv.Atoi(os.Args[2])

	switch os.Args[3] {
	case "+":
		fmt.Println(num1 + num2)

	case "-":
		fmt.Println(num1 - num2)

	default:
		fmt.Println("Huh?")
	}
}
