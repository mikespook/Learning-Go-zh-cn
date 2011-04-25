package main
import "fmt"

/* 定义一个空的 interface 类型 */
type e interface{}

func mult2(f e) e {
    switch f.(type) {
    case int:
        return f.(int) * 2
    case string:
        return f.(string) + f.(string) + f.(string) + f.(string)
    }
    return f
}

func Map(n []e, f func(e) e) []e {
    m := make([]e, len(n))
    for k, v := range n {
        m[k] = f(v)
    }
    return m
}

func main() {
    m := []e{1, 2, 3, 4}
    s := []e{"a", "b", "c", "d"}
    mf := Map(m, mult2)
    sf := Map(s, mult2)
    fmt.Printf("%v\n", mf)
    fmt.Printf("%v\n", sf)
}
