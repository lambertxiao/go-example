package main

import (
    "fmt"
    "time"
)

func main() {
    // 假设有5个请求放置在channel中
    requests := make(chan int, 5)

    for i := 0; i < 5; i++ {
        requests <- i
    }

    close(requests)

    limiter := time.Tick(time.Second)

    // 这里要控制处理的速率
    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    burstyLimiter := make(chan time.Time, 3)

    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    go func() {
        for t := range time.Tick(time.Second) {
            burstyLimiter <- t
        }
    } ()

    burstyRequests := make(chan int, 5)

    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }

    close(burstyRequests)

    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}
