package main

import "fmt"

func main() {
	var word string
	for {
		fmt.Print("Enter a word: ")
		fmt.Scanln(&word)
		if word == "quit" {
			break
		}
		if len(word)%2 == 0 {
			fmt.Println(word)
		} else {
			defer fmt.Println(word)
		}
	}
}
