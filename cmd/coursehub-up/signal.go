package main

import (
	"errors"
	"os"
	"os/signal"
	"syscall"
)

func signalCatcher(cancel <-chan struct{}) error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	select {
	case <-stop:
		return nil
	case <-cancel:
		return errors.New("catch signal canceled")
	}
}
