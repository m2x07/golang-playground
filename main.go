package main

import (
    "time"
    "fmt"
)

func main()  {
    msg := "hello world!"

    for _, v := range msg {
        fmt.Printf("%c", v)
        time.Sleep(time.Millisecond * 100)
    }
    time.Sleep(time.Millisecond * 500)
    fmt.Println("\n:)")
}
