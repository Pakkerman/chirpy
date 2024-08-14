package main

import (
	"fmt"
	"net/http"
)

type apiConfig struct {
	fileserverHits int
}

func main() {
	mux := http.NewServeMux()
	apiCfg := apiConfig{}

	fsHandler := apiCfg.middlewareMatricsInc(http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mux.Handle("/app/*", fsHandler)

	mux.HandleFunc("GET /api/healthz", handlerHealthz)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("GET /api/reset", apiCfg.handlerReset)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	fmt.Println("starting server...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
