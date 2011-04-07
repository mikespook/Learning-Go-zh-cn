package main

import "fmt"

func main() {
        var arr [10]int |\coderemark{Create an array with 10 elements}|
	for i := 0; i < 10; i++ {
                arr[i] = i |\coderemark{Fill it one by one}|
	}
        fmt.Printf("%v", arr) |\coderemark{With \%v Go prints the type}|
}
