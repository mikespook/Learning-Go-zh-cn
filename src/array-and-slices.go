package main

func main() {
var array [100]int   |\coderemark{Create array, index from 0 to 99}|
slice := array[0:99] |\coderemark{Create slice, index from 0 to 98}|

slice[98] = 'a'	     |\coderemark{OK}|
slice[99] = 'a'      |\coderemark{Error: "throw: index out of range"}|
}
