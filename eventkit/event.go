package eventkit

import (
	"context"
	"reflect"
	"sync"
	"time"
)

type contextKey int

const (
	keyPublisher contextKey = iota
)

// Event represents event.
type Event struct {
	Name        string
	Body        interface{}
	OccuredTime time.Time
}

// Publisher publish event.
type Publisher interface {
	Publish(Event)
}

// Handler handles event.
type Handler interface {
	Handle(Event)
}

// HandlerFunc is the func adapter for Han
type HandlerFunc func(Event)

// Handle executes h(e).
func (h HandlerFunc) Handle(e Event) {
	h(e)
}

// Bus represents the event bus.
type Bus struct {
	mu       sync.RWMutex
	handlers []Handler
}

// Publish the event.
func (b *Bus) Publish(e Event) {
	b.mu.RLock()
	for _, h := range b.handlers {
		h.Handle(e)
	}
	b.mu.RUnlock()
}

// Subscribe for events.
func (b *Bus) Subscribe(h Handler) {
	b.mu.Lock()
	b.handlers = append(b.handlers, h)
	b.mu.Unlock()
}

// SubscribeFunc subscribe for event using inline func.
func (b *Bus) SubscribeFunc(f func(e Event)) {
	b.Subscribe(HandlerFunc(f))
}

// PublisherFromContext returns publisher from context.
func PublisherFromContext(ctx context.Context) Publisher {
	pub, ok := ctx.Value(keyPublisher).(Publisher)
	if !ok {
		return nil
	}

	return pub
}

// ContextWithPublisher creates context with publisher inside.
func ContextWithPublisher(parent context.Context, pub Publisher) context.Context {
	return context.WithValue(parent, keyPublisher, pub)
}

// Publish the event body via Publisher.
func Publish(pub Publisher, body interface{}) {
	pub.Publish(Event{
		Name:        reflect.TypeOf(body).Name(),
		Body:        body,
		OccuredTime: time.Now(),
	})
}

// PublishContext the event body via Publisher on the context.
func PublishContext(ctx context.Context, body interface{}) {
	pub := PublisherFromContext(ctx)
	if pub == nil {
		return
	}

	pub.Publish(Event{
		Name:        reflect.TypeOf(body).Name(),
		Body:        body,
		OccuredTime: time.Now(),
	})
}
