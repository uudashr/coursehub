package course

import "errors"

// Course represents course.
type Course struct {
	id   string
	name string
}

// New creates new Course.
func New(id, name string) (*Course, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}

	if name == "" {
		return nil, errors.New("empty name")
	}

	return &Course{
		id:   id,
		name: name,
	}, nil
}

// ID of the course.
func (c Course) ID() string {
	return c.id
}

// Name of the course.
func (c Course) Name() string {
	return c.name
}
