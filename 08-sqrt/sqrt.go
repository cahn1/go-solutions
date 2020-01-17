package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func sqrt(num float64) float64 {
	oldGuess := 0.0
	newGuess := 1.0
	for math.Abs(oldGuess-newGuess) > 1e-8 {
		oldGuess = newGuess
		fmt.Println("...guess is", oldGuess)
		newGuess = oldGuess - (oldGuess*oldGuess-num)/(2*oldGuess)
	}
	return newGuess
}

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Println(sqrt(float64(num)))
}
