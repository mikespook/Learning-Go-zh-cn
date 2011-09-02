package stack

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
        return
}
