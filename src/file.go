package main

import "os"

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd", os.O_RDONLY, 0666)
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 { break }
		os.Stdout.Write(buf[0:n])
	}
}
