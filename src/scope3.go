package main

var a int

func main() {
        a = 5
        println(a)
        f()
}

func f() {
        a := 6
        println(a)
        g()
}

func g() {
        println(a)
}
