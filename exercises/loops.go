package main

import "fmt"

func main() {
	fmt.Printf("classic for\n")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\nfor 'while'\n")
	j := 0
	for j < 3 {
		fmt.Println(j)
		j++
	}

	fmt.Printf("\nfor each\n")
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
	fmt.Printf("\nfor each hash\n")
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}

}
