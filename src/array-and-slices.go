package main

func main() {
	var array [100]int     // create array, index from 0 to 99
	slice := array[0:99]   // create slice, index from 0 to 98

	slice[98] = 'a'	       // OK
	slice[99] = 'a'        // Error: "throw: index out of range"
}
