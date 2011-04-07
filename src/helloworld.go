package main |\longremark{This first line is just required. All Go files start with %
\lstinline{package <something>}, \lstinline{package main} is required for a %
standalone executable;}|

import "fmt"  // Implements formatted I/O. \longremark{This says we need \package{"fmt"} in %
addition to \package{main}. A package other than \package{main} is commonly called a %
library, a familiar concept of many programming languages (see chapter \ref{chap:packages}). %
The line ends with a comment which is started with \lstinline|//|;}

/* Print something */ |\longremark{This is also a comment, but this one is enclosed in %
\lstinline|/*| and \lstinline|*/|;}|
func main() { |\longremark{%
Just as \lstinline{package main} was required to be first, %
\lstinline{import} may come next. In Go, \lstinline{package} is always first, then %
\lstinline{import}, then everything else. %
When your Go program is executed, the first function called will be %
\lstinline{main.main()}, which mimics the behavior from C. Here we declare %
that function;}|
|\longremark{%
On line 8 we call a function from the package \package{fmt} to print a %
string to the screen. The string is enclosed with \lstinline{"} and may %
contain non-ASCII characters. Here we use Greek and Japanese.}|
 fmt.Printf("Hello, world; or |$\kappa\alpha\lambda\eta\mu\acute{\epsilon}\rho\alpha\hspace{1em}\kappa$\'o$ \sigma\mu\epsilon$|; or|\begin{cjk}こんにちは 世界\end{cjk}|\n")
}

