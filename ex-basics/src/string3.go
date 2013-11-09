package main

import (
	"fmt"
)

func main() {
	str := "asSASA ddd dsjkdsjs dk";
	runes := []rune(str)
	runes[4] = rune('a')
	runes[5] = rune('b')
	runes[6] = rune('c')
	fmt.Printf("Before: %s\n", str);
	fmt.Printf("After : %s\n", string(runes))
}
