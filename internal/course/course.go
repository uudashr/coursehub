package course

// Course represents course.
type Course struct {
	name string
}

// NewCourse creates new Course.
func NewCourse(name string) (*Course, error) {
	return &Course{
		name: name,
	}, nil
}

// Name of the course.
func (c Course) Name() string {
	return c.name
}
