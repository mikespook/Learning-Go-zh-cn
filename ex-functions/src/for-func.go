package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		show(i)
	}
}

func show(j int) {
	fmt.Printf("%d\n", j)
}
