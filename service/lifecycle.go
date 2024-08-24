package service

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitForTerminate() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
