package applifecycle

import (
	"context"

	"github.com/uudashr/coursehub/eventkit"
)

// Lifecycle is the application service lifecycle.
// It dispatch events to the EventsHandler by providing eventkit.Publlisher via Context.
type Lifecycle struct {
	bus          *eventkit.Bus
	EventHandler eventkit.Handler
	events       []eventkit.Event
}

func (lc *Lifecycle) getBus() *eventkit.Bus {
	if lc.bus == nil {
		lc.bus = new(eventkit.Bus)
		lc.bus.SubscribeFunc(func(e eventkit.Event) {
			lc.events = append(lc.events, e)
		})
	}

	return lc.bus
}

// Context of the lifecycle.
func (lc *Lifecycle) Context(parent context.Context) context.Context {
	return eventkit.ContextWithPublisher(parent, lc.getBus())
}

// End the lifecycle.
func (lc *Lifecycle) End(err error) {
	handler := lc.EventHandler
	if err != nil || handler == nil {
		return
	}

	for _, e := range lc.events {
		handler.Handle(e)
	}
}
