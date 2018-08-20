package lecture

// Lecture represents lecture.
type Lecture struct {
	id       string
	title    string
	videoURL string
}

// New constructs new Lecture.
func New(id, title, videoURL string) (*Lecture, error) {
	return &Lecture{
		id:       id,
		title:    title,
		videoURL: videoURL,
	}, nil
}
