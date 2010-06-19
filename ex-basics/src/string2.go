package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "dsjkdshdjsdh|\ldots|js"
	fmt.Printf("String %s\nLenght: %v, Runes: %d\n", str,
		len(strings.Bytes(str)), len(strings.Runes(str)))
}
