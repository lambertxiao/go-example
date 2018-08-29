package main

import "fmt"

func f(text string) {
    for _, c := range text {
        fmt.Printf("%c\n", c)
    }
}

func main() {
    f("1233456789")

    go f("abcdefghi")

    go func(text string) {
        for _, c := range text {
            fmt.Println(c)
        }
    }(",./;'p[]")

    fmt.Scanln()
    fmt.Println("Done")
}