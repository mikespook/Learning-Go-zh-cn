package main

import "fmt"

func main() {
  s := "foobar"
  a := []rune(s) |\coderemark{Again a conversion}|
  for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
     a[i], a[j] = a[j], a[i] |\coderemark{Parallel assignment}|
  }
  fmt.Printf("%s\n", string(a)) |\coderemark{Convert it back}|
}
