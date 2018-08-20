package course

// Course represents course.
type Course struct {
	id   string
	name string
}

// New creates new Course.
func New(id, name string) (*Course, error) {
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
