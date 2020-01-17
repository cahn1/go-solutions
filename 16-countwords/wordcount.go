package main

import (
	"bufio" // for reading from files
	"fmt"
	"os"
	"regexp"  // for regular expressions
	"strings" // for split and other string functions
)

func main() {
	// Get filename from command-line args
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s filename\n", os.Args[0])
		os.Exit(1)
	}
	// Bail out if file cannot be opened
	f, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	// Create a map to hold the words and counts
	counts := map[string]int{}

	// Read a line at a time
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		// Get rid of extra spaces by converting 2 spaces to 1
		line = strings.Replace(line, "  ", " ", -1)
		// Make it all lower case
		line = strings.ToLower(line)
		// Regexp to describe all non-alphanumeric characters
		reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
		if err != nil {
			panic(err)
		}
		// Remove all non-alphaumeric characters
		line = reg.ReplaceAllString(line, "")
		// Split line into words
		words := strings.Split(line, " ")

		// For each word, update its count in the map
		for i := range words {
			counts[words[i]]++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// Iterate through map, printing out key and value. It's actually
	// difficult to sort by values as opposed to keys (unlike Python),
	// so we won't do that here. Instead, run this program on hamlet.txt
	// and then pipe the result to sort, i.e.,
	//
	// go run wordcount.go hamlet.txt | sort -k 2 -rn | head -20
	//
	// Check out Andrew Gerrand's solution for sorting a map by value:
	// https://groups.google.com/forum/#!topic/golang-nuts/FT7cjmcL7gw
	for word, count := range counts {
		fmt.Println(word, count)
	}
}
