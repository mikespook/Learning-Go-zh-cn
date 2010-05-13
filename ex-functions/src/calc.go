package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var st = new(Stack)

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
	s.i--
	if s.i < 0 {
		s.i = 0
		return 0, false
	}
	ret = s.data[s.i]
	return ret, true
}

func (s *Stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" +
			strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

func main() {
	for {
		s, err := reader.ReadString('\n')
		var token string
		if err != nil {
			return
		}
		for _, c := range s {
			switch {
			case c >= '0' && c <= '9':
				token = token + string(c)
			case c == ' ':
				x, _ := (strconv.Atoi(token))
				st.push(x)
				token = ""
			case c == '+':
				var p, _ = st.pop()
				var q, _ = st.pop()
				fmt.Printf("%d\n", p + q)
			case c == '*':
				var p, _ = st.pop()
				var q, _ = st.pop()
				fmt.Printf("%d\n", p * q)
			case c == '-':
				var p, _ = st.pop()
				var q, _ = st.pop()
				fmt.Printf("%d\n", q - p)
			case c == 'q':
				return
			default:
				//error
			}
		}
	}
}
