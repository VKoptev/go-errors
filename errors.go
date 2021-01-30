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

// WithReason returns new Error with set error reason.
func WithReason(code string) *Error {
	return &Error{reason: code, err: nil, x: nil}
}

// WithErr returns new Error with set error.
func WithErr(err error) *Error {
	return &Error{reason: "", err: err, x: nil}
}

// WithX returns new Error with set error data.
func WithX(x interface{}) *Error {
	return &Error{reason: "", err: nil, x: x}
}

// IfErr returns new Error if specified err is not nil. Otherwise it returns nil.
func IfErr(err error) *Error {
	if err == nil {
		return nil
	}

	return WithErr(err)
}

// WithReason sets error reason and returns self.
func (e *Error) WithReason(reason string) *Error {
	if e == nil {
		return WithReason(reason)
	}

	if e.reason != "" {
		return &Error{reason: reason, err: e, x: nil}
	}

	return &Error{reason: reason, err: e.err, x: e.x}
}

// WithErr sets error and returns self.
func (e *Error) WithErr(err error) *Error {
	if e == nil {
		return WithErr(err)
	}

	if e.err != nil {
		return &Error{reason: e.Error(), err: err, x: nil}
	}

	return &Error{reason: e.reason, err: err, x: e.x}
}

// WithX sets error data and returns self.
func (e *Error) WithX(x interface{}) *Error {
	if e == nil {
		return WithX(x)
	}

	if e.x != nil {
		return &Error{reason: e.Error(), err: nil, x: x}
	}

	return &Error{reason: e.reason, err: e.err, x: x}
}

// IfErr sets error and returns self if specified err is not nil. Otherwise it returns nil.
func (e *Error) IfErr(err error) *Error {
	if err == nil {
		return nil
	}

	return e.WithErr(err)
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
