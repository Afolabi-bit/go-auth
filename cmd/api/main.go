package main

import (
	"go-auth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main() {
	r := httpserver.NewRouter()

	// standard Go type that runs a http server
	srv := &http.Server{
		Addr:              ":5000",
		Handler:           r,
		ReadHeaderTimeout: time.Second * 5,
	}
	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("server closed")
		} else {
			log.Fatalf("server failed to start: %v", err)
		}
	}
}
