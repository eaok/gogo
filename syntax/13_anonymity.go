package main

import (
    "fmt"
)

func getIntValue() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {
    nextInt := getIntValue()
    fmt.Printf("%d, %d, %d\n", nextInt(), nextInt(), nextInt())
}
