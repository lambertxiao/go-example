package main

import (
	"fmt"
)

// func main() {
//     message := make(chan string)

//     go func() {
//         message <- "ping"
//     } ()

//     msg := <-message
//     fmt.Println(msg)
// }

// func main() {
//     message := make(chan string)

//     go func() {
//         for {
//             time.Sleep(10)
//             message <- "->"
//         }
//     } ()

//     go func() {
//         for {
//             msg := <-message
//             fmt.Println(msg)
//         }
//     } ()

//     fmt.Scanln()
// }

// func main() {
//     message := make(chan string)

//     go func() {
//         message <- "msg1"
//     } ()

//     fmt.Println(<-message)

//     go func() {
//         message <- "msg2"
//     } ()

//     fmt.Println(<-message)
// }

func main() {
    message := make(chan string, 2)

    message <- "msg1"
    message <- "msg2"

    fmt.Println(<-message)
    fmt.Println(<-message)
}