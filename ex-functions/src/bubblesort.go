package main

import "fmt"

func main() {
	n := []int{5, -1, 0, 12, 3, 5}
	fmt.Printf("unsorted %v\n", n)
	// even though it's call by value, a slice is a
	// reference type, so the underlaying array is changed!
	bubblesort(n)
	fmt.Printf("sorted %v\n", n)
}

func bubblesort(n []int) {
	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
}
