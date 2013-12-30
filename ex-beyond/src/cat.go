package main

|\longremark{包含所有需要用到的包；}|
import (
        "io"
	"os"
	"fmt"
	"bufio"
	"flag"
)

var numberFlag = flag.Bool("n", false, "number each line") |\longremark{定义新的开关 "n"，默认是关闭的。%
注意很容易写的帮助文本；}|

|\longremark{实际上读取并且显示文件内容的函数；}|
func cat(r *bufio.Reader) {
	i := 1
	for {
		buf, e := r.ReadBytes('\n')	|\longremark{每次读一行；}|
		if e == io.EOF {		|\longremark{如果到达文件结尾；}|
			break
		}
		if *numberFlag {		|\longremark{如果设定了行号，打印行号然后是内容本身；}|
			fmt.Fprintf(os.Stdout, "%5d  %s", i, buf)
			i++
		} else {			|\longremark{否则，仅仅打印该行内容。}|
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, e := os.Open(flag.Arg(i))
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n",
				os.Args[0], flag.Arg(i), e.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
