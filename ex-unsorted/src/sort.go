// orignally from:
// http://www.thebitsource.com/programming-software-development/google-go-programming-introduction/
// heavily modified and influenced by the Go sort package
// Why can't you use a type switch to detect the type
// of the array and fire off the correct sort function?

package main

import "fmt"

// Generic Interface to any sortable type 
type S interface {
	Len() int           // len() as a method
	Less(i, j int) bool // <     as a method
	Swap(i, j int)      // p[i], p[j] = p[j], p[i] as a method
}
// For each type we want to sort we have to implement the above interface

// this is what needs to be done, but you cannot redefine non-local
// and thus built-in types.
//func (p []float) Len() int               { return len(p) }
//func (p []float) Less(i int, j int) bool { return p[i] < p[j] }
//func (p []float) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

// So define a type for an array of floats
// Needed because you can not define methods
// on built-in types.
type Xf []float

func (p Xf) Len() int               { return len(p) }
func (p Xf) Less(i int, j int) bool { return p[i] < p[j] }
func (p Xf) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

// same for strings
type Xs []string

func (p Xs) Len() int               { return len(p) }
func (p Xs) Less(i int, j int) bool { return p[i] < p[j] }
func (p Xs) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

// The generic sort function become sorter, because
// of all the work we did above
func Sort(S Interface) {
	for i := 1; i < S.Len(); i++ {
		for j := i; j > 0 && S.Less(j, j-1); j-- {
			S.Swap(j, j-1)
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
