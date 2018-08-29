package main

import (
	"time"
	"fmt"
)

func worker(c chan bool) {
    fmt.Println("Working...");
    time.Sleep(time.Second)
    fmt.Println("Done");

    c <- true
}

func main() {
    done := make(chan bool)
    go worker(done)

    <- done
}
