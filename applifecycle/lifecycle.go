package applifecycle

import (
	"context"

	"github.com/uudashr/coursehub/eventkit"
)

// EventsHandler handles the events.
type EventsHandler interface {
	HandleEvents([]eventkit.Event)
}

// EventsHandlerFunc is the function adapter of EventsHandler.
type EventsHandlerFunc func([]eventkit.Event)

// HandleEvents invoke h(events).
func (h EventsHandlerFunc) HandleEvents(events []eventkit.Event) {
	h(events)
}

// Lifecycle is the application service lifecycle.
// It dispatch events to the EventsHandler by providing eventkit.Publlisher via Context.
type Lifecycle struct {
	bus     *eventkit.Bus
	Handler EventsHandler
	events  []eventkit.Event
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
	if err != nil || lc.Handler == nil {
		return
	}

	lc.Handler.HandleEvents(lc.events)
}
