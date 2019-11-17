package wlerr

import (
	"fmt"
)

// WLError is an error containing that wrapps ancestor errors.
type WLError struct {
	Errs []error
}

func (wle *WLError) Error() string {
	s := ""
	for _, e := range wle.Errs {
		if s != "" {
			s += ": "
		}
		s += e.Error()
	}

	return s
}

// Errorf produces a WLError with a formatted string.
func Errorf(format string, args ...interface{}) error {
	return &WLError{
		Errs: []error{fmt.Errorf(format, args...)},
	}
}

// ToWLError converts an error to a WLError.
func ToWLError(err error) *WLError {
	if err == nil {
		return nil
	}

	wle, ok := err.(*WLError)
	if !ok {
		wle = &WLError{
			Errs: []error{err},
		}
	}

	return wle
}

// Wrapf wraps an error in a WLError.
func Wrapf(err error, format string, args ...interface{}) error {
	wle := ToWLError(err)
	if wle == nil {
		wle = &WLError{}
	}
	wle.Errs = append(wle.Errs, fmt.Errorf(format, args...))

	return wle
}
