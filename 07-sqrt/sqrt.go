package main

import (
	"fmt"
	"os"
	"strconv"
)

func sqrt(num float64) float64 {
	guess := 1.0
	for i := 1; i <= 10; i++ {
		fmt.Println("...guess is", guess)
		guess = guess - (guess*guess-num)/(2*guess)
	}
	return guess
}

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Println(sqrt(float64(num)))
}
