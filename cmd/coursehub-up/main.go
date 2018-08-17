package main

import (
	"context"
	"fmt"
	"log"
	nethttp "net/http"
	"time"

	"github.com/oklog/run"

	"github.com/uudashr/coursehub/internal/http"
	"github.com/uudashr/coursehub/internal/inmem"
)

const port = 8080

func main() {
	accRepo := inmem.NewAccountRepository()
	accSvc, err := http.NewAccountService(accRepo)
	if err != nil {
		log.Fatal("Fail to create account service:", err)
	}

	handler, err := http.NewHandler(accSvc)
	if err != nil {
		log.Fatal("Fail to create http handler:", err)
	}

	var g run.Group
	// Signal catcher
	{
		cancel := make(chan struct{})
		g.Add(func() error {
			return signalCatcher(cancel)
		}, func(error) {
			close(cancel)
		})
	}

	// HTTP Server
	{
		server := &nethttp.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: handler,
		}
		g.Add(func() error {
			log.Printf("Listening on port %d\n", port)
			if err := server.ListenAndServe(); err != nethttp.ErrServerClosed {
				return err
			}

			return err
		}, func(error) {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				log.Println("Fail to shutdown:", err)
			}
		})
	}

	if err := g.Run(); err != nil {
		log.Println("System error:", err)
	}

	log.Println("Stopped")
}
