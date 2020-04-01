package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	Hobbies   []string `json:"hobbies"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the index page")
}

//将状态码写入响应中
func writeHeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, "This API not implemented!!!")
}

// 重定向到baidu.com
func moveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://baidu.com")
	w.WriteHeader(http.StatusFound)
}

//返回json数据
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u := &User{
		FirstName: "ls",
		LastName:  "ls",
		Age:       18,
		Hobbies:   []string{"reading", "learning"},
	}
	data, _ := json.Marshal(u)
	w.Write(data)
}

//接收json数据
func getjsonHandler(w http.ResponseWriter, r *http.Request) {
	ctValue := r.Header.Get("Content-Type")
	fmt.Println(ctValue) // 请求体的内容类型

	data := make([]byte, r.ContentLength) // 造切⽚，切⽚的⻓度是请求体内容的⻓度
	r.Body.Read(data)                     // 从请求体读数据

	// json反序列化
	u1 := new(User)
	json.Unmarshal(data, u1)
	fmt.Printf("%#v\n", u1)
	fmt.Fprint(w, "ok")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/writeHeader", writeHeaderHandler)
	mux.HandleFunc("/move", moveHandler)
	mux.HandleFunc("/json", jsonHandler)
	mux.HandleFunc("/getjson", getjsonHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
