package main

import (
	"fmt"
	"sort"
	"exec"
	"strings"
	"strconv"
	"container/vector"
)

const (
	PID = iota
	PPID
)

func main() {
	ps := exec.Command("ps", "-e", "-opid,ppid,comm")
	output, _ := ps.Output()
	child := make(map[int]*vector.IntVector)
	for i, s := range strings.Split(string(output), "\n") {
		if i == 0 || len(s) == 0 { // Kill first line and last
			continue
		}
		f := strings.Fields(s)
		fpp, _ := strconv.Atoi(f[PPID])
		fp, _ := strconv.Atoi(f[PID])
		if _, present := child[fpp]; !present {
			v := new(vector.IntVector)
			child[fpp] = v
		}
		// Save the child PIDs on a vector
		child[fpp].Push(fp)
	}

	// Sort the PPIDs
	schild := make([]int, len(child))
	i := 0
	for k, _ := range child {
		schild[i] = k
		i++
	}
	sort.Ints(schild)
	for _, ppid := range schild { // Walk through the sorted list
		fmt.Printf("Pid %d has %d child", ppid, child[ppid].Len())
		if child[ppid].Len() == 1 {
			fmt.Printf(": %v\n", []int(*child[ppid]))
		} else {
			fmt.Printf("ren: %v\n", []int(*child[ppid]))
		}
	}
}
