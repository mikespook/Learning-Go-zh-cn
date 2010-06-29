package main

import "fmt"

func main() {
	var a struct {
		name string
		age  int
	}

	a.name = "Pete"
	a.age = 42

	fmt.Printf("%v\n", a)
}
