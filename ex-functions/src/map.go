package main

import "fmt"

/* define the empty interface as a type */
type e interface{}

func mult2(f e) e {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string) + f.(string) + f.(string)
	}
	return f
}

func Map(n []e, f func(e) e) {
	for k, v := range n {
		n[k] = f(v)
	}
}

func main() {
	m := []e{1, 2, 3, 4}
	s := []e{"a", "b", "c", "d"}
	Map(m, mult2)
	Map(s, mult2)
	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", s)
}
