package errors

import (
	"errors"
	"fmt"
	"strings"
)

// Error is custom error.
type Error struct {
	reason string
	err    error
	x      interface{}
}

// WithReason returns new error with set error reason.
func WithReason(code string) *Error {
	return &Error{reason: code, err: nil, x: nil}
}

// WithErr returns new error with set error.
func WithErr(err error) *Error {
	return &Error{reason: "", err: err, x: nil}
}

// WithX returns new error with set error data.
func WithX(x interface{}) *Error {
	return &Error{reason: "", err: nil, x: x}
}

// WithReason sets error reason and returns self.
func (e *Error) WithReason(reason string) *Error {
	e.reason = reason

	return e
}

// WithErr sets error and returns self.
func (e *Error) WithErr(err error) *Error {
	var t *Error

	if errors.As(err, &t) {
		return e.chainError(t)
	}

	return e.chain(err)
}

// WithX sets error data and returns self.
func (e *Error) WithX(x interface{}) *Error {
	e.x = x

	return e
}

// Error implements builtin/error interface.
func (e *Error) Error() string {
	if e == nil {
		return "nil"
	}

	var buf strings.Builder

	if e.reason != "" {
		buf.WriteString(e.reason)
	}

	if e.err != nil {
		if buf.Len() > 0 {
			buf.WriteString(" because: ")
		}

		buf.WriteString(e.err.Error())
	}

	if e.x != nil {
		if buf.Len() > 0 {
			buf.WriteString(" with data: ")
		}

		buf.WriteString(fmt.Sprintf("%T{%+v}", e.x, e.x))
	}

	return buf.String()
}

// Unwrap implements errors/Unwrap interface.
func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.err
}

// Is implements errors/Is interface.
func (e *Error) Is(target error) bool {
	var t *Error

	if errors.As(target, &t) {
		return t.reason == e.reason
	}

	return errors.Is(e.err, target)
}

func (e *Error) chain(err error) *Error {
	if e.err != nil {
		err = WithReason(err.Error()).WithErr(e.err)
	}

	e.err = err

	return e
}

func (e *Error) chainError(err *Error) *Error {
	if e.reason != err.reason {
		return e.chain(err)
	}

	if e.x == nil && err.x != nil {
		e.x = err.x
	}

	return e.chain(err.err)
}
