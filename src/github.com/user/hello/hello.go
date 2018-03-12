package main

import (
    "fmt"
    "github.com/user/stringutil"
    "github.com/user/data"
    _ "io"
    _ "strings"
    "io"
)

func qwe(i, j int) (int, int) {
    return i, j
}

func main() {
    var _ = io.Copy
    var s = stringutil.Reverse("Hello, world!\n")
    fmt.Printf(stringutil.Reverse("Hello, world!\n"))
    fmt.Printf(s)
    var fd float64
    _ = fd

    var i int64
    i = 10
    if i < 60 {
        fmt.Println(i)
    }

    for i = 0; i < 10; i++ {
        fmt.Println(i)
    }
    var a, b = qwe(1, 2)
    fmt.Println(a)
    fmt.Println(b)

    const str = "QWERWERQ"
    fmt.Println(str)

    figure := data.Data{
        Width:  1,
        Height: 2,
    }
    figure.IncreaseHeight(5)
    figure.IncreaseWidth(10)
    fmt.Println(figure.Square())
}
