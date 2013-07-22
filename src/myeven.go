package main

import (|\longremark{导入下面的包；}|
	"even"	|\longremark{\emph{本地}包 \package{even} 在这里导入；}|
	"fmt"	|\longremark{官方 \package{fmt} 包导入；}|
)

func main() {
	i := 5
	fmt.Printf("Is %d even? %v\n", i, even.Even(i)) |\longremark{调用 \package{even} 包中的函数。%
访问一个包中的函数的语法是 \lstinline{<package>.Function()}。}|
}
