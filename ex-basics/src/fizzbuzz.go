package main

import "fmt"

func main() {
	const (
		FIZZ = 3 |\longremark{Define two constants to make the code more %
readable. See section "\titleref{sec:constants}";}|
		BUZZ = 5
	)
	var p bool      |\longremark{Holds if we already printed someting;}|
	for i := 1; i < 100; i++ { |\longremark{for-loop, see section "\titleref{sec:for}"};|
		p = false
		if i%FIZZ == 0 { |\longremark{If divisible by FIZZ, print "Fizz";}|
			fmt.Printf("Fizz")
			p = true
		}
		if i%BUZZ == 0 { |\longremark{And if divisble by BUZZ, print %
"Buzz". Note that we have also taken care of the FizzBuzz case;}|
			fmt.Printf("Buzz")
			p = true
		}
		if !p {	|\longremark{If neither FIZZ nor BUZZ printed, print the value;}|
			fmt.Printf("%v", i)
		}
		fmt.Println() |\longremark{Format each output on a new line.}|
	}
}
