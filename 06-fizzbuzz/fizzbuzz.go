package main

import (
	"fmt"
	"strconv"
)

func fizzbuzzX(num int) string {
	switch {
	case num%15 == 0:
		return "fizzbuzz"
	case num%5 == 0:
		return "buzz"
	case num%3 == 0:
		return "fizz"
	default:
		return strconv.Itoa(num)
	}
}

func fizzbuzz(num int) string {
	result := ""
	if num%3 == 0 {
		result += "fizz" // append "fizz"
	}
	if num%5 == 0 {
		result += "buzz" // append "buzz"
	}
	if result == "" { // neither "fizz" nor "buzz"
		result = strconv.Itoa(num)
	}
	return result
}

func main() {
	fmt.Println(fizzbuzz(3))
	fmt.Println(fizzbuzz(5))
	fmt.Println(fizzbuzz(15))
	fmt.Println(fizzbuzz(4))
}
