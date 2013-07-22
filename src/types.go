package main

func main() {
	var a int   |\coderemark{通用整数类型}|
	var b int32 |\coderemark{32 位整数类型}|
	a = 15
	b = a + a   |\coderemark{混合这些类型是非法的}|
	b = b + 5   |\coderemark{5 是一个（未定义类型的）常量，所以这没啥问题}|
}
