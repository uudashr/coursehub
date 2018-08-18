package main

import (
	"log"

	"github.com/uudashr/coursehub/eventkit"
)

type logEventHandler struct {
}

func (h *logEventHandler) Handle(e eventkit.Event) {
	log.Printf("Event %+v\n", e)
}
