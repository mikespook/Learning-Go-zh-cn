package main

import (
	"fmt"
	"strconv"
)

// stack impl
// simple array with pointer

type stack struct {
	i    int
	data [10]int
}


func (s stack) push(k int) (ok bool) {
	if s.i+1 > 9 {
		return false
	}
	fmt.Printf("hier\n")
	s.data[s.i] = k
	s.i++
	return true
}

func (s stack) pop() (ret int, ok bool) {
	if s.i-1 < 0 {
		return 0, false
	}
	ret = s.data[s.i]
	s.i--
	return ret, true
}

func (s stack) String() string {
	var str string
	for i := 0; i <= s.i; i++ {
		str = str + "[" +
			strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}


func main() {
	var s stack

	s.push(25)
	fmt.Printf("stack %d\n", s.i);
	s.push(14)
	fmt.Printf("stack %d\n", s.i);

	fmt.Printf("stack %v\n", s)

}
