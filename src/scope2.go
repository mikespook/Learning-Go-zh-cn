package main

var a = 6

func main() {
	p()
	q()
	p()
}

func p() {
	println(a)
}

func q() {
	a = 5
	println(a)
}
