package main

import "fmt"

func main() {
	s := "foobar"
	a := []byte(s) |\coderemark{Again a conversion}|
	// Reverse a
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i] |\coderemark{Parallel assignment}|
	}
	fmt.Printf("%s\n", string(a)) |\coderemark{Convert it back}|
}
