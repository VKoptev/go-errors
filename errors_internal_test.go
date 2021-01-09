package errors

import (
	"errors"
	"testing"
)

var errEx = errors.New("ex")

func TestError_Is(t *testing.T) {
	t.Parallel()

	e1 := WithReason("e1")
	e2 := WithReason("e2")
	e3 := WithReason("e3")
	err := e1.WithErr(e2.WithErr(e3.WithErr(errEx)))

	if !errors.Is(err, e1) {
		t.Errorf("err{%v} is not e1{%v}", err, e1)
	}

	if !errors.Is(err, e2) {
		t.Errorf("err{%v} is not e2{%v}", err, e2)
	}

	if !errors.Is(err, e3) {
		t.Errorf("err{%v} is not e3{%v}", err, e3)
	}

	if !errors.Is(err, errEx) {
		t.Errorf("err{%v} is not ex{%v}", err, errEx)
	}
}

func TestError_WithErr(t *testing.T) {
	t.Parallel()

	e1 := WithReason("e1")
	e2 := e1.WithErr(errEx)
	err := e1.WithErr(e2)

	if !errors.Is(err, e1) {
		t.Errorf("err{%v} is not e1{%v}", err, e1)
	}

	if !errors.Is(err, e2) {
		t.Errorf("err{%v} is not e2{%v}", err, e2)
	}

	var terr *Error

	if errors.As(err.err, &terr) {
		t.Errorf("err{%v} contains e2{%v} insteadof replaced by e1{%v}", err, e2, e1)
	}

	e3 := e1.WithX(1)
	err = e1.WithErr(e3)

	if !errors.Is(err, e1) {
		t.Errorf("err{%v} is not e1{%v}", err, e1)
	}

	if !errors.Is(err, e2) {
		t.Errorf("err{%v} is not e2{%v}", err, e2)
	}

	if errors.As(err.err, &terr) {
		t.Errorf("err{%v} contains e3{%v} insteadof replaced by e1{%v}", err, e2, e1)
	}
}
