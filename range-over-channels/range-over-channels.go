package main

import "fmt"

func main() {
    queue := make(chan string, 2)
    queue <- "1"
    queue <- "2"

    close(queue)

    // 当queue未close时，range将继续阻塞，等待下一个值
    for element := range queue {
        fmt.Println("element = ", element)
    }
}
