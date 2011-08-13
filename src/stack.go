// This package implementes a simple stack
package stack

import (
	"strconv"
)

// Stack hold 10 integer element and has an index
type Stack struct {
	i    int
	data [10]int
}

// Push an element on the stack and increment the index
func (s *Stack) Push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

// Pop an element from the stack and decrement the index
func (s *Stack) Pop() (ret int) {
	if s.i-1 < 0 {
		return 0
	}
	ret = s.data[s.i]
	s.i--
	return ret
}

// String returns a string representation of the stack
func (s *Stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" +
			strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
