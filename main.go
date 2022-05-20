package main

import (
	"net/http"
	"time"

	"github.com/axzed/blog-service/internal/routers"
)

func main() {
	// create router
	router := routers.NewRouter()
	// define http.Server by ourself
	s := &http.Server{
		Addr:           ":8080", // TCP Endpoint
		Handler:        router,  // programme to process
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
