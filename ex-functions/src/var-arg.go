package main

import "fmt"

func main() {
	prtthem(1, 4, 5, 7, 4)
	prtthem(1, 2, 4)
}

func prtthem(numbers ... int) { |\coderemark{numbers 现在是整数类型的~slice}|
	for _, d := range numbers {
		fmt.Printf("%d\n", d)
	}
}
