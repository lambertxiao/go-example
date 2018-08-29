package main

import "fmt"

func main() {
    nums := []int{2, 3, 4}
    sum := 0

    for _, num := range nums {
        sum += num
        fmt.Println(num)
    }

    fmt.Println("sum = ", sum)

    for index, num := range nums {
        fmt.Printf("index = %d, num = %d", index, num)
    }

    kvs := map[string]int {"foo": 1, "bar": 2}

    for k, v := range kvs {
        fmt.Printf("%s -> %d\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key = ", k)
    }

    for index, c := range "golang" {
        fmt.Printf("index = %d, c = %c", index, c)
    }
}