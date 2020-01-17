package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func sqrt(num float64) (float64, error) {
	// Create a simple error to be returned
	negError := errors.New("Can't take sqrt of negative number!")

	if num < 0.0 {
		return 0.0, negError
	}
	oldGuess := 0.0
	newGuess := 1.0

	for math.Abs(oldGuess-newGuess) > 1e-8 {
		oldGuess = newGuess
		fmt.Println("...guess is", oldGuess)
		newGuess = oldGuess - (oldGuess*oldGuess-num)/(2*oldGuess)
	}
	return newGuess, nil
}

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	if result, err := sqrt(float64(num)); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
