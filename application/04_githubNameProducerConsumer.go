//用生产者消费者模式实现查找github上还未注册的4位用户名
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
)

//result 结果切片
var result []string
var mutex sync.Mutex

//resultInFile 把结果写进文件
func resultInFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for k, value := range result {
		file.WriteString(value)
		file.WriteString("\t")
		if (k+1)%10 == 0 {
			file.WriteString("\n")
		}
	}
	fmt.Println(result)
}

//producer 生产者
func producer(wg *sync.WaitGroup, products chan<- string) {
	defer wg.Done()
	str := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			for k := 0; k < 26; k++ {
				for m := 0; m < 26; m++ {
					sub := string(str[i]) + string(str[j]) + string(str[k]) + string(str[m])
					product := "https://github.com/" + sub
					products <- product
				}
			}
		}
	}
}

//consumer 消费者
func consumer(wg *sync.WaitGroup, products <-chan string, name int, fileName string) {
	defer wg.Done()

	for product := range products {
		fmt.Printf("%d\t%s\t%d\n", name, product, runtime.NumGoroutine())
		resp, err := http.Get(product)
		if err != nil {
			log.Println(err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == 404 {
			mutex.Lock()
			result = append(result, resp.Request.URL.Path)
			if len(result) >= 100 {
				resultInFile(fileName)
				result = []string{}
			}
			mutex.Unlock()
		}
	}
}

func main() {
	var wgp, wgc sync.WaitGroup
	products := make(chan string, 1024)
	fileName := `./url.txt`

	_, err := os.Stat(fileName)
	if err == nil {
		os.Remove(fileName)
	}

	wgp.Add(1)
	go producer(&wgp, products) //生产者

	//创建消费者
	for i := 0; i < 512; i++ {
		wgc.Add(1)
		go consumer(&wgc, products, i, fileName)
	}

	wgp.Wait()
	close(products)
	wgc.Wait()

	resultInFile(fileName)
}
