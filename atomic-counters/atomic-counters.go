package main

import (
    "fmt"
    "time"
    "sync/atomic"
    "runtime"
)

func main() {
    var ops uint64 = 0

    for i := 0; i < 1100000; i++ {
        go func() {
            atomic.AddUint64(&ops, 1)

            // 允许其他携程执行
            runtime.Gosched()
        } ()
    }

    time.Sleep(time.Millisecond)

    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops:", opsFinal)
}
