package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing home handler by printing URL path %q", html.EscapeString(r.URL.Path))
}

func main() {
	PORT := ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/home", homeHandler)

	
	s := &http.Server{
		Addr:           PORT,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB = 2^20 bytes
	}

	log.Fatal(s.ListenAndServe())
}
