package main

import "os"
import "strings"
import "fmt"

func main() {
    os.Setenv("foo", "1")
    fmt.Println("GOROOT:", os.Getenv("GOROOT"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
        fmt.Println(pair[0])
    }
}