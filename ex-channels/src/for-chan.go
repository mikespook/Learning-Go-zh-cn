package main

import "fmt"

func main() {
	ch := make(chan int)
	go shower(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func shower(c chan int) {
	for {
		j := <-c
		fmt.Printf("%d\n", j)
	}
}
