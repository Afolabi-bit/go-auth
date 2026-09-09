package main

import (
	"context"
	"go-auth/internal/app"
	"go-auth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Failed to start app: %v", err)
	}
	defer func() {
		if err := app.Close(ctx); err != nil {
			log.Printf("Failed to close app: %v", err)
		}
	}()

	r := httpserver.NewRouter(app)

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
