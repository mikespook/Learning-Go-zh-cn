package main

import "fmt"

func fibonacci(value int) []int {
	x := make([]int, value) |\longremark{创建一个用于保存函数执行结果的~\key{array}；}|
	a, b := 0, 1 |\longremark{开始计算斐波那契数列；}|
	for n := 0; n < value; n++ {
		a, b = b, a + b |\longremark{$x_n = x_{n+1}$；}| |\longremark{$x_{n+1} = x_n + x_{n+1}$；}|
		x[n] = a
	}
	return x |\longremark{返回\emph{整个}~array；}|
}

func main() {
	for _, term := range fibonacci(10) { |\longremark{使用关键字~\key{range} 可以``遍历''数字得到斐波那契函数返回的序列。这里有~10 个，且打印了出来。}|
		fmt.Printf("%v ", term)
	}
}
