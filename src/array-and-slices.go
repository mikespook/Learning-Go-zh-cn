package main

func main() {
	var array [100]int     // Create array, index from 0 to 99
	slice := array[0:99]   // Create slice, index from 0 to 98

	slice[98] = 'a'	       // OK
	slice[99] = 'a'        // Error: "throw: index out of range"
}
