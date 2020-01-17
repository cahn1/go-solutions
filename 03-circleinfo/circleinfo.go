package main

import "fmt"
import "math"

func circleinfo(radius float64) (float64, float64) {
	return math.Pi * radius * radius, math.Pi * radius * 2.0
}

func main() {
	area, circ := circleinfo(6.5)
	fmt.Println(area, circ)
}
