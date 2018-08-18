package main

import (
	"log"

	"github.com/uudashr/coursehub/eventkit"
)

type logEventsHandler struct {
}

func (h *logEventsHandler) HandleEvents(events []eventkit.Event) {
	for _, e := range events {
		log.Printf("Event %+v\n", e)
	}
}
