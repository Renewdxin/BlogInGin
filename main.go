package main

import (
	"BloginGin/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()

	// 自定义http.server
	s := &http.Server{
		Addr:           ":9090",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
