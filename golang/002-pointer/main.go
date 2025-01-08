package main

import "fmt"

func changeValue(s int) {
    s = 0
}

func changeValueByPointer(ptr *int) {
    *ptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    changeValue(i)
    fmt.Println("first value:", i)

    changeValueByPointer(&i)
    fmt.Println("second value:", i)

    fmt.Println("pointer:", &i)
}