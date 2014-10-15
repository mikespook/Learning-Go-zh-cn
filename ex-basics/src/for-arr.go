package main

import "fmt"

func main() {
   var arr [10]int |\coderemark{创建一个有 10 个元素的数组}|
   for i := 0; i < 10; i++ {
       arr[i] = i |\coderemark{一个一个填充它们}|
   }
   fmt.Printf("%v", arr) |\coderemark{利用 \%v Go 可以将其中的值打印出来}|
}
