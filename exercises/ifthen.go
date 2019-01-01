package main

import "fmt"

func main() {
	x := 5

	if x > 6 {
		fmt.Println("Więcej niż 6")
	} else if x < 2 {
		fmt.Println("Mniej niż 2")
	} else if x == 5 {
		fmt.Println("Równe 5")
	} else {
		fmt.Println("Poza zakresem")
	}
}
