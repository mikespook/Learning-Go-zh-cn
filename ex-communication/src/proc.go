package main

import (
	"os"
	"fmt"
	"sort"
	"bufio"
	"strings"
	"strconv"
	"container/vector"
)

const (
	PID = iota
	PPID
)

func atoi(s string) (x int) {
	x, _ = strconv.Atoi(s)
	return
}

func main() {
	pr, pw, _ := os.Pipe()
	defer pr.Close()
	r := bufio.NewReader(pr)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	pid, _ := os.StartProcess("/bin/ps", []string{"ps", "-e", "-opid,ppid,comm"}, nil, "", []*os.File{nil, pw, nil})
	defer os.Wait(pid, os.WNOHANG)
	pw.Close()

	child := make(map[int]*vector.IntVector)
	s, ok := r.ReadString('\n') // Discard the header line
	s, ok = r.ReadString('\n')
	for ok == nil {
		f := strings.Fields(s)
		if _, present := child[atoi(f[PPID])]; !present {
			v := new(vector.IntVector)
			child[atoi(f[PPID])] = v
		}
		// Save the child PIDs on a vector
		child[atoi(f[PPID])].Push(atoi(f[PID]))
		s, ok = r.ReadString('\n')
	}

	// Sort the PPIDs
	schild := make([]int, len(child))
	i := 0
	for k, _ := range child {
		schild[i] = k
		i++
	}
	sort.SortInts(schild)
	// Walk throught the sorted list
	for _, ppid := range schild {
		fmt.Printf("Pid %d has %d child", ppid, child[ppid].Len())
		if child[ppid].Len() == 1 {
			fmt.Printf(": %v\n", []int(*child[ppid]))
		} else {
			fmt.Printf("ren: %v\n", []int(*child[ppid]))
		}
	}
}
