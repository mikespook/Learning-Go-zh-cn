package main

import "fmt"

func main() {
	var array [100]int      // Create array, index from 0 to 99
	slice := array[0:100]   // Create slice, index from 0 to 99

	fmt.Printf("%d %d  - %d %d\n",
		cap(array), len(array), cap(slice), len(slice));

	slice[0] = 'a'
	slice[1] = 'a'
	slice[2] = 'a'
	slice[3] = 'a'
	slice[4] = 'a'
	slice[5] = '0'

	fmt.Printf("%d %d  - %d %d\n",
		cap(array), len(array), cap(slice), len(slice));
}
