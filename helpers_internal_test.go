package errors

import (
	"errors"
	"testing"
)

var (
	err1 = errors.New("error")
	err2 = errors.New("error")
	err3 = WithErr(err1)
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

	testTable(t, Is, []test{
		{nam: "(err1,err1)", err: err1, tgt: []error{err1}, exp: true},
		{nam: "(err1,err2)", err: err1, tgt: []error{err2}, exp: false},
		{nam: "(err1,err3)", err: err1, tgt: []error{err3}, exp: false},
		{nam: "(err1,err1,err2)", err: err1, tgt: []error{err1, err2}, exp: true},
		{nam: "(err1,err2,err1)", err: err1, tgt: []error{err2, err1}, exp: false},
		{nam: "(err3,err1)", err: err3, tgt: []error{err1}, exp: true},
		{nam: "(err3,err2)", err: err3, tgt: []error{err2}, exp: false},
		{nam: "(err3,err1,err2)", err: err3, tgt: []error{err1, err2}, exp: true},
		{nam: "(err3,err2,err1)", err: err3, tgt: []error{err2, err1}, exp: false},
	})
}

//nolint:dupl
func TestOneOf(t *testing.T) {
	t.Parallel()

	testTable(t, OneOf, []test{
		{nam: "(err1,err1)", err: err1, tgt: []error{err1}, exp: true},
		{nam: "(err1,err2)", err: err1, tgt: []error{err2}, exp: false},
		{nam: "(err1,err3)", err: err1, tgt: []error{err3}, exp: false},
		{nam: "(err1,err1,err2)", err: err1, tgt: []error{err1, err2}, exp: true},
		{nam: "(err1,err2,err1)", err: err1, tgt: []error{err2, err1}, exp: true},
		{nam: "(err3,err1)", err: err3, tgt: []error{err1}, exp: true},
		{nam: "(err3,err2)", err: err3, tgt: []error{err2}, exp: false},
		{nam: "(err3,err1,err2)", err: err3, tgt: []error{err1, err2}, exp: true},
		{nam: "(err3,err2,err1)", err: err3, tgt: []error{err2, err1}, exp: true},
	})
}

//nolint:dupl
func TestEachOf(t *testing.T) {
	t.Parallel()

	testTable(t, EachOf, []test{
		{nam: "(err1,err1)", err: err1, tgt: []error{err1}, exp: true},
		{nam: "(err1,err2)", err: err1, tgt: []error{err2}, exp: false},
		{nam: "(err1,err3)", err: err1, tgt: []error{err3}, exp: false},
		{nam: "(err1,err1,err2)", err: err1, tgt: []error{err1, err2}, exp: false},
		{nam: "(err1,err2,err1)", err: err1, tgt: []error{err2, err1}, exp: false},
		{nam: "(err3,err1)", err: err3, tgt: []error{err1}, exp: true},
		{nam: "(err3,err2)", err: err3, tgt: []error{err2}, exp: false},
		{nam: "(err3,err1,err2)", err: err3, tgt: []error{err1, err2}, exp: false},
		{nam: "(err3,err2,err1)", err: err3, tgt: []error{err2, err1}, exp: false},
	})
}

func testTable(t *testing.T, f func(error, ...error) bool, tt []test) {
	for i := range tt {
		tc := tt[i]

		t.Run(tc.nam, func(t *testing.T) {
			t.Parallel()

			if f(tc.err, tc.tgt...) != tc.exp {
				t.Errorf("actual=%v expected=%v", f(tc.err, tc.tgt...), tc.exp)
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
