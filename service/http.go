package service

import (
	"fmt"
	"log/slog"
	"net/http"
)

func ServeHttp() {
	slog.Info("Starting HTTP server")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Got request")
		fmt.Fprintln(w, "Hello")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Failed to service HTTP requests")
		slog.Error(err.Error())
	} else {
		slog.Info("HTTP server stopped without any errors")
	}

}
