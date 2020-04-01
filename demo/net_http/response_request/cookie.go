package main

import (
	"fmt"
	"log"
	"net/http"
)

//使用ResponseWriter.Header().Set/Add设置cookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := &http.Cookie{
		Name:     "name",
		Value:    "lianshi",
		HttpOnly: true,
	}
	c2 := &http.Cookie{
		Name:     "age",
		Value:    "18",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
}

//使用http.SetCookie⽅法设置cookie
func setCookie2(w http.ResponseWriter, r *http.Request) {
	c1 := &http.Cookie{
		Name:     "name",
		Value:    "lianshishi",
		HttpOnly: true,
	}
	c2 := &http.Cookie{
		Name:     "age",
		Value:    "28",
		HttpOnly: true,
	}
	http.SetCookie(w, c1)
	http.SetCookie(w, c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Host:", r.Host)
	fmt.Fprintln(w, "Cookies:", r.Header["Cookie"])
}

//使用http.Request提供了⼀些⽅法可以更容易地获取cookie：
func getCookie2(w http.ResponseWriter, r *http.Request) {
	name, err := r.Cookie("name")
	if err != nil {
		fmt.Fprintln(w, "cannot get cookie of name")
	}

	cookies := r.Cookies()
	fmt.Fprintln(w, name)
	fmt.Fprintln(w, cookies)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/set_cookie", setCookie)
	mux.HandleFunc("/set_cookie2", setCookie2)
	mux.HandleFunc("/get_cookie", getCookie)
	mux.HandleFunc("/get_cookie2", getCookie2)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
