package main

import "os"

func main() {
    // panic("a problem")

    if _, err := os.Create("/root/file"); err != nil {
        panic(err)
    }
}
