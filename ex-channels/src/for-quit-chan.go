package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go shower(ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- false	// 或者是 true，这没啥关系
}

func shower(c chan int, quit chan bool) {
	for {
		select {
		case j := <-c:
			fmt.Printf("%d\n", j)
		case <-quit:
			break
		}
	}
}
