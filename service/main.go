package main

import (
	"fmt"
	"net/http"
)

func health(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("start")
	fmt.Fprintf(rw, "hello, world!")
}
func main() {
	http.HandleFunc("/hello", health)
	http.ListenAndServe(":8080",nil)
}
