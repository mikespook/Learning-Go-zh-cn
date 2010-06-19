package main

import "fmt"

func fibonacci(value int) []int {
	x := make([]int, value) |\longremark{We create an \key{array} to hold the integers up to the value given in the function call;}|
	x[0], x[1] = 1, 1 |\longremark{Starting point of the Fibonicai calculation;}|
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2] |\longremark{$x_n = x_{n-1} + x_{n-2}$;}|
	}
	return x |\longremark{Return the \emph{entire} array;}|
}

func main() {
	for _, term := range fibonacci(10) { |\longremark{Using the \key{range} keyword we "walk" the numbers returned by the fibonacci funcion. Here up to 10. And we print them.}|
		fmt.Printf("%v ", term)
	}
}
