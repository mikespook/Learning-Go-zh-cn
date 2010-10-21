package main

import (		|\longremark{Import the following packages};|
	"./even"	|\longremark{The \emph{local} package \package{even};}|
	"fmt"		|\longremark{The official \package{fmt} package;}|
)

func main() {
	i := 5
	fmt.Printf("Is %d even? %v\n", i, even.Even(i)) |\longremark{Use the function from the %
\package{even} package. The syntax for accessing a function from a package is %
\lstinline{<package>.Function}.}|
}
