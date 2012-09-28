package main

import (
	"fmt"
)

const (
	size = 8
	norm = size - 1
)

var B = [size][size]string{
	{".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", "."}}

var leftdiag [size*2 - 1]bool
var rightdiag [size*2 - 1]bool
var col [size]bool

func main() {
	tryQueen(0)
}

func tryQueen(x int) {
	for y := 0; y < size; y++ {
		if col[y] == false && leftdiag[x-y+norm] == false && rightdiag[x+y] == false {
			B[y][x] = "Q"	// reverse coordinates so it prints correctly
			col[y] = true
			leftdiag[x-y+norm] = true
			rightdiag[x+y] = true
			if x < size-1 {
				tryQueen(x + 1)
			} else {
				fmt.Printf("Solution found\n%s", func() string {
					s := ""
					for i := 0; i < 8; i++ {
						for j := 0; j < 8; j++ {
							s += B[i][j]
						}
						s += "\n"
					}
					return s
				}())
				return
			}
			B[y][x] = "."
			col[y] = false
			leftdiag[x-y+norm] = false
			rightdiag[x+y] = false
		}
	}
}
