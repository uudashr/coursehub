package lecture

// Lecture represents lecture.
type Lecture struct {
	title    string
	videoURL string
}

// New constructs new Lecture.
func New(title, videoURL string) (*Lecture, error) {
	return &Lecture{
		title:    title,
		videoURL: videoURL,
	}, nil
}
