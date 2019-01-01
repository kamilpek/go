package main

import "fmt"

func main() {
	i := 7
	fmt.Println("memory address:", &i)
	inc1(i)
	fmt.Println("inc with func by value:", i)
	inc2(&i)
	fmt.Println("inc with func by pointer:", i)
}

func inc1(x int) {
	x++
}

func inc2(x *int) {
	*x++
}
