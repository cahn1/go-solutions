package main

import (
	"fmt"
	"os"
	"strconv"
)

func factorial(num int) int {
	result := num
	for i := 1; i < num; i++ {
		result *= i
	}
	return result
}

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Println(factorial(num))
}
