package main

import (
    "fmt"
    "sync"
    "runtime"
)

type stPlayer struct {
    cnt int
}

func main() {
    runtime.GOMAXPROCS(2)

    var wg sync.WaitGroup
    var player stPlayer
    wg.Add(2)
    player.cnt = 0

    fmt.Println("Starting Go Routines")
    go func() {
        defer wg.Done()
        for value := 'a'; value < 'a' + 26; value++ {
            player.cnt++
            fmt.Printf("%c player.cnt=%d\n", value, player.cnt)
        }
    }()

    go func() {
        defer wg.Done()
        for number := 1; number < 27; number++ {
            player.cnt++
            fmt.Printf("%d player.cnt=%d\n", number, player.cnt)
        }
    }()

    fmt.Println("Waiting To Finish")
    wg.Wait()
    fmt.Println("Terminating Program")
}
