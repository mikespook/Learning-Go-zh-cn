package main
import ( "os"; "bufio")

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd") |\longremark{Open the file, \key{os.Open} returns %
a \type{*os.File} which implements the \type{io.Reader} interface;}|
	defer f.Close()
	r := bufio.NewReader(f) |\longremark{Turn \var{f} into a buffered \func{Reader}. %
\func{NewReader} expects an \type{io.Reader}, so you %
might think this will fail. But it does not. %
\emph{Anything} that has such a \func{Read()} function implements this %
interface. And from listing \ref{src:read} we can see %
that \type{*os.File} indeed does so;}|
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		n, _ := r.Read(buf)     |\longremark{Read from the Reader and write to the Writer, and thus %
print the file to the screen.}|
		if n == 0 { break }
		w.Write(buf[0:n])
	}
}

