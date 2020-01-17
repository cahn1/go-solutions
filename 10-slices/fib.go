package main

import (
	"fmt"
)

func fib(n int) []int {
	if n <= 1 { // must ask for > 1 Fibonacci numbers
		return nil
	}
	nums := make([]int, n)
	nums[0], nums[1] = 1, 1
	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums
}

func main() {
	if fibs := fib(11); fibs != nil {
		fmt.Println(fibs)
	} else {
		fmt.Println("> 2, please!")
	}
}
