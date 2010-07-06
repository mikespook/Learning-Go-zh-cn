package main

type PersonAge struct { |\coderemark{First we define two structures as a new type, %
\texttt{PersonAge};}|
	name string
	age  int
}

type PersonShoe struct { |\coderemark{And \texttt{PersonShoe};}|
	name     string
	shoesize int
}

func WhichOne(x interface{}) {
	switch t := x.(type) {
	case *PersonAge:
		println("Age person")
	case *PersonShoe:
		println("Shoe person")
	}
}

func main() {
	p1 := new(PersonAge)
	p2 := new(PersonShoe)
	WhichOne(p1)
	WhichOne(p2)
}
