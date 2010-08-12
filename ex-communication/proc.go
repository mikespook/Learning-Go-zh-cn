package main

import (
	"os"
	"fmt"
	"sort"
	"bufio"
	"strings"
	"strconv"
	"container/vector"
	"exp/iterable"
)

const (
	PID = iota
	PPID
	COMM
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

	pid, _ := os.ForkExec("/bin/ps", []string{"ps", "-e", "-opid,ppid,comm"}, nil, "",
		[]*os.File{nil, pw, nil})
	defer os.Wait(pid, os.WNOHANG)
	pw.Close()

	child := make(map[int]*vector.IntVector)	// strings that are numbers (pid)
	r.ReadString('\n')		// header
	for {
		s, ok := r.ReadString('\n')
		if ok != nil {
			break
		}
		f := strings.Fields(s)

		if _, present := child[atoi(f[PPID])]; !present {
			v := new(vector.IntVector)
			child[atoi(f[PPID])] = v
		}
		child[atoi(f[PPID])].Push(atoi(f[PID]))
	}

	schild := make([]int, len(child))
	i := 0
	for k, _ := range child {
		schild[i] = k
		i++
	}
	sort.SortInts(schild)

	for _, ppid := range schild {
		pids := child[ppid]
		fmt.Printf("Pid %d has %d child",
			ppid, pids.Len())

		if pids.Len() == 1 {
			fmt.Print(": ")
		} else {
			fmt.Print("eren: ")
		}
		do := func(e int) {
			fmt.Printf(" %d", e)
		}
		pids.Do(do)
		fmt.Println()
	}
}
