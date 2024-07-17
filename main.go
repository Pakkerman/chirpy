package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/app/*", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/healthz", healthz)

	server := &http.Server{
		Handler: mux,
		Addr:    "localhost:8080",
	}

	fmt.Println("starting server...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
