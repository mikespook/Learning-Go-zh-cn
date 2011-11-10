package main

import "fmt"

type NameAge struct {
	name string	|\coderemark{不导出}|
	age  int	|\coderemark{不导出}|
}

func main() {
	a := new(NameAge)
	a.name = "Pete"; a.age = 42

	fmt.Printf("%v\n", a)
}
