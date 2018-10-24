package core

import (
	"errors"
)

// ErrDuplicate returns if there is duplicate entries on repository.
var ErrDuplicate = errors.New("Duplicate")
