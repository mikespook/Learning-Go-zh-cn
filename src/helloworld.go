package main |\longremark{首行这个是必须的。所有的~Go 文件以~\mbox{\lstinline{package <something>}}
开头，对于独立运行的执行文件必须是~\lstinline{package main}；}|

import "fmt" |\longremark{这是说需要将~\package{fmt} 包加入~\package{main}。%
不是~\package{main} 的其他包都被称为库，其他许多编程语言有着类似的概念（参阅第~%
\ref{chap:packages} 章）。末尾以 \lstinline|//| 开头的内容是注释；}// 实现格式化的~I/O。

/* Print something */ |\longremark{这同样是注释，不过这是被包裹于 \lstinline|/*| 和 \lstinline|*/| 之间的；}|
func main() { |\longremark{%
\lstinline{package main} 必须首先出现，紧跟着是 \lstinline{import}。%
在 Go 中，\lstinline{package} 总是首先出现，然后是 \lstinline{import}，%
然后是其他所有内容。%
当 Go 程序在执行的时候，首先调用的函数是 \lstinline{main.main()}，这是从 C 中继承而来。%
这里定义了这个函数；}|
|\longremark{%
第 8 行调用了来自于 \package{fmt} 包的函数打印字符串到屏幕。%
字符串由 \lstinline{"} 包裹，并且可以包含非 ASCII 的字符。%
这里使用了希腊文和日文。}|
        fmt.Printf("Hello, world; or |$\kappa\alpha\lambda\eta\mu\acute{\epsilon}\rho\alpha\hspace{1em}\kappa$\'o$ \sigma\mu\epsilon$|; or|こんにちは 世界|")
}
