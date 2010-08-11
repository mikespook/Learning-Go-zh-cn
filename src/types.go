package main

func main() {
	var a int	|\coderemark{Generic integer type}|
	var b int32	|\coderemark{32 bits integer type}|
	a = 15
	b = a + a	|\coderemark{Illegal mixing of these types}|
	b = b + 5	|\coderemark{5 is a (typeless) constant, so this is OK}|
}
