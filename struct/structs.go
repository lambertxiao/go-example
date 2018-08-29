package main

import "fmt"

type person struct {
    name string
    age int
}

func main() {
    p := person{name: "xsx", age: 24}
    fmt.Println("p = ", p)

    p2 := person{"xsx", 24}
    fmt.Println("p2 = ", p2)

    p3 := person{name: "p3man"}
    fmt.Println("p3 = ", p3)

    p3.age = 44
    fmt.Println("p3 = ", p3)

    p4 := &person{name: "haha", age: 52}
    fmt.Println("p4 = ", p4)

    p4.age = 25
    fmt.Println("p4 = ", p4)
}