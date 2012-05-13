package main

import "os"

func main() {
	buf := make([]byte, 1024)
        f, _ := os.Open("/etc/passwd") |\longremark{打开文件，\func{os.Open} 返回一个实现了 \type{io.Reader} 和 \type{io.Writer} 的 \type{*os.File}；}|
        defer f.Close() |\longremark{确保关闭了 \var{f}；}|
	for {
                n, _ := f.Read(buf) |\longremark{一次读取 1024 字节；}|
                if n == 0 { break } |\longremark{到达文件末尾；}|
                os.Stdout.Write(buf[:n]) |\longremark{将内容写入 \type{os.Stdout}}|
	}
}
