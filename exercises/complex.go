package main

import (
	"fmt"
)

func main() {
	x := 1 + 2i
	y := 3 + 4i

	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}
