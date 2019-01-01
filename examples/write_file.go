package main

import (
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	text := "hello\nworld\nwith\ngo\n"
	dat := []byte(text)
	err := ioutil.WriteFile("hello.txt", dat, 0644)
	check(err)
}
