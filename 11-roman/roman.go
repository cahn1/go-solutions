package main

import (
	"fmt"
)

func romanToArabic(numeral string) int {
	// Here we map runes to ints because each item in
	// a string is itself a rune (not a string). So
	// when we iterate through a string, we'll get a
	// sequence of runes.
	romanMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	// Create a slice big enough to hold numeric equivalents.
	// We'll actually make it one larger than it need to be
	// so that we don't fall off the end when we iterate thru
	// the slice of Arabic equivalents a bit later.

	arabicSlice := make([]int, len(numeral)+1)

	// Convert each Roman digit to Arabic equivalent and store
	for index, digit := range numeral {
		if val, inmap := romanMap[digit]; inmap {
			arabicSlice[index] = val
		} else {
			fmt.Printf("Invalid Roman digit: %c\n", digit)
			return 0
		}
	}
	// EXTRA: Now iterate through slice of ints, looking for
	// any values which are smaller than their neighbors. If
	// so, multiply that value by -1, making it negative. At
	// the same time, sum up the Arabic equivalents.

	sum := 0
	for index := 0; index < len(numeral); index++ {
		if arabicSlice[index] < arabicSlice[index+1] {
			arabicSlice[index] = -arabicSlice[index]
		}
		sum += arabicSlice[index]
	}
	return sum
}

func main() {
	romanStrings := []string{"MCLX", "MCMXCIX", "XLJI", "DCCIX"}
	for _, numeral := range romanStrings {
		fmt.Println(romanToArabic(numeral))
	}
}
