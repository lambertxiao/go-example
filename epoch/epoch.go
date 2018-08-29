package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now)

    secs := now.Unix()
    fmt.Println(secs)

    nanos := now.UnixNano()
    millis := nanos / 1000000
    fmt.Println(nanos)
    fmt.Println(millis)

    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}
