package stack

import (
	"strconv"
)

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) Push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *Stack) Pop() (ret int) {
	s.i--
	ret = s.data[s.i]
}

func (s *Stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":"
			+ strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
