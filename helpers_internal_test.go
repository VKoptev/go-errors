package errors

import (
	"errors"
	"testing"
)

var (
	err1 = errors.New("error")
	err2 = errors.New("error")
	err3 = WithErr(err1)
	err4 = errors.New("another error")
)

type test struct {
	nam string
	err error
	tgt []error
	exp bool
}

//nolint:dupl
func TestIs(t *testing.T) {
	t.Parallel()

	tt := []test{
		{nam: "(err1,err1)", err: err1, tgt: []error{err1}, exp: true},
		{nam: "(err1,err2)", err: err1, tgt: []error{err2}, exp: false},
		{nam: "(err1,err3)", err: err1, tgt: []error{err3}, exp: false},
		{nam: "(err1,err1,err2)", err: err1, tgt: []error{err1, err2}, exp: true},
		{nam: "(err1,err2,err1)", err: err1, tgt: []error{err2, err1}, exp: false},
		{nam: "(err3,err1)", err: err3, tgt: []error{err1}, exp: true},
		{nam: "(err3,err2)", err: err3, tgt: []error{err2}, exp: true},
		{nam: "(err3,err4)", err: err3, tgt: []error{err4}, exp: false},
		{nam: "(err3,err1,err4)", err: err3, tgt: []error{err1, err4}, exp: true},
		{nam: "(err3,err4,err1)", err: err3, tgt: []error{err4, err1}, exp: false},
	}

	for _, tc := range tt {
		tc := tc

		t.Run(tc.nam, func(t *testing.T) {
			t.Parallel()

			if Is(tc.err, tc.tgt...) != tc.exp {
				t.Errorf("actual=%v expected=%v", Is(tc.err, tc.tgt...), tc.exp)
			}
		})
	}
}

//nolint:dupl
func TestOneOf(t *testing.T) {
	t.Parallel()

	tt := []test{
		{nam: "(err1,err1)", err: err1, tgt: []error{err1}, exp: true},
		{nam: "(err1,err2)", err: err1, tgt: []error{err2}, exp: false},
		{nam: "(err1,err3)", err: err1, tgt: []error{err3}, exp: false},
		{nam: "(err1,err1,err2)", err: err1, tgt: []error{err1, err2}, exp: true},
		{nam: "(err1,err2,err1)", err: err1, tgt: []error{err2, err1}, exp: true},
		{nam: "(err3,err1)", err: err3, tgt: []error{err1}, exp: true},
		{nam: "(err3,err2)", err: err3, tgt: []error{err2}, exp: true},
		{nam: "(err3,err4)", err: err3, tgt: []error{err4}, exp: false},
		{nam: "(err3,err1,err4)", err: err3, tgt: []error{err1, err4}, exp: true},
		{nam: "(err3,err4,err1)", err: err3, tgt: []error{err4, err1}, exp: true},
	}

	for _, tc := range tt {
		tc := tc

		t.Run(tc.nam, func(t *testing.T) {
			t.Parallel()

			if OneOf(tc.err, tc.tgt...) != tc.exp {
				t.Errorf("actual=%v expected=%v", OneOf(tc.err, tc.tgt...), tc.exp)
			}
		})
	}
}

//nolint:dupl
func TestEachOf(t *testing.T) {
	t.Parallel()

	tt := []test{
		{nam: "(err1,err1)", err: err1, tgt: []error{err1}, exp: true},
		{nam: "(err1,err2)", err: err1, tgt: []error{err2}, exp: false},
		{nam: "(err1,err3)", err: err1, tgt: []error{err3}, exp: false},
		{nam: "(err1,err1,err2)", err: err1, tgt: []error{err1, err2}, exp: false},
		{nam: "(err1,err2,err1)", err: err1, tgt: []error{err2, err1}, exp: false},
		{nam: "(err3,err1)", err: err3, tgt: []error{err1}, exp: true},
		{nam: "(err3,err2)", err: err3, tgt: []error{err2}, exp: true},
		{nam: "(err3,err4)", err: err3, tgt: []error{err4}, exp: false},
		{nam: "(err3,err1,err4)", err: err3, tgt: []error{err1, err4}, exp: false},
		{nam: "(err3,err4,err1)", err: err3, tgt: []error{err4, err1}, exp: false},
	}

	for _, tc := range tt {
		tc := tc

		t.Run(tc.nam, func(t *testing.T) {
			t.Parallel()

			if EachOf(tc.err, tc.tgt...) != tc.exp {
				t.Errorf("actual=%v expected=%v", EachOf(tc.err, tc.tgt...), tc.exp)
			}
		})
	}
}

type dummyErr struct{}

func (d *dummyErr) Error() string {
	return ""
}

func (d *dummyErr) As(error) bool {
	return false
}

func TestAs(t *testing.T) {
	t.Parallel()

	var (
		err *Error
		dum *dummyErr
	)

	tt := []struct {
		nam string
		err error
		tgt interface{}
		exp bool
	}{
		{nam: "(err1,*Error)", err: err1, tgt: &err, exp: false},
		{nam: "(err3,*Error)", err: err3, tgt: &err, exp: true},
		{nam: "(err1,*dummyErr)", err: err1, tgt: &dum, exp: false},
		{nam: "(err3,*dummyErr)", err: err3, tgt: &dum, exp: false},
	}

	for _, tc := range tt {
		tc := tc

		t.Run(tc.nam, func(t *testing.T) {
			t.Parallel()

			if As(tc.err, tc.tgt) != tc.exp {
				t.Errorf("actual=%v expected=%v", As(tc.err, tc.tgt), tc.exp)
			}
		})
	}
}
