package errors

import "fmt"

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
// Be aware that high-level WithX couldn't be recognised by Is and errors.Is.
// Do not use this constructor without especial necessity.
func WithX(x interface{}) *Error {
	return &Error{reason: "", err: nil, x: x}
}

// IfErr returns WithErr if specified err is not nil. Otherwise it returns nil.
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
		var t *Error

		if As(e.err, &t) {
			err = t.WithErr(err)
		} else {
			err = &Error{reason: e.err.Error(), err: err, x: nil}
		}
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

	s := ""

	if e.reason != "" {
		s += e.reason
	}

	if e.err != nil {
		if len(s) > 0 {
			s += " because: "
		}

		s += e.err.Error()
	}

	if e.x != nil {
		if len(s) > 0 {
			s += " with data: "
		}

		s += fmt.Sprintf("%T{%+v}", e.x, e.x)
	}

	return s
}

// Unwrap implements errors/Unwrap interface.
func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.err
}

//nolint:gocognit
// Is implements errors/Is interface.
func (e *Error) Is(target error) bool {
	if e.Error() == target.Error() || e.reason == target.Error() {
		return true
	}

	var t *Error

	if !As(target, &t) {
		if e.err == nil {
			return target == nil
		}

		return Is(e.err, target)
	}

	return e.reason == t.reason && e.reason != "" ||
		Is(e.err, t) || Is(e, t.err)
}
