package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "John", age: 50}
	fmt.Println(p)
	fmt.Println("name:", p.name, "age:", p.age)
}
