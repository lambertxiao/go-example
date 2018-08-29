# Golang语法

- [x] Hello World

    - import时包名要加上双引号
    - Println方法是首字母大写的(首字母大写的方法代表是public的)
    - 同一个目录下面只能有一个main函数

- [x] Values
- [x] Variables

    - var 变量名 类型 = 值
    - var 变量名 = 值 (变量类型GO会推断)
    - 变量名 := 值

- [ ] Constants
- [ ] For
- [ ] If/Else
- [ ] Switch
- [ ] Arrays
- [x] Slices

    - s := make([]string, 3) 声明了一个string类型的slice, 注意赋值符号用的是 `:=`
    - 变量第一次赋值是用的是 `:=`, 之后的赋值只可以用 `=`
    - append函数
    - copy函数
    - slice[startIndex:endIndex]

- [x] Maps

    - m := make(map[string]int) 声明一个key类型的string, value类型为int的map
    - n := map[string]int {"foo": 1, "bar": 2} 第二种声明方式
    - delete(mapName, "keyName")
    - _, prs := m["k2"] 通过prs可以判断k2这个key在不在map中
    - len(m) 可以获取到键值对的数量

- [x] Range

    - range 可用来循环数组的下标和值

        ```
        for i, num := range nums {
        }
        ```

    - range 可用来循环map的键值对

        ```
        for k, v := range kvs {
        }
        ```

    - range 可用来循环字符串

        ```
        for index, c := range "golang" {

        }
        ```

- [x] Functions

    - 函数返回值类型写在后面

        ```
        func plus(a int, b int) int {}
        ```

    - 当多个连续的参数为同样类型时，可以只声明最后一个参数类型

        ```
        func plusPlus(a, b, c int) int {}
        ```

    - 函数可以有多个返回值

        ```
        func vals() (int, int) {}
        ```

    - 函数支持可变长参数

        ```
        func sum(nums ...int) {}

        如果你有一个含有多个值的 slice，想把它们作为参数使用，你要这样调用 `func(slice...)`。
        ```

    - 函数的返回值类型支持函数,即闭包

        ```
        func intSeq() func() int {
            i := 0

            return func() int {
                i++
                return i++
            }
        }
        ```

    - 函数的参数类型支持指针

        ```
        func zeroptr(iptr *int) {
            *iptr = 0
        }

        i := 1
        zeroptr(&i)
        ```

- [x] Multiple Return Values
- [x] Variadic Functions
- [x] Closures
- [x] Recursion
- [x] Pointers
- [x] Structs
- [x] Methods

     - 可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接收的结构体

- [x] Interfaces

    - 当实现一个接口时,必须实现接口的所有方法
    - 可以基于接口编程

- [x] Errors

    - 引入 errors 包
    - 可以自定义error
    - 如果想在程序中使用一个自定义错误类型中的数据，需要通过类型断言来得到这个错误类型的实例。

- [x] Goroutines

    - 使用goroutines时需注意,当main函数执行完时,程序会直接退出的,不会等routines执行完

- [x] Channels

    - chan用于routines之间的通信
    - channel的用法

        ```
        // 声明一个channel
        message := make(chan string)
        // 将数据放入channel
        message <- "msg"
        // 从channel中取出数据
        msg := <-message

        // 声明一个带缓冲的channel
        message := make(chan string, 2)
        ```

- [x] Channel Buffering

    * 错误实例代码,当一个chan没有缓冲区时,往里面 **存取** 数据会直接造成deadlock

        ```
        func main() {
            message := make(chan string)

            message <- "msg1"
            fmt.Println(<-message)

            message <- "msg2"
            fmt.Println(<-message)
        }

        func main() {
            message := make(chan string)

            fmt.Println(<-message)
            message <- "msg1"

            message <- "msg2"
            fmt.Println(<-message)
        }
        ```

    * 解开deadlock的方法

        ```
        // 使用goroutines放数据
        func main() {
            message := make(chan string)

            go func() {
                message <- "msg1"
            } ()

            fmt.Println(<-message)

            go func() {
                message <- "msg2"
            } ()

            fmt.Println(<-message)
        }

        // 通过缓冲区存放数据
        func main() {
            message := make(chan string, 2)

            message <- "msg1"
            message <- "msg2"

            fmt.Println(<-message)
            fmt.Println(<-message)
        }
        ```

- [x] Channel Synchronization
- [x] Channel Directions

    - 只读的channel

        ```
        <-chan
        ```

    - 只写的channel

        ```
        chan <-
        ```

- [x] Select

    * 用来等待多个channel返回时,可以用select,当有一个channel返回了值,程序会继续向下运行

- [x] Timeouts

    * select + time.After 实现超时处理

- [x] Non-Blocking Channel Operations

    * 使用 select + default 实现非阻塞的channel通信

- [x] Closing Channels

    - 当close了一个channel后，从channel中获取数据时，返回值的第二个值代表channel是否已经关闭

- [x] Range over Channels

    - range 和 channel的close搭配使用

- [x] Timers
- [x] Tickers
- [x] Worker Pools
- [x] Rate Limiting

    - time.Tick 和 channel 实现速率控制

- [ ] Atomic Counters

    - 了解 `"sync/atomic"`
    - runtime.Gosched()

- [ ] Mutexes
- [ ] Stateful Goroutines
- [ ] Sorting
- [ ] Sorting by Functions
- [x] Panic
- [x] Defer
- [x] Collection Functions
- [x] String Functions
- [x] String Formatting
- [x] Regular Expressions
- [x] JSON
- [ ] Epoch
- [ ] Time Formatting / Parsing
- [ ] Random Numbers
- [ ] Number Parsing
- [x] URL Parsing
- [x] SHA1 Hashes
- [ ] Base64 Encoding
- [ ] Reading Files
- [ ] Writing Files
- [ ] Line Filters
- [ ] Command-Line Arguments
- [ ] Command-Line Flags
- [ ] Environment Variables
- [ ] Spawning Processes
- [ ] Exec'ing Processes
- [ ] Signals
- [ ] Exit