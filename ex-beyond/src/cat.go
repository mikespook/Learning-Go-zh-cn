package main

|\longremark{Include all the packages we need;}|
import (
	"os"
	"fmt"
	"bufio"
	"flag"
)

var numberFlag = flag.Bool("n", false, "number each line") |\longremark{Define a %
new flag "n", which defaults to off. Note that we get the help for free;}|

|\longremark{Start the function that actually reads the file's contents and displays it;}|
func cat(r *bufio.Reader) {
	i := 1
	for {
		buf, e := r.ReadBytes('\n')	|\longremark{Read one line at the time;}|
		if e == os.EOF {		|\longremark{Or stop if we hit the end;}|
			break
		}
		if *numberFlag {		|\longremark{If we should number %
each line, print the line number and then the line itself;}|
			fmt.Fprintf(os.Stdout, "%5d  %s", i, buf)
			i++
		} else {			|\longremark{Otherwise we could just print the line.}|
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
		f, e := os.Open(flag.Arg(i), os.O_RDONLY, 0)
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n",
				os.Args[0], flag.Arg(i), e.String())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
