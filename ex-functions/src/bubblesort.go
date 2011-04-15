package main

import "fmt"

func main() {
	n := []int{5, -1, 0, 12, 3, 5}
	fmt.Printf("unsorted %v\n", n)
	// 虽然 slice 是存储的值，但是它是个引用类型，因为底层的 array 已经改变了！
	bubblesort(n)
	fmt.Printf("sorted %v\n", n)
}

func bubblesort(n []int) {
	for i := 0; i < len(n) - 1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
}
