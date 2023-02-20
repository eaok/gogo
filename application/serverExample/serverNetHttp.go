package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello1")
	})

	fmt.Println("server1 listen ...")
	server.ListenAndServe()
}
