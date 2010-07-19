package main

func main() {
	var a int	|\coderemark{Generic integer type}|
	var b int32	|\coderemark{32 bits integer type}|
	a = 15
	b = a + a	|\coderemark{Illegal mixing of these types}|
}
