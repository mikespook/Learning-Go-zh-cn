package main

import "fmt"

func main() {
	list := []string{"a", "b", "a", "a", "c", "d", "e", "f"}
	first := list[0]

	fmt.Printf("%s ", first)
	for _, v := range list[1:] {
		if first != v {
			fmt.Printf("%s ", v)
			first = v
		}
	}

}
