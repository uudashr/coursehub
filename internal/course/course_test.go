package course_test

import (
	"fmt"
	"testing"

	"github.com/uudashr/coursehub/internal/course"
)

func TestCourse(t *testing.T) {
	cases := []struct {
		id    string
		name  string
		valid bool
	}{
		{id: "id-01", name: "Beginers Guide to Test", valid: true},
		{id: "id-02", name: "Beginers Guide to Integration test", valid: true},
		{id: "id-03", name: "", valid: false},
		{id: "", name: "Mutation Test", valid: false},
		{id: "", name: "", valid: false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			cs, err := course.New(c.id, c.name)
			if c.valid {
				if err != nil {
					t.Fatal(err)
				}

				if got, want := cs.ID(), c.id; got != want {
					t.Errorf("id got: %q, want: %q", got, want)
				}

				if got, want := cs.Name(), c.name; got != want {
					t.Errorf("name got: %q, want: %q", got, want)
				}
			} else {
				if err == nil {
					t.Error("expect error")
				}
			}

		})
	}
}
