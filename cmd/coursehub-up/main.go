package main

import (
	"fmt"
	"log"
	nethttp "net/http"

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

	server := &nethttp.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	log.Printf("Listening on port %d\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Fail to serve:", err)
	}
}
