/*
	1484210530.pdf
	page 68
	Basic http server
*/
package main

import "net/http"
import "time"

func messageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Go development"))
}

func main() {
	handler := http.NewServeMux()
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        handler,
	}
	println("Listening...")
	handler.HandleFunc("/welcome", messageHandler)
	server.ListenAndServe()
}
