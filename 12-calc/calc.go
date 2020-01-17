package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func main() {
	opmapper := map[rune]func(x, y int) int{
		'+': add,
		'-': sub,
		'*': mul,
		'/': div,
	}
	fmt.Print("Enter num op num: ")
	var num1, num2 int
	var op rune
	fmt.Scanf("%d %c %d", &num1, &op, &num2)
	fmt.Println(opmapper[op](num1, num2))
}
