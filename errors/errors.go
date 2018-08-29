package main

import (
    "errors"
    "fmt"
)

func f1(arg int) (int, error) {
    if arg == 0 {
        return -1, errors.New("can't work with 0")
    }

    return 10 / arg, nil
}

type customerError struct {
    arg int
    prob string
}

func (e *customerError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 0 {
        return -1, &customerError{arg, "can't work with it"}
    }

    return arg + 3, nil
}

func f3() (int, int, error, int) {
    return 1, 1, errors.New("hahah"), 1
}

func main() {
    if result, e := f1(0); e != nil {
        fmt.Println(e)
    } else {
        fmt.Println(result)
    }

    result, e := f1(2)

    if e != nil {
        fmt.Println(e)
    } else {
        fmt.Println(result)
    }

    if result, e := f2(0); e != nil {
        fmt.Println(e)
    } else {
        fmt.Println(result)
    }

    v1, v2, e, _ := f3()

    fmt.Println(v1)
    fmt.Println(v2)
    fmt.Println(e)

    _, e2 := f2(0)

    if ae, ok := e2.(*customerError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
