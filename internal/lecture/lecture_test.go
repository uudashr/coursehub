package lecture_test

import (
	"fmt"
	"testing"

	"github.com/uudashr/coursehub/internal/lecture"
)

func TestLecture(t *testing.T) {
	cases := []struct {
		id       string
		courseID string
		title    string
		videoURL string
		valid    bool
	}{
		{id: "lc-01", courseID: "cs-01", title: "Introduction", videoURL: "http://coursehub.com/lc-01/cs-01.mp4", valid: true},
		{id: "lc-02", courseID: "cs-01", title: "Why unit test?", videoURL: "http://coursehub.com/lc-01/cs-02.mp4", valid: true},
		{id: "", courseID: "cs-01", title: "Introduction", videoURL: "http://coursehub.com/lc-01/cs-01.mp4", valid: false},
		{id: "lc-01", courseID: "", title: "Introduction", videoURL: "http://coursehub.com/lc-01/cs-01.mp4", valid: false},
		{id: "lc-01", courseID: "cs-01", title: "", videoURL: "http://coursehub.com/lc-01/cs-01.mp4", valid: false},
		{id: "lc-01", courseID: "cs-01", title: "Introduction", videoURL: "", valid: false},
	}
	_ = cases

	for i, c := range cases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			lc, err := lecture.New(c.id, c.courseID, c.title, c.videoURL)
			if c.valid {
				if err != nil {
					t.Fatal(err)
				}

				if got, want := lc.ID(), c.id; got != want {
					t.Errorf("id got: %q, want: %q", got, want)
				}

				if got, want := lc.CourseID(), c.courseID; got != want {
					t.Errorf("courseID got: %q, want: %q", got, want)
				}

				if got, want := lc.Title(), c.title; got != want {
					t.Errorf("title got: %q, want: %q", got, want)
				}

				if got, want := lc.VideoURL(), c.videoURL; got != want {
					t.Errorf("courseID got: %q, want: %q", got, want)
				}
			} else {
				if err == nil {
					t.Error("expect error")
				}
			}
		})
	}
}
