package main

import "os"

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd", os.O_RDONLY, 0666)
	defer f.Close()

        f, _ := os.Open("/etc/passwd") |\longremark{打开文件；}|
        defer f.Close() |\longremark{确保关闭（close）它；}|

	for {
                n, _ := f.Read(buf) |\longremark{一次读取 1024 字节；}|
                if n == 0 { break } |\longremark{到达文件末尾；}|
                os.Stdout.Write(buf[:n]) |\longremark{将内容写入 \var{Stdout}；}|
	}
}
