package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["k1"] = 1
    m["k2"] = 2

    fmt.Println("maps = ", m)
    fmt.Println("k1 = ", m["k1"])
    fmt.Println("len = ", len(m))

    delete(m, "k2")
    fmt.Println("maps = ", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int {"foo": 1, "bar": 2}
    fmt.Println("n = ", n)
}