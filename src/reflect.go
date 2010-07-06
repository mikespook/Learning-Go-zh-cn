package main

type PersonAge struct { |\longremark{First we define two structures as a new type, %
\texttt{PersonAge};}|
	name string
	age  int
}

type PersonShoe struct { |\longremark{And \texttt{PersonShoe};}|
	name     string
	shoesize int
}

func main() {
	p1 := new(PersonAge)
	p2 := new(PersonShoe)
	WhichOne(p1)
	WhichOne(p2)
}

func WhichOne(x interface{}) { |\longremark{This function must accept \emph{both} %
types as valid input, so we use the empty interface, which every type implements;}|
	switch t := x.(type) { |\longremark{The type switch;}|
	case *PersonAge:	|\longremark{When allocated with \func{new} its a %
pointer;}|
		println("Age person")
	case *PersonShoe:
		println("Shoe person")
	}
}

