package main

func main() {
	a := func() {     |\coderemark{定义一个匿名函数，并且赋值给~\var{a}}|
	  println("Hello")
	}		  |\coderemark{这里没有~()}|
	a()               |\coderemark{调用函数}|
}
