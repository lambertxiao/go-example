package main

import (
    "time"
    "fmt"
)

func main() {
    ticker := time.NewTicker(time.Second)

    go func() {
        for t := range ticker.C {
            fmt.Println(t)
        }
    } ()

    time.Sleep(10 * time.Second)
    ticker.Stop()

    fmt.Println("Ticker stopped")
}
