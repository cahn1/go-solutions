package main

import (
	"fmt"
	"os"
	"strconv"
)

func collatz(num int) {
	if num < 1 {
		fmt.Println("no neg")
		return
	}
	for num != 1 {
		fmt.Printf("%d ", num)
		switch {
		case num%2 == 0:
			num /= 2
		default:
			num = num*3 + 1
		}
	}
	fmt.Println(1)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s number\n", os.Args[0])
	} else {
		num, _ := strconv.Atoi(os.Args[1])
		collatz(num)
	}
}
