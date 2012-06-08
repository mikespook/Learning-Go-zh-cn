package main

import "fmt"

func main() {
	const (
		FIZZ = 3 |\longremark{为了提高代码的可读性，定义两个常量。%
参阅``\titleref{sec:constants}''；}|
		BUZZ = 5
	)
	var p bool      |\longremark{判断是否需要打印内容；}|
	for i := 1; i < 100; i++ { |\longremark{for 循环，参阅``\titleref{sec:for}''}；|
		p = false
		if i%FIZZ == 0 { |\longremark{如果能被 FIZZ 整除，打印``Fizz''；}|
			fmt.Printf("Fizz")
			p = true
		}
		if i%BUZZ == 0 { |\longremark{如果能被 BUZZ 整除，打印``Buzz''。%
注意，FizzBuzz 的情况已经被处理了；}|
			fmt.Printf("Buzz")
			p = true
		}
		if !p {	|\longremark{如果 FIZZ 和 BUZZ 都没有打印，打印原始值；}|
			fmt.Printf("%v", i)
		}
                fmt.Println() |\longremark{换行。}|
        }
}
