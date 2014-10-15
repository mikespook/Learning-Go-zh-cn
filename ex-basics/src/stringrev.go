package main

import "fmt"

func main() {
  s := "foobar"
  a := []rune(s) |\coderemark{转换}|
  for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
     a[i], a[j] = a[j], a[i] |\coderemark{并列赋值}|
  }
  fmt.Printf("%s\n", string(a)) |\coderemark{转换回字符串}|
}
