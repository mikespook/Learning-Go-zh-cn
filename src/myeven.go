package main

import (
	"./even"
	"fmt"
)

func main() {
	i := 5
	fmt.Printf("Is %d even? %v\n", i, even.uneven(i))
}
