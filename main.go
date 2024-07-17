package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	server := http.Server{
		Handler: mux,
		Addr:    "localhost:8080",
	}

	fmt.Println("starting server...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error: ", err)
	}
}
