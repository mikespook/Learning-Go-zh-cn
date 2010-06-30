package main

import "fmt"

func main() {
	fizz := 3
	buzz := 5
	for i := 1; i < 100; i++ {
		if i%fizz == 0 {
			fmt.Printf("Fizz")
		} else if i%buzz == 0 {
			fmt.Printf("FizzBuzz")
		} else {
			fmt.Printf("%v", i)
		}
		fmt.Println()

	}
}
