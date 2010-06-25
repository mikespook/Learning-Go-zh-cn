package main

func main() {
	var a int |\coderemark{generic integer type}|
	var b int32 |\coderemark{32 bits integer type}|
	a = 15
	b = a + a |\coderemark{illegal mixing of these types}|
}
