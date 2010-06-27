ssuming slices of ints, you can do (not really tested):

func concat(old1, old2 []int) []int {
   newslice := make([]int, len(old1) + len(old2))
   copy(newslice, old1)
   copy(newslice[len(old1):], old2)
   return newslice
}

A fun little exercise might be to write

func concat(slices ...[]int) []int

That is, a function to concatenate efficiently an arbitrary number of slices,
as opposed to just two.

TODO, but nice exercise
