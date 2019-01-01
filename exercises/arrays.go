package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	a = append(a, 13)

	b := [5]int{1, 2, 3, 4, 5}
	b[2] = 7

	fmt.Println(a)
	fmt.Println(b)
}
