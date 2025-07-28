package main

import (
	"fmt"
	"go-web/chat"
	"log"
	"net/http"
)

// add by ai
func enableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// add by ai
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w, r)
		if r.Method == "OPTIONS" {
			return
		}
		next(w, r)
	}
}

func health(rw http.ResponseWriter, req *http.Request) {

	log.Printf("request to /hello")

	fmt.Fprintf(rw, "hello, world!")
}
func main() {
	http.HandleFunc("/hello", corsMiddleware(health))
	http.HandleFunc("/chat", corsMiddleware(chat.Chat))
	http.ListenAndServe(":8080", nil)
}
