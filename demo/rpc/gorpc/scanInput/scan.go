package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func scan() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func scanWithTicker() {
	// 心跳发生器器
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for {
			fmt.Printf("heart beat.\n")
			<-ticker.C
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//scan()
	scanWithTicker()
}
