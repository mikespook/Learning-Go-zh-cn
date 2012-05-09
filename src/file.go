package main
import "os"

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd") |\longremark{Open the file, \func{os.Open} returns a \type{*os.File}, which implements \type{io.Reader} and \type{io.Writer};}|
	defer f.Close() |\longremark{Make sure we close \var{f} again;}|
	for {
		n, _ := f.Read(buf) |\longremark{Read up to 1024 bytes at the time;}|
		if n == 0 { break } |\longremark{We have reached the end of the file;}|
		os.Stdout.Write(buf[:n]) |\longremark{Write the contents to \type{os.Stdout}}|
	}
}
