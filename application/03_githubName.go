//查找github上还未注册的4位用户名
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
)

//MAXNUM 最大goroutine数量
const MAXNUM = 512

var result []string
var mutex sync.Mutex

//testURL 验证用户网址是否还未注册
func testURL(url string, wg *sync.WaitGroup, ch chan bool, fileName string) {
	defer func() {
		<-ch
		wg.Done()
		runtime.Goexit()
	}()

	fmt.Println(url, runtime.NumGoroutine(), len(result))

	// client := &http.Client{
	// 	Timeout: 1 * time.Second,
	// }
	// resp, err := client.Get(url)
	resp, err := http.Head(url)
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

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool, MAXNUM)
	str := "abcdefghijklmnopqrstuvwxyz"
	fileName := `./url.txt`

	_, err := os.Stat(fileName)
	if err == nil {
		os.Remove(fileName)
	}

	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			for k := 0; k < 26; k++ {
				for m := 0; m < 26; m++ {
					sub := string(str[i]) + string(str[j]) + string(str[k]) + string(str[m])
					subName := "https://github.com/" + sub
					wg.Add(1)
					ch <- true
					go testURL(subName, &wg, ch, fileName)
				}
			}
		}
	}

	wg.Wait()

	resultInFile(fileName)
}
