package stack

// Stack holds the items.
type Stack struct {
	i    int
	data [10]int
}

// Push pushes an item on the stack.
func (s *Stack) Push(k int) {
	s.data[s.i] = k
	s.i++
}

// Pop pops an item from the stack.
func (s *Stack) Pop() (ret int) {
	s.i--
	ret = s.data[s.i]
        return
}
