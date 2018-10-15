package main //<<1>>

import "fmt" //<<2>> // 引用格式化输出

/* Print something */ //<<3>>

func main() {
	文字 := "Hello, world!\n" +
		"καλημέρα κóσμε!" +
		"こんにちは 世界!" +
		"你好，世界！" //<<5>>
	fmt.Println(文字) //<<6>>
}
