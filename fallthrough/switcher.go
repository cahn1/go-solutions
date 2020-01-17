package main

import "fmt"
import "math/rand"
import "time"

func main() {
	sec := time.Now().Unix()
	rand.Seed(int64(sec))
	num := rand.Int() % 1111
	fmt.Println("random number is", num)

	switch {
	case num > 1000:
		fmt.Println("greater than 1000")
		fallthrough

	case num > 100:
		fmt.Println("greater than 100")
		fallthrough

	case num > 10:
		fmt.Println("greater than 10")

	default:
		fmt.Println("less than 10")
	}
}
