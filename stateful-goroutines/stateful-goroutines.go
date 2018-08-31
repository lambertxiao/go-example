package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

type readOp struct {
    key int
    resp chan int
}

type writeOp struct {
    key int
    val int
    resp chan bool
}

func main()  {
    var readOps uint64 = 0
    var writeOps uint64 = 0

    // 请求读的队列
    reads := make(chan *readOp)
    // 请求写的队列
    writes := make(chan *writeOp)

    go func () {
        // state为这个携程独占,所有需要读/写的操作都需要发请求到这个携程中
        var state = make(map[int]int)

        for {
            select {
            case read := <- reads:
                // 读state,需要读哪个key,读到的结果发到read里的channel中
                read.resp <- state[read.key]
            case write := <- writes:
                // 写state,需要写哪个key,写入的值是什么,写完返回个true
                state[write.key] = write.val
                write.resp <- true
            }
        }
    } ()

    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := &readOp{
                    key: rand.Intn(5),
                    resp: make(chan int),
                }

                reads <- read
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := &writeOp{
                    key: rand.Intn(5),
                    val: rand.Intn(100),
                    resp: make(chan bool),
                }

                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // 主线程在1s后查看两个操作数的值
    time.Sleep(time.Second)
    fmt.Println("readOps = ", atomic.LoadUint64(&readOps))
    fmt.Println("writeOps = ", atomic.LoadUint64(&writeOps))
}
