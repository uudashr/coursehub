package eventmatcher

import (
	"github.com/uudashr/coursehub/eventkit"
)

// Matcher matches the event.
type Matcher interface {
	Match(eventkit.Event) bool
}

// MatcherFunc is functio adapter of Matcher.
type MatcherFunc func(eventkit.Event) bool

// Match invoke m(e).
func (m MatcherFunc) Match(e eventkit.Event) bool {
	return m(e)
}

// Handler implements eventkit.Handler which only pass the event based on Matcher.
type Handler struct {
	next eventkit.Handler
	m    Matcher
}

// NewHandler constructs new Handler.
func NewHandler(next eventkit.Handler, m Matcher) (*Handler, error) {
	return &Handler{
		next: next,
		m:    m,
	}, nil
}

// MustNewHandler constructs new Handler which panic if error found.
func MustNewHandler(next eventkit.Handler, m Matcher) *Handler {
	h, err := NewHandler(next, m)
	if err != nil {
		panic(err)
	}

	return h
}

// Handle implements eventkit.Handler.
func (h Handler) Handle(e eventkit.Event) {
	if !h.m.Match(e) {
		return
	}

	h.next.Handle(e)
}

// MatchOnly pass only event based on it names.
func MatchOnly(next eventkit.Handler, eventNames ...string) *Handler {
	return MustNewHandler(next, MatcherFunc(func(e eventkit.Event) bool {
		for _, name := range eventNames {
			if e.Name == name {
				return true
			}
		}

		return false
	}))
}
