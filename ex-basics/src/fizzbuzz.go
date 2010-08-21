package main

import "fmt"

func main() {
	const (
		FIZZ = 3
		BUZZ = 5
	)
	for i := 1; i < 100; i++ {
		p := false
		if i%FIZZ == 0 {
			fmt.Printf("Fizz")
			p = true
		}
		if i%BUZZ == 0 {
			fmt.Printf("Buzz")
			p = true
		}
		if !p {
			fmt.Printf("%v", i)
		}
		fmt.Println()
	}
}
