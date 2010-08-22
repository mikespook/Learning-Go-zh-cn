package main

func main() {
	s := "foobar"
	a := []byte(s)
	// Reverse a
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	println(string(a))
}
