package eventkit_test

import (
	"context"
	"fmt"
	"time"

	"github.com/uudashr/coursehub/eventkit"
)

func ExampleBus() {
	b := new(eventkit.Bus)

	var h eventkit.Handler
	// TODO: need to initialize h

	b.Subscribe(h)

	b.Publish(eventkit.Event{
		Name: "OrderInitiated",
		Body: map[string]interface{}{
			"ID": "some-id",
		},
		OccuredTime: time.Now(),
	})
}

func ExampleBus_simplify() {
	type OrderInitiated struct {
		ID string
	}

	b := new(eventkit.Bus)

	b.SubscribeFunc(func(e eventkit.Event) {
		fmt.Printf("Got event %q", e.Name)
	})

	eventkit.Publish(b, OrderInitiated{
		ID: "some-id",
	})

	// Output: Got event "OrderInitiated"
}

func ExamplePublishContext() {
	type OrderInitiated struct {
		ID string
	}

	b := new(eventkit.Bus)
	b.SubscribeFunc(func(e eventkit.Event) {
		fmt.Printf("Got event %q\n", e.Name)
	})

	ctx := eventkit.ContextWithPublisher(context.Background(), b)
	eventkit.PublishContext(ctx, OrderInitiated{ID: "some-id"})

	// Output: Got event "OrderInitiated"
}
