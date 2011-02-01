package main

func main() {
	a := func() {     |\coderemark{Define a nameless function and assign to \var{a}}|
	  println("Hello")
	}		  |\coderemark{No () here}|
	a()               |\coderemark{Call the function}|
}
