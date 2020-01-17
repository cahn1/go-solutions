package main

import "fmt"

func doubler(s *string) {
	doublerHelper(s) // pass the pointer by value to doublerHelper
}

func doublerHelper(s *string) {
	// This function receives a COPY of the pointer, but that's OK.
	// The pointer enables us to change the object it points to.
	*s = *s + *s
}

func main() {
	str := "Golang"
	doubler(&str)
	fmt.Println(str)
}
