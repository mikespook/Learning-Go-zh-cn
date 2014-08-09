package main

import ( "os"; "bufio")

func main() {
	buf := make([]byte, 1024)
        f, _ := os.Open("/etc/passwd") |\longremark{打开文件；}|
	defer f.Close()
        r := bufio.NewReader(f) |\longremark{转换 \var{f} 为有缓冲的 \func{Reader}。%
\func{NewReader} 需要一个 \type{io.Reader}，因此或许你认为这会出错。%
但其实不会。\emph{任何}有 \func{Read()} 函数就实现了这个接口。%
同时，从列表\ref{src:read}可以看到，\type{*os.File} 已经这样做了；}|	
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
                n, _ := r.Read(buf)     |\longremark{从 Reader 读取，而向 Writer 写入，然后%
向屏幕输出文件。}|
		if n == 0 { break }
		w.Write(buf[0:n])
	}
}
