package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	var chars, words, lines int
	r := bufio.NewReader(os.Stdin) |\longremark{Start a new reader that reads from standard input;}|
	for {
		switch s, ok := r.ReadString('\n'); true { |\longremark{Read a line from the input;}|
		case ok != nil: |\longremark{If we received an error, we assume it was because %
of a EOF. So we print the current values;}|
			fmt.Printf("%d %d %d\n", chars, words, lines);
			return
		default:        |\longremark{Otherwise we count the charaters, words and increment the %
lines.}|
			chars += len(s)
			words += len(strings.Fields(s))
			lines++
		}
	}
}
