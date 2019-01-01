package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Printf("sum\n")
	sum_result := sum(2, 3)
	fmt.Println(sum_result)

	fmt.Printf("\nsqrt\n")
	sqrt_result, sqrt_err := sqrt(16)
	if sqrt_err != nil {
		fmt.Println(sqrt_err)
	} else {
		fmt.Println(sqrt_result)
	}
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Negative number")
	}

	return math.Sqrt(x), nil
}
