package main

import (
	"fmt"
)

func minmaxavg(nums []float64) (float64, float64, float64, bool) {
	if len(nums) == 0 {
		return 0.0, 0.0, 0.0, false
	}
	min, max, total := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		} else if nums[i] > max {
			max = nums[i]
		}
		total += nums[i]
	}
	return min, max, total / float64(len(nums)), true
}

func main() {
	nums := []float64{-1.0, 4.5, -7.2, 65.4, -65.4, -4.5, -34.7, 1.0, -19.8, 19.8, 7.2, 34.7}
	if min, max, avg, success := minmaxavg(nums); success {
		fmt.Println(min, max, avg)
	} else {
		fmt.Println("empty slice?")
	}
}
