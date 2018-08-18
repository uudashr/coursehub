package applifecycle_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/uudashr/coursehub/eventkit"

	"github.com/uudashr/coursehub/applifecycle"
)

func ExampleLifecycle() {
	// --- Event declaration ---
	type OrderInitiated struct {
		ID     string
		ItemID string
	}

	// --- Function declaration ---
	InitiateOrder := func(ctx context.Context, orderID string, itemID string) error {
		// ...

		eventkit.PublishContext(ctx, OrderInitiated{
			ID:     orderID,
			ItemID: itemID,
		})
		return nil
	}

	// --- Lifecycle usage ---
	h := applifecycle.EventsHandlerFunc(func(events []eventkit.Event) {
		for _, e := range events {
			fmt.Printf("Got event %q\n", e.Name)
		}
	})

	lc := &applifecycle.Lifecycle{ // begin: lifecycle
		Handler: h,
	}

	err := InitiateOrder(lc.Context(context.Background()), "order-98", "item-52")

	lc.End(err) // end: lifecycle

	// Output: Got event "OrderInitiated"
}

func TestLifecycle_minimum(t *testing.T) {
	lc := &applifecycle.Lifecycle{}
	lc.End(nil)
}
