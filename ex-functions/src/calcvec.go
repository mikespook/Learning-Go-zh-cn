package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"container/vector"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var st = new(vector.IntVector)

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
				r, _ := strconv.Atoi(token)
				st.Push(r)
				token = ""
			case c == '+':
				fmt.Printf("%d\n", st.Pop()+st.Pop())
			case c == '*':
				fmt.Printf("%d\n", st.Pop()*st.Pop())
			case c == '-':
				p := st.Pop()
				q := st.Pop()
				fmt.Printf("%d\n", q-p)
			case c == 'q':
				return
			default:
				//error
			}
		}
	}
}
