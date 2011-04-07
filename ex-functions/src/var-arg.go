package main

import "fmt"

func main() {
	printthem(1, 4, 5, 7, 4)
	printthem(1, 2, 4)
}

func printthem(numbers ... int) { |\coderemark{\var{numbers} is now a slice of ints}|
	for _, d := range numbers {
		fmt.Printf("%d\n", d)
	}
}
