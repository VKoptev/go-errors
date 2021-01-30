package errors

import (
	"errors"
	"testing"
)

var errEx = errors.New("errEx")

//nolint:funlen,gocognit
func TestError_Is(t *testing.T) {
	t.Parallel()

	e1 := WithReason("e1")
	e2 := WithErr(errEx)
	e3 := WithX(1)

	tt := []struct {
		name string
		err  error
		is   []error
		nis  []error
	}{
		{name: "nil", err: nil, is: []error{}, nis: []error{e1, e2, e3, errEx}},
		{name: "e1", err: e1, is: []error{e1}, nis: []error{e2, e3, errEx}},
		{name: "with e1", err: WithErr(e1), is: []error{e1}, nis: []error{e2, e3, errEx}},
		{name: "e1 with e1", err: e1.WithErr(e1), is: []error{e1}, nis: []error{e2, e3, errEx}},
		{name: "with e1 with e1", err: WithErr(e1).WithErr(e1), is: []error{e1}, nis: []error{e2, e3, errEx}},
		{name: "e2", err: e2, is: []error{e2, errEx}, nis: []error{e1, e3}},
		{name: "with e2", err: WithErr(e2), is: []error{e2, errEx}, nis: []error{e1, e3}},
		{name: "e2 with e2", err: e2.WithErr(e2), is: []error{e2, errEx}, nis: []error{e1, e3}},
		{name: "with e2 with e2", err: WithErr(e2).WithErr(e2), is: []error{e2, errEx}, nis: []error{e1, e3}},
		{name: "e1 with e2", err: e1.WithErr(e2), is: []error{e1, e2, errEx}, nis: []error{e3}},
		{name: "e2 with e1", err: e2.WithErr(e1), is: []error{e1, e2, errEx}, nis: []error{e3}},
		{name: "with e1 with e2", err: WithErr(e1).WithErr(e2), is: []error{e1, e2, errEx}, nis: []error{e3}},
		{name: "with e2 with e1", err: WithErr(e2).WithErr(e1), is: []error{e1, e2, errEx}, nis: []error{e3}},
		{name: "e1 with e3", err: e1.WithErr(e3), is: []error{e1, e3}, nis: []error{e2, errEx}},
		// high-level WithX couldn't be determined
		{name: "e3 with e1", err: e3.WithErr(e1), is: []error{e1}, nis: []error{e2, e3, errEx}},
		{name: "with e1 with e3", err: WithErr(e1).WithErr(e3), is: []error{e1, e3}, nis: []error{e2, errEx}},
		// high-level WithX couldn't be determined
		{name: "with e3 with e1", err: WithErr(e3).WithErr(e1), is: []error{e1}, nis: []error{e2, e3, errEx}},
		{name: "e1 with e2 with e3", err: e1.WithErr(e2).WithErr(e3), is: []error{e1, e2, e3, errEx}, nis: []error{}},
		{name: "e1 with e3 with e2", err: e1.WithErr(e3).WithErr(e2), is: []error{e1, e2, errEx}, nis: []error{e3}},
		{name: "e2 with e1 with e3", err: e2.WithErr(e1).WithErr(e3), is: []error{e1, e2, e3, errEx}, nis: []error{}},
		{name: "e2 with e3 with e1", err: e2.WithErr(e3).WithErr(e1), is: []error{e1, e2, errEx}, nis: []error{e3}},
		{name: "e3 with e1 with e2", err: e3.WithErr(e1).WithErr(e2), is: []error{e1, e2, errEx}, nis: []error{e3}},
		{name: "e3 with e2 with e1", err: e3.WithErr(e2).WithErr(e1), is: []error{e1, e2, errEx}, nis: []error{e3}},
		{
			name: "thing",
			err: WithReason("q").WithErr(
				e1.WithErr(WithReason("w").WithErr(e2.WithErr(e3))).WithX(100),
			),
			is:  []error{e1, e2, e3, errEx},
			nis: []error{},
		},
	}

	for _, ti := range tt {
		ti := ti

		t.Run(ti.name, func(t *testing.T) {
			t.Parallel()

			for _, is := range ti.is {
				if !errors.Is(ti.err, is) {
					t.Errorf("err{%v} IS NOT exp{%v}", ti.err, is)
				}
			}

			for _, nis := range ti.nis {
				if errors.Is(ti.err, nis) {
					t.Errorf("err{%v} IS exp{%v}", ti.err, nis)
				}
			}
		})
	}
}

func TestIfErr(t *testing.T) {
	t.Parallel()

	if err := IfErr(nil); err != nil {
		t.Errorf("err{%v} is not nil", err)
	}

	e1 := WithReason("e1")
	err := IfErr(e1)

	if err == nil {
		t.Errorf("err{%v} is nil", err)
	}

	if !errors.Is(err, e1) {
		t.Errorf("err{%v} is not e1{%v}", err, e1)
	}
}

func TestError_IfErr(t *testing.T) {
	t.Parallel()

	e1 := WithReason("e1")

	if err := e1.IfErr(nil); err != nil {
		t.Errorf("err{%v} is not nil", err)
	}

	e2 := WithReason("e2")
	err := e1.IfErr(e2)

	if err == nil {
		t.Errorf("err{%v} is nil", err)
	}

	if !errors.Is(err, e1) {
		t.Errorf("err{%v} is not e1{%v}", err, e1)
	}

	if !errors.Is(err, e2) {
		t.Errorf("err{%v} is not e2{%v}", err, e2)
	}
}
