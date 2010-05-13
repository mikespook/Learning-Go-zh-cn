package stack

import (
	"strconv"
)

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) push(k int) (ok bool) {
	if s.i+1 > 9 {
		return false
	}
	s.data[s.i] = k
	s.i++
	return true
}

func (s *Stack) pop() (ret int, ok bool) {
	if s.i-1 < 0 {
		return 0, false
	}
	ret = s.data[s.i]
	s.i--
	return ret, true
}

func (s *Stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":"
			+ strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
