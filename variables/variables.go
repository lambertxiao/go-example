package main

import "fmt"

func main() {
    var a = "initial"

    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go会自动推断变量类型
    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    // var f string = "short" 的精简写法
    f := "short"
    fmt.Println(f)
}
