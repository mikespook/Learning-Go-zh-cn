package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	var chars, words, lines int
	r := bufio.NewReader(os.Stdin)
	for {
		switch s, ok := r.ReadString('\n'); true {
		case ok != nil:
			fmt.Printf("%d %d %d\n", chars, words, lines);
			return
		default:
			chars += len(s)
			words += len(strings.Fields(s))
			lines++
		}
	}
}
