// orignally from:
// http://www.thebitsource.com/programming-software-development/google-go-programming-introduction/
// heavily modified and influenced by the Go sort package

package main

import (
	"fmt"
	"rand"
	"time"
)

// Generic Interface to any sortable type 
type Sorter interface {
	Len() int		// len() as a method
	Less(i, j int) bool	// <     as a method
	Swap(i, j int)		// p[i], p[j] = p[j], p[i] as a method
}
// For each type we want to sort we have to implement the above interface


// define a type for an array of floats
// Needed because you can not define methods
// on built-in types.
type FloatArray []float

func (p FloatArray) Len() int               { return len(p) }
func (p FloatArray) Less(i int, j int) bool { return p[i] < p[j] }
func (p FloatArray) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

// same for strings
type StringArray []string

func (p StringArray) Len() int               { return len(p) }
func (p StringArray) Less(i int, j int) bool { return p[i] < p[j] }
func (p StringArray) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

// The generic sort function become sorter, because
// of all the work we did above
func Sort(Sorter Interface) {
	for i := 1; i < Sorter.Len(); i++ {
		for j := i; j > 0 && Sorter.Less(j, j-1); j-- {
			Sorter.Swap(j, j-1)
		}
	}
}

func main() {
	numbers := FloatArray{44.5, 67.3, 3.9, 17.8, 89.0, 10.2, 76.3, 9.4, 14.7, 89.1}
	strings := StringArray{"kobe", "odom", "farmar", "fisher", "artest", "bynum", "brown", "gasol"}

	Sort(numbers)
	fmt.Printf("%v\n", numbers)

	Sort(strings)
	fmt.Printf("%v\n", strings)
}
