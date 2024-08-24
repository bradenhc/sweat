package main

import (
	"log/slog"

	"hitchcock.codes/sweat/service"
)

func main() {
	slog.Info("Instance is starting")

	go service.ServeHttp()

	service.WaitForTerminate()

	slog.Info("Instance is shutting down")
}
