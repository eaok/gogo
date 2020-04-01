package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// 请求URL部分信息
func urlHandler(w http.ResponseWriter, r *http.Request) {
	URL := r.URL

	fmt.Fprintf(w, "Scheme: %s\n", URL.Scheme)
	fmt.Fprintf(w, "Host: %s\n", URL.Host)
	fmt.Fprintf(w, "Path: %s\n", URL.Path)
	fmt.Fprintf(w, "RawPath: %s\n", URL.RawPath)
	fmt.Fprintf(w, "RawQuery: %s\n", URL.RawQuery)
	fmt.Fprintf(w, "Fragment: %s\n", URL.Fragment)
}

// 查看Header信息的函数
func headerHandler(w http.ResponseWriter, r *http.Request) {
	//获取所有首部
	for key, value := range r.Header {
		fmt.Fprintf(w, "%s: %v\n", key, value)
	}

	//获取某个特定首部
	fmt.Fprintln(w, r.Header["Accept-Encoding"])     //返回切片
	fmt.Fprintln(w, r.Header.Get("Accept-Encoding")) //返回字符串
}

// 查看请求的内容体
func bodyHandler(w http.ResponseWriter, r *http.Request) {
	data := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(data); err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	fmt.Fprintln(w, string(data))
}

// 返回HTML表单页面
func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
<html>
    <head>
        <title>Go Web</title>
    </head>
    
    <body>
        <form action="/formhandle?lang=cpp&name=ls" method="post" enctype="application/x-www-form-urlencoded">
            <label>Form:</label>
            <input type="text" name="lang" />
            <input type="text" name="age" />
            <button type="submit">提交</button>
        </form>
    </body>
</html>`)
}

// 当请求时form表单数据时，可以使用r.ParseForm解析请求数据
func formhandleHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)     //结果包含表单键值对和url键值对
	fmt.Fprintln(w, r.PostForm) //只获取表单键值对

	//下面的在需要是会自动调用ParseForm/ParseMultipartForm，并且只返回第一个
	fmt.Fprintln(w, r.FormValue("lang"))
	fmt.Fprintln(w, r.PostFormValue("lang"))
}

// 返回上传文件页面
func form2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
<html>
    <head>
        <title>Go Web</title>
    </head>

    <body>
		<form action="/multipartform?lang=cpp&name=dj" method="post" enctype="multipart/form-data">
			<label>MultipartForm:</label>
    		<input type="text" name="lang" />
    		<input type="text" name="age" />
    		<input type="file" name="uploaded" />
    		<button type="submit">提交</button>
		</form>
    </body>
</html>`)
}

//使用MultipartForm字段接收上传的文件
func multipartFormHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm) //MultipartForm字段只包含表单键值对

	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("Open failed: ", err)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ReadAll failed: ", err)
		return
	}

	ioutil.WriteFile("xx.md", data, os.ModePerm)
	fmt.Fprintln(w, string(data))
}

//使用FormFile方法接收上传的文件
func multipartForm2Handler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/url", urlHandler)
	mux.HandleFunc("/header", headerHandler)
	mux.HandleFunc("/body", bodyHandler)

	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/formhandle", formhandleHandler)

	mux.HandleFunc("/form2", form2Handler)
	mux.HandleFunc("/multipartform", multipartFormHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
