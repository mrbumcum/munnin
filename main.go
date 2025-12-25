package main

import (
	"flag"
	"net/http"
	"time"
)

var (
	port string
	originServer string
)

func main() {
	flag.StringVar(&port, "port", "8080", "TCP connection the server listens on")
	flag.StringVar(&originServer, "originServer", "origin", "URL of the origin server")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRequest)

	srv := &http.Server{
		Addr: ":"+port,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	srv.ListenAndServe()

}

func handleRequest(w http.ResponseWriter, r *http.Request) {

} 
