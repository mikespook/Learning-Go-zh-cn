package main

import (
	"time"
	"fmt"
)

func ready(w string, min int) {
	time.Sleep(min * 60 * 1e9)
	fmt.Println(w, "is ready!")
}

func main() {
	go ready("Tee", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
}
