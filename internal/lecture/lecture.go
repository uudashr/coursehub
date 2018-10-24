package lecture

import "errors"

// Lecture represents lecture.
type Lecture struct {
	id       string
	title    string
	videoURL string
	courseID string
}

// New constructs new Lecture.
func New(id, courseID, title, videoURL string) (*Lecture, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}

	if courseID == "" {
		return nil, errors.New("empty courseID")
	}

	if title == "" {
		return nil, errors.New("empty title")
	}

	if videoURL == "" {
		return nil, errors.New("empty videoURL")
	}

	return &Lecture{
		id:       id,
		courseID: courseID,
		title:    title,
		videoURL: videoURL,
	}, nil
}

// ID of the lecture.
func (l Lecture) ID() string {
	return l.id
}

// CourseID of the lecture.
func (l Lecture) CourseID() string {
	return l.courseID
}

// Title of the lecture.
func (l Lecture) Title() string {
	return l.title
}

// VideoURL of the lecture.
func (l Lecture) VideoURL() string {
	return l.videoURL
}
