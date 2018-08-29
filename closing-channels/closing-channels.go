package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            // 当close了某个channel后，从channel中获取数据时，第二个参数会返回false
            j, more := <- jobs

            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("receiverd all jobs")
                done <- true
                return
            }
        }
    } ()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("send job", j)
    }

    close(jobs)

    <-done
}