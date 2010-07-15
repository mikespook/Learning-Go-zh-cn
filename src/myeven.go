package main

import (		|\coderemark{Import the following packages}|
	"./even"	|\coderemark{The local package even}|
	"fmt"		|\coderemark{The official fmt package}|
)

func main() {
	i := 5
	fmt.Printf("Is %d even? %v\n", i, even.Even(i))
}
