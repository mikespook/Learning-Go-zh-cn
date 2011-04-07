package main

import (
	"time"
	"fmt"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(int64(sec) * 1e9)
	fmt.Println(w, "is ready!")
	c <- 1
}

func main() {
	c = make(chan int)
	go ready("Tee", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting, but not too long")
	<-c
	<-c
}
