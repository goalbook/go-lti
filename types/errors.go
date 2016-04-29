package types

import (
	"errors"
	"fmt"
)

var (
	errInvalidContextType = errors.New("Bad Context Type")

	errUnsupportedLTIVersion = func(s ltiVersion) error { return errors.New(fmt.Sprintf("Unsupported version %s.", s)) }
	errUnsupportedLTIMessage = func(s ltiMessage) error { return errors.New(fmt.Sprintf("Unsupported message type %s.", s)) }
)
