package eventmatcher_test

import (
	"fmt"

	"github.com/uudashr/coursehub/eventkit/eventmatcher"

	"github.com/uudashr/coursehub/eventkit"
)

func ExampleMatchOnly() {
	type OrderInitiated struct {
		ID string
	}

	type OrderCanceled struct {
		ID string
	}

	b := new(eventkit.Bus)

	h := eventkit.HandlerFunc(func(e eventkit.Event) {
		fmt.Printf("Got event %q", e.Name)
	})

	b.Subscribe(eventmatcher.MatchOnly(h, "OrderCanceled"))

	// this event will surpressed
	eventkit.Publish(b, OrderInitiated{
		ID: "some-id",
	})

	eventkit.Publish(b, OrderCanceled{
		ID: "some-id",
	})

	// Output: Got event "OrderCanceled"
}
