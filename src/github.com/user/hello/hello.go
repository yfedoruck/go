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

    var arr [5]int
    var ii int
    var j int
    j = 4
    fmt.Println(arr, ii, j)
    //arr[0] = 1
    //arr2 := [5]int{1, 2, 3}
    //fmt.Println(arr2, arr)

    var qq [4]int
    ww := [4]int{}
    fmt.Println(qq, ww)

    primes := [5]int{1, 2, 3, 5, 7}
    fmt.Println(primes[4])

    sl := primes[1:3]
    fmt.Println(sl)



    t := make(map[string]int)
    t["a"] = 3
    t["b"] = 5
    t["c"] = 7777777

    t["c"] = 8888
    //delete(t, "c")

    fmt.Println(t)
    fmt.Println(
        t["b"],
        t["c"],
        t["CCCSdf"],
        t["d  v"],
    )

    el, ok := t["aa"]
    fmt.Println(el, ok)





}
