package main

import "fmt"

func main() {
    s := make([]string, 3)

    // 像数组一样赋值
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"

    fmt.Println("set:", s)
    fmt.Println("get:", s[1])

    fmt.Println("len = ", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("s = ", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("c = ", c)

    l := s[2:5]
    fmt.Println("sl1 = ", l)

    l = s[:5]
    fmt.Println("sl2 = ", l)

    l = s[2:]
    fmt.Println("sl3 = ", l)

    twoD := make([][]int, 3)

    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)

        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }

    fmt.Println("2d = ", twoD)
}